package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CommuniTEAM/CommuniTEA/router"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRouter(t *testing.T) {
	mockDBPool, err := pgxmock.NewPool()
	if err != nil {
		t.Error(err)
	}

	r := router.NewRouter(mockDBPool)

	req, err := http.NewRequest(http.MethodGet, "/docs", nil)
	require.NoError(t, err)

	rw := httptest.NewRecorder()

	r.ServeHTTP(rw, req)
	assert.Equal(t, http.StatusOK, rw.Code)

	// actualSchema, err := assertjson.MarshalIndentCompact(json.RawMessage(rw.Body.Bytes()), "", "  ", 120)
	// require.NoError(t, err)

	// expectedSchema, err := os.ReadFile("_testdata/openapi.json")
	// require.NoError(t, err)

	// if !assertjson.Equal(t, expectedSchema, rw.Body.Bytes(), string(actualSchema)) {
	// 	require.NoError(t, os.WriteFile("_testdata/openapi_last_run.json", actualSchema, 0o600))
	// }
}
