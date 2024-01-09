package api_test

import (
	"context"
	"fmt"
	"log"
	"net/http/httptest"
	"path/filepath"
	"time"

	"github.com/CommuniTEAM/CommuniTEA/api"
	"github.com/CommuniTEAM/CommuniTEA/auth"
	"github.com/CommuniTEAM/CommuniTEA/router"
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
	authTokens  []auth.TokenCookie // user[0], business[1], admin[2] jwt cookies
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

	api := &api.API{DBPool: dbPool}
	// suite.authTokens = []auth.TokenCookie{
	// 	auth.GenerateNewJWT(&auth.TokenData{}).TokenCookie,

	// }
	suite.server = httptest.NewServer(router.NewRouter(api))
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
