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
	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"
)

func main() {
	// Initialize database connection pool
	dbPool, err := db.NewDBPool(os.Getenv("DB_URI"))
	if err != nil {
		panic(err)
	}

	// Initialize web service
	s := web.NewService(openapi31.NewReflector())

	// Initialize API documentation schema
	s.OpenAPISchema().SetTitle("CommuniTEA API")
	s.OpenAPISchema().SetDescription("Bringing your community together over a cuppa")
	s.OpenAPISchema().SetVersion("v0.0.1")

	// Create custom schema mapping for 3rd party type uuid.
	uuidDef := jsonschema.Schema{}
	uuidDef.AddType(jsonschema.String)
	uuidDef.WithFormat("uuid")
	uuidDef.WithExamples("248df4b7-aa70-47b8-a036-33ac447e668d")
	s.OpenAPIReflector().JSONSchemaReflector().AddTypeMapping(uuid.UUID{}, uuidDef)
	s.OpenAPIReflector().JSONSchemaReflector().InlineDefinition(uuid.UUID{})

	// Set up middleware wraps
	s.Wrap(
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
		middleware.Logger,
	)

	// Forgive appended slashes on URLs
	s.Use(middleware.StripSlashes)

	// ! remove for prod - debug profiler
	s.Mount("/debug", middleware.Profiler())

	// Add API endpoints to router.
	// greeter (example endpoint to be removed for prod)
	s.Get("/hello/{name}", api.Greet())

	// auth
	s.Post("/login", api.UserLogin(dbPool))
	s.Delete("/logout", api.UserLogout())

	// accounts
	s.Post("/users", api.CreateUser(dbPool))

	// locations
	s.Post("/locations/cities", api.CreateCity(dbPool))

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

	// Check for VITE_API_HOST environment variable
	// ! Remove this code block for prod
	pubURL := os.Getenv("VITE_API_HOST")
	if pubURL == "" {
		log.Println("WARN: Could not find VITE_API_HOST var. Update .env file and rebuild docker containers.")
	} else {
		log.Printf("Starting server at %v/docs", pubURL)
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
