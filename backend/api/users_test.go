package api_test

import (
	"bytes"
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

func (suite *UsersTestSuite) TestCreateUser() {
	t := suite.T()

	// * Check 400 response
	reqBody := []byte(`{
		"username": "TestUser",
		"state_code": "WA",
		"email": "email@email.com",
		"first_name": "Testy",
		"last_name": "Testington",
		"password": "string",
		"password_confirmation": "string",
		"role": "user"
	}`)

	req, err := http.NewRequest(http.MethodPost, suite.server.URL+"/users", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// * Check 200 response & body
	reqBody = []byte(`{
		"username": "TestUser",
		"city_name": "Seattle",
		"state_code": "WA",
		"email": "email@email.com",
		"first_name": "Testy",
		"last_name": "Testington",
		"password": "string",
		"password_confirmation": "string",
		"role": "user"
	}`)

	req, err = http.NewRequest(http.MethodPost, suite.server.URL+"/users", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	expectedBody := []byte(`{
		"access_token": "<ignore-diff>",
		"token_type": "bearer",
		"expires_in": 3600,
		"id": "<ignore-diff>",
		"username": "TestUser",
		"email": "email@email.com",
		"first_name": "Testy",
		"last_name": "Testington",
		"role": "user",
		"location": {
			"id": "4c33e0bc-3d43-4e77-aed0-b7aff09bb689",
			"name": "Seattle",
			"state": "WA"
		}
	}`)

	assertjson.Equal(t, expectedBody, respBody)

	// * Check 409 response & body
	req, err = http.NewRequest(http.MethodPost, suite.server.URL+"/users", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusConflict, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)
}
