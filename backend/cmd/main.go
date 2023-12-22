package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CommuniTEAM/CommuniTEA/api"
	"github.com/CommuniTEAM/CommuniTEA/db"
	"github.com/rs/cors"
	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest/response/gzip"
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

	// Setup middlewares
	s.Wrap(gzip.Middleware) // Response compression with support for direct gzip pass through

	// Add API endpoints to router.
	// greeter (example endpoint to be removed for prod)
	s.Get("/hello/{name}", api.Greet())

	// locations
	s.Post("/locations/cities", api.CreateCity(dbPool))

	// wikiteadia
	s.Get("/teas/{published}", api.GetAllTeas(dbPool))
	s.Post("/teas", api.CreateTea(dbPool))

	// Swagger UI endpoint at /docs.
	s.Docs("/docs", swgui.New)

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:             []string{"http://localhost:3000"}, // Set the allowed origins here
		AllowedMethods:             []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:             []string{"Content-Type"},
		ExposedHeaders:             []string{},
		OptionsPassthrough:         false,
		OptionsSuccessStatus:       0,
		Debug:                      false,
		AllowOriginFunc:            nil,
		AllowOriginRequestFunc:     nil,
		AllowOriginVaryRequestFunc: nil,
		MaxAge:                     0,
		AllowCredentials:           false,
		AllowPrivateNetwork:        false,
		Logger:                     nil,
	})

	// Configure and start the server
	const serverTimeout = 5
	server := &http.Server{
		Addr:              ":8000",
		Handler:           c.Handler(s), // Wrap the service with CORS middleware
		ReadHeaderTimeout: serverTimeout * time.Second,
	}

	// Check for PUBLIC_URL environment variable
	pubURL := os.Getenv("PUBLIC_URL")
	if pubURL == "" {
		log.Println("WARN: Could not find PUBLIC_URL var. Update .env file and rebuild docker containers.")
	} else {
		log.Printf("Starting server at %v/docs", pubURL)
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
