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
	s := router.NewRouter(endpoints)

	// Configure and start the server
	const serverTimeout = 5
	server := &http.Server{
		Addr:              serverPort,
		Handler:           s,
		ReadHeaderTimeout: serverTimeout * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(fmt.Errorf("could not start the http server: %w", err))
	}
}
