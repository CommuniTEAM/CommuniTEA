// Documentation you'll need to get acquainted with the APIs:

// https://pkg.go.dev/github.com/swaggest/rest#section-readme

// https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html

// https://github.com/jackc/pgx/wiki/Getting-started-with-pgx

// https://dev.to/vearutop/tutorial-developing-a-restful-api-with-go-json-schema-validation-and-openapi-docs-2490

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/CommuniTEAM/CommuniTEA/api"
	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest/response/gzip"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"
)

func main() {
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

	s.Post("/users", api.CreateUser())

	// Swagger UI endpoint at /docs.

	s.Docs("/docs", swgui.New)

	// Start server.

	log.Println("http://localhost:8000/docs")

	const three = 3

	// Run the server
	server := &http.Server{
		Addr:              ":8000",
		ReadHeaderTimeout: three * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
