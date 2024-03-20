package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CommuniTEAM/CommuniTEA/api"
	"github.com/CommuniTEAM/CommuniTEA/auth"
	"github.com/CommuniTEAM/CommuniTEA/router"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

const (
	serverPort  = ":8000"
	prodEnvPath = "/usr/lib/communitea-api/.env"
)

//nolint:gochecknoglobals // Changed with ldflags at build time
var Env = "dev"

func main() {
	if Env == "prod" {
		log.Println("INFO: Initializing PRODUCTION environment.")

		// Get production environment variables
		err := godotenv.Load(prodEnvPath)
		if err != nil {
			log.Fatal(fmt.Errorf("could not read environment variables: %w", err))
		}
		log.Printf("Starting server on port %v\n", serverPort)
	} else {
		log.Println("INFO: Initializing DEVELOPMENT environment.")

		// Check for VITE_API_HOST environment variable if dev environment
		pubURL := os.Getenv("VITE_API_HOST")
		if pubURL == "" {
			log.Println("WARN: Could not find VITE_API_HOST var. Update .env file and rebuild docker containers.")
		} else {
			log.Printf("Starting server at %v/docs", pubURL)
		}
	}

	// Initialize database connection
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal(fmt.Errorf("could not create new db pool: %w", err))
	}

	// Initialize user authentication system
	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		log.Fatal(fmt.Errorf("could not initialize authenticator: %w", err))
	}

	// Initialize endpoints
	endpoints := &api.API{DBPool: dbPool, Auth: authenticator}

	// Initialize router
	s := router.NewRouter(endpoints, Env)

	// Wrap the router with the CORS middleware
	wrappedRouter := corsMiddleware(s)

	// Configure and start the server
	const serverTimeout = 5
	server := &http.Server{
		Addr:              serverPort,
		Handler:           wrappedRouter,
		ReadHeaderTimeout: serverTimeout * time.Second,
	}

	if Env == "prod" {
		err = server.ListenAndServeTLS(os.Getenv("SSL_CERT"), os.Getenv("SSL_KEY"))
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		log.Fatal(fmt.Errorf("could not start the http(s) server: %w", err))
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
