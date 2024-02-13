package router

import (
	"net/http"

	"github.com/CommuniTEAM/CommuniTEA/api"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
	"github.com/rs/cors"
	"github.com/swaggest/jsonschema-go"
	oapi "github.com/swaggest/openapi-go"
	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"
)

const securityName = "authCookie"

// httpResponse provides generic json schema for an http response's
// accompanying json body.
type httpResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func addEndpoints(s *web.Service, endpoints *api.API) *web.Service {
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

	// auth
	s.Post("/login", endpoints.UserLogin())
	s.Delete("/logout", endpoints.UserLogout(), requireAuth)

	// users
	s.Post("/users", endpoints.CreateUser())
	s.Get("/users/{id}", endpoints.GetUser())
	s.Put("/users/{id}", endpoints.UpdateUser())
	s.Put("/users/{id}/change-password", endpoints.ChangePassword())
	s.Put("/users/{id}/promote", endpoints.PromoteToAdmin())
	s.Delete("/users/{id}", endpoints.DeleteUser())

	// locations
	s.Post("/locations/cities", endpoints.CreateCity(), requireAuth)
	s.Get("/locations/cities", endpoints.GetAllCities())
	s.Get("/locations/cities/{id}", endpoints.GetCity())
	s.Get("/locations/states", endpoints.GetAllStates())
	s.Get("/locations/states/{state-code}/cities", endpoints.GetAllCitiesInState())
	s.Put("/locations/cities/{id}", endpoints.UpdateCity(), requireAuth)
	s.Delete("/locations/cities/{id}", endpoints.DeleteCity(), requireAuth)

	// wikiteadia
	s.Get("/teas/{published}", endpoints.GetAllTeas())
	s.Post("/teas", endpoints.CreateTea())

	return s
}

// NewRouter creates a custom router for the http server in line with
// openapi specifications. It bundles the included api endpoints into
// the Swagger UI in the browser, available at /docs.
func NewRouter(endpoints *api.API, envType string) http.Handler {
	// Initialize openAPI 3.1 reflector
	reflector := openapi31.NewReflector()

	// Declare security scheme
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

	// Set up middleware wraps
	s.Wrap(
		middleware.Logger,
		cors.New(cors.Options{
			AllowedOrigins:      []string{"http://localhost:3000", "https://communitea.life"},
			AllowedMethods:      []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:      []string{"Content-Type"},
			AllowCredentials:    true,
			AllowPrivateNetwork: true,
		}).Handler,

		// Describe additional response schema
		nethttp.OpenAPIAnnotationsMiddleware(s.OpenAPICollector, func(oc oapi.OperationContext) error {
			oc.AddRespStructure(httpResponse{}, func(cu *oapi.ContentUnit) {
				cu.HTTPStatus = http.StatusBadRequest
			})
			oc.AddRespStructure(httpResponse{}, func(cu *oapi.ContentUnit) {
				cu.HTTPStatus = http.StatusConflict
			})
			return nil
		}),
	)

	// Prepend "/api" to endpoint URIs in production
	if envType == "prod" {
		s.Wrap(func(handler http.Handler) http.Handler {
			var withRoute rest.HandlerWithRoute
			if nethttp.HandlerAs(handler, &withRoute) {
				return nethttp.HandlerWithRouteMiddleware(
					withRoute.RouteMethod(),
					"/api"+withRoute.RoutePattern(),
				)(handler)
			}
			return handler
		})
	}

	// Forgive appended slashes on URLs
	s.Use(middleware.StripSlashes)

	// Add API endpoints to router
	addEndpoints(s, endpoints)

	// Swagger UI endpoint at /docs
	s.Docs("/docs", swgui.New)
	s.Handle("/", http.RedirectHandler(
		func(env string) string {
			if env == "prod" {
				return "/api/docs"
			}
			return "/docs"
		}(envType),
		http.StatusSeeOther,
	))

	return s
}
