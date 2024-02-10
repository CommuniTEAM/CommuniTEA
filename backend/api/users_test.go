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

type UsersTestSuite struct {
	TestSuite
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}

func (suite *UsersTestSuite) TestGetUser() {
	t := suite.T()

	// ID of the manually added user "user"
	userID := "372bcfb3-6b1d-4925-9f3d-c5ec683a4294"

	// * Check 200 response & body
	req, err := http.NewRequest(http.MethodGet, suite.server.URL+"/users/"+userID, nil)
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	expectedBody, err := os.ReadFile("_testdata/users/get_user.json")
	require.NoError(t, err)

	assertjson.Equal(t, expectedBody, body)

	// * Check 404 response
	userID = "4c33e0bc-3d43-4e77-aed0-b7aff09bb600" // a uuid not in users table

	req, err = http.NewRequest(http.MethodGet, suite.server.URL+"/users/"+userID, nil)
	require.NoError(t, err)

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	// * Check 400 response
	userID = "4c33e0bc-3d43-4e77-aed0-b7af" // not a valid uuid

	req, err = http.NewRequest(http.MethodGet, suite.server.URL+"/users/"+userID, nil)
	require.NoError(t, err)

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
