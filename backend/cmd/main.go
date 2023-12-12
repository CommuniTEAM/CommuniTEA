// Documentation you'll need to get acquainted with the APIs:
// https://pkg.go.dev/github.com/swaggest/rest#section-readme
// https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html
// https://github.com/jackc/pgx/wiki/Getting-started-with-pgx
// https://dev.to/vearutop/tutorial-developing-a-restful-api-with-go-json-schema-validation-and-openapi-docs-2490

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CommuniTEAM/CommuniTEA/api"
	"github.com/CommuniTEAM/CommuniTEA/db"
	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"
)

func main() {
	// Init database connection pool
	dbPool, err := db.NewDBPool(os.Getenv("DB_URI"))
	if err != nil {
		panic(err)
	}

	s := web.NewService(openapi31.NewReflector())

	// Init API documentation schema.
	s.OpenAPISchema().SetTitle("CommuniTEA API")
	s.OpenAPISchema().SetDescription("Bringing your community together over a cuppa")
	s.OpenAPISchema().SetVersion("v0.0.1")

	// Setup middlewares.
	s.Wrap(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)

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

	// Start server.
	pubURL := os.Getenv("PUBLIC_URL")

	if pubURL == "" {
		log.Println("WARN: Could not find PUBLIC_URL var. Update .env file and rebuild docker containers.")
	} else {
		log.Printf("Starting server at %v/docs", pubURL)
	}

	// Run the server
	const serverTimeout = 5
	server := &http.Server{
		Addr:              ":8000",
		Handler:           s,
		ReadHeaderTimeout: serverTimeout * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
