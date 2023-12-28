package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CommuniTEAM/CommuniTEA/api"
	"github.com/CommuniTEAM/CommuniTEA/db"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/rs/cors"
	"github.com/swaggest/jsonschema-go"
	oapi "github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"
)

type httpResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func main() {
	// Set environment status
	isDevEnv := os.Getenv("DEV")

	// Initialize database connection pool
	dbPool, err := db.NewDBPool(os.Getenv("DB_URI"))
	if err != nil {
		panic(err)
	}

	// Initialize openAPI 3.1 reflector
	reflector := openapi31.NewReflector()

	// Declare security scheme
	securityName := "authCookie"
	reflector.SpecEns().SetHTTPBearerTokenSecurity(securityName, "cookie", "User Authentication")

	// Initialize web service
	s := web.NewService(reflector)

	// Initialize API documentation schema
	s.OpenAPISchema().SetTitle("CommuniTEA API")
	s.OpenAPISchema().SetDescription("Bringing your community together over a cuppa")
	s.OpenAPISchema().SetVersion("v0.0.1")

	// Create custom schema mapping for 3rd party type uuid
	uuidDef := jsonschema.Schema{}
	uuidDef.AddType(jsonschema.String)
	uuidDef.WithFormat("uuid")
	uuidDef.WithExamples("248df4b7-aa70-47b8-a036-33ac447e668d")
	s.OpenAPIReflector().JSONSchemaReflector().AddTypeMapping(uuid.UUID{}, uuidDef)
	s.OpenAPIReflector().JSONSchemaReflector().InlineDefinition(uuid.UUID{})

	// Set up middleware wraps
	s.Wrap(
		middleware.Logger,
		cors.New(cors.Options{
			AllowedOrigins:             []string{"http://localhost:3000", "https://communitea.life"},
			AllowedMethods:             []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:             []string{"Content-Type"},
			ExposedHeaders:             []string{},
			OptionsPassthrough:         false,
			OptionsSuccessStatus:       http.StatusNoContent,
			Debug:                      false,
			AllowOriginFunc:            nil,
			AllowOriginRequestFunc:     nil,
			AllowOriginVaryRequestFunc: nil,
			MaxAge:                     0,
			AllowCredentials:           true,
			AllowPrivateNetwork:        false,
			Logger:                     nil,
		}).Handler,
		// Describe bad request (400) response
		nethttp.OpenAPIAnnotationsMiddleware(s.OpenAPICollector, func(oc oapi.OperationContext) error {
			oc.AddRespStructure(httpResponse{}, func(cu *oapi.ContentUnit) {
				cu.HTTPStatus = http.StatusBadRequest
			})
			return nil
		}),
	)

	// Forgive appended slashes on URLs
	s.Use(middleware.StripSlashes)

	// Mount debug profiler in dev environment
	if isDevEnv == "true" {
		s.Mount("/debug", middleware.Profiler())
	}

	// Set up auth requirement option for routes
	requireAuth := nethttp.AnnotateOpenAPIOperation(func(oc oapi.OperationContext) error {
		// Add security requirement to operation
		oc.AddSecurity(securityName)

		// Describe unauthenticated response
		oc.AddRespStructure(httpResponse{}, func(cu *oapi.ContentUnit) {
			cu.HTTPStatus = http.StatusUnauthorized
		})

		// Describe unauthorized (forbidden) response
		oc.AddRespStructure(httpResponse{}, func(cu *oapi.ContentUnit) {
			cu.HTTPStatus = http.StatusForbidden
		})
		return nil
	})

	// Add API endpoints to router
	// greeter (example endpoint to be removed for prod)
	s.Get("/hello/{name}", api.Greet())

	// auth
	s.Post("/login", api.UserLogin(dbPool))
	s.Delete("/logout", api.UserLogout(), requireAuth)

	// users
	s.Post("/users", api.CreateUser(dbPool))

	// locations
	s.Post("/locations/cities", api.CreateCity(dbPool), requireAuth)

	// wikiteadia
	s.Get("/teas/{published}", api.GetAllTeas(dbPool))
	s.Post("/teas", api.CreateTea(dbPool))

	// Swagger UI endpoint at /docs.
	s.Docs("/docs", swgui.New)

	// Configure and start the server
	const serverTimeout = 5
	server := &http.Server{
		Addr:              ":8000",
		Handler:           s,
		ReadHeaderTimeout: serverTimeout * time.Second,
	}

	// Check for VITE_API_HOST environment variable if dev environment
	if isDevEnv == "true" {
		pubURL := os.Getenv("VITE_API_HOST")
		if pubURL == "" {
			log.Println("WARN: Could not find VITE_API_HOST var. Update .env file and rebuild docker containers.")
		} else {
			log.Printf("Starting server at %v/docs", pubURL)
		}
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
