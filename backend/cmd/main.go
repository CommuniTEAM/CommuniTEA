// Documentation you'll need to get acquainted with the APIs:
// https://pkg.go.dev/github.com/swaggest/rest#section-readme
// https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html
// https://github.com/jackc/pgx/wiki/Getting-started-with-pgx
// https://dev.to/vearutop/tutorial-developing-a-restful-api-with-go-json-schema-validation-and-openapi-docs-2490

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CommuniTEAM/CommuniTEA/api"
	"github.com/jackc/pgx/v5"
	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://admin:secret@postgres/communitea-db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	s := web.NewService(openapi31.NewReflector())

	// Init API documentation schema.
	s.OpenAPISchema().SetTitle("CommuniTEA API")
	s.OpenAPISchema().SetDescription("Let's gooooooooooooooo!")
	s.OpenAPISchema().SetVersion("v1.0.0")

	// Setup middlewares.
	s.Wrap(
		gzip.Middleware, // Response compression with support for direct gzip pass through.
	)

	// Add use case handler to router.
	s.Get("/hello/{name}", api.Greet())
	s.Get("/users", api.GetAllUsers())

	// Swagger UI endpoint at /docs.
	s.Docs("/docs", swgui.New)

	// Start server.
	log.Println("http://localhost:8000/docs")
	if err := http.ListenAndServe(":8000", s); err != nil {
		log.Fatal(err)
	}
}
