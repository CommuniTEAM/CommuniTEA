package api_test

import (
	"context"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"path/filepath"
	"time"

	"github.com/CommuniTEAM/CommuniTEA/api"
	"github.com/CommuniTEAM/CommuniTEA/auth"
	"github.com/CommuniTEAM/CommuniTEA/router"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

// TestSuite is a helper struct that enables multiple
// tests to be run against one test database. To create a
// test suite for a series of endpoints, build a new struct
// that embeds this one.
type TestSuite struct {
	suite.Suite
	pgContainer *PostgresContainer
	server      *httptest.Server
	ctx         context.Context
	authTokens  userTokens // jwts for every kind of user role
	errBody     []byte     // response body for error codes
	successBody []byte     // response body for success messages
}

// SetupSuite instantiates a test suite by setting up a new test
// database container, connecting it to the API, and starting an
// httptest server.
func (suite *TestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := CreatePostgresContainer(suite.ctx)
	if err != nil {
		// errors from CreatePostgresContainer are already wrapped
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer

	dbPool, err := pgxpool.New(context.Background(), pgContainer.ConnectionString)
	if err != nil {
		log.Fatalf("could not create new db pool: %v", err)
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		log.Fatalf("could not initialize authenticator: %v", err)
	}

	api := &api.API{DBPool: dbPool, Auth: authenticator}
	suite.server = httptest.NewServer(router.NewRouter(api))

	suite.errBody, err = os.ReadFile("_testdata/error_response.json")
	if err != nil {
		log.Fatalf("could not read _testdata/error_response.json")
	}
	suite.successBody, err = os.ReadFile("_testdata/success_response.json")
	if err != nil {
		log.Fatalf("could not read _testdata/success_response.json")
	}

	// Generate jwts for different user roles
	// data aligns with _testdata/db_userdata_migration.sql
	userData := auth.TokenData{
		ExpiresIn: 3600,
		ID:        uuid.MustParse("372bcfb3-6b1d-4925-9f3d-c5ec683a4294"),
		Role:      "user",
		Username:  "user",
		Location:  uuid.MustParse("4c33e0bc-3d43-4e77-aed0-b7aff09bb689"),
	}
	userToken, err := api.Auth.GenerateNewJWT(&userData, false)
	if err != nil {
		log.Fatalf("could not generate user token: %v", err)
	}
	suite.authTokens.user = userToken.TokenCookie

	businessData := auth.TokenData{
		ExpiresIn: 3600,
		ID:        uuid.MustParse("140e4411-a7f7-4c50-a2d4-f3d3fc9fc550"),
		Role:      "business",
		Username:  "business",
		Location:  uuid.MustParse("4c33e0bc-3d43-4e77-aed0-b7aff09bb689"),
	}
	businessToken, err := api.Auth.GenerateNewJWT(&businessData, false)
	if err != nil {
		log.Fatalf("could not generate business token: %v", err)
	}
	suite.authTokens.business = businessToken.TokenCookie

	adminData := auth.TokenData{
		ExpiresIn: 3600,
		ID:        uuid.MustParse("e6473137-f4ef-46cc-a5e5-96ccb9d41043"),
		Role:      "admin",
		Username:  "admin",
		Location:  uuid.MustParse("4c33e0bc-3d43-4e77-aed0-b7aff09bb689"),
	}
	adminToken, err := api.Auth.GenerateNewJWT(&adminData, false)
	if err != nil {
		log.Fatalf("could not generate admin token: %v", err)
	}
	suite.authTokens.admin = adminToken.TokenCookie
}

// TearDownSuite terminates the test database container used by the suite and closes the httptest server.
func (suite *TestSuite) TearDownSuite() {
	suite.server.Close()
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %v", err)
	}
}

// CreatePostgresContainer is a helper function that sets up a test database
// container and returns its pointer.
func CreatePostgresContainer(ctx context.Context) (*PostgresContainer, error) {
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(
			filepath.Join("..", "db", "01_schema.sql"),
			filepath.Join("..", "db", "migrations", "20231201090112_populate_location_states.sql"),
			filepath.Join("..", "db", "migrations", "20231212061145_populate_user_roles.sql"),
			filepath.Join("..", "db", "migrations", "20240106002631_add_inital_locations.sql"),
			filepath.Join("_testdata", "db_testdata_migration.sql"),
		),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, fmt.Errorf("could not create postgres container: %w", err)
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("could not acquire test-db connection string: %w", err)
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connStr,
	}, nil
}

type userTokens struct {
	user     auth.TokenCookie
	business auth.TokenCookie
	admin    auth.TokenCookie
}
