package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CommuniTEAM/CommuniTEA/api"
	"github.com/CommuniTEAM/CommuniTEA/router"
	"github.com/pashagolub/pgxmock/v3"
)

func TestNewRouter(t *testing.T) {
	mockDBPool, err := pgxmock.NewPool()

	if err != nil {
		t.Error(err)
	}

	endpoints := &api.API{DBPool: mockDBPool}

	r := router.NewRouter(endpoints)

	req, err := http.NewRequest(http.MethodGet, "/docs", nil)

	if err != nil {
		t.Fatalf("unexpected error while creating request: %v", err)
	}

	rw := httptest.NewRecorder()

	r.ServeHTTP(rw, req)

	if http.StatusOK != rw.Code {
		t.Fatalf("expected status code to be 200, but got: %v", rw.Code)
	}
}
