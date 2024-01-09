package api_test

import (
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/swaggest/assertjson"
)

type LocationsTestSuite struct {
	TestSuite
}

func TestLocationTestSuite(t *testing.T) {
	suite.Run(t, new(LocationsTestSuite))
}

func (suite *LocationsTestSuite) TestGetAllStates() {
	t := suite.T()

	req, err := http.NewRequest(http.MethodGet, suite.server.URL+"/locations/states", nil)
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	expectedBody, err := os.ReadFile("_testdata/locations/get_all_states.json")
	require.NoError(t, err)

	assertjson.Equal(t, expectedBody, body)
}

func (suite *LocationsTestSuite) TestGetCity() {
	t := suite.T()

	// ID of the manually added Seattle entry
	cityID := "4c33e0bc-3d43-4e77-aed0-b7aff09bb689"

	// Check 200 response & body
	req, err := http.NewRequest(http.MethodGet, suite.server.URL+"/locations/cities/"+cityID, nil)
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	expectedBody, err := os.ReadFile("_testdata/locations/get_city.json")
	require.NoError(t, err)

	assertjson.Equal(t, expectedBody, body)

	// Check 404 response
	cityID = "4c33e0bc-3d43-4e77-aed0-b7aff09bb600" // a uuid not in cities table

	req, err = http.NewRequest(http.MethodGet, suite.server.URL+"/locations/cities/"+cityID, nil)
	require.NoError(t, err)

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
