package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/CommuniTEAM/CommuniTEA/auth"
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

func (suite *UsersTestSuite) TestCreateUserAndPasswords() {
	t := suite.T()

	// CREATE USER
	// * Check 400 response for CreateUser
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

	// * Check 200 response, body, & cookie for CreateUser
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

	expectedBody, err := os.ReadFile("_testdata/users/token_data.json")
	require.NoError(t, err)

	assertjson.Equal(t, expectedBody, respBody)

	// Keep user data for ChangePassword
	var userData auth.TokenData
	err = json.Unmarshal(respBody, &userData)
	require.NoError(t, err)

	// * Check 409 response & body for CreateUser
	badInputs := []string{`{
		"username": "TestUser",
		"city_name": "Seattle",
		"state_code": "WA",
		"email": "test@email.com",
		"first_name": "Testy",
		"last_name": "Testington",
		"password": "string",
		"password_confirmation": "string",
		"role": "user"
	}`,
		`{
		"username": "User123",
		"city_name": "Seattle",
		"state_code": "WA",
		"email": "email@email.com",
		"first_name": "Testy",
		"last_name": "Testington",
		"password": "string",
		"password_confirmation": "string",
		"role": "user"
	}`}

	for _, input := range badInputs {
		req, err = http.NewRequest(http.MethodPost, suite.server.URL+"/users", bytes.NewBufferString(input))
		require.NoError(t, err)

		resp, err = http.DefaultTransport.RoundTrip(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusConflict, resp.StatusCode)

		respBody, err = io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.NoError(t, resp.Body.Close())

		assertjson.Equal(t, suite.errBody, respBody)
	}

	// * Check 400 response for CreateUser
	reqBody = []byte(`{
		"username": "User123",
		"city_name": "Seattle",
		"state_code": "WA",
		"email": "test@email.com",
		"first_name": "Testy",
		"last_name": "Testington",
		"password": "string",
		"password_confirmation": "badPass",
		"role": "user"
	}`)

	req, err = http.NewRequest(http.MethodPost, suite.server.URL+"/users", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// CHANGE PASSWORD
	// * Check 401 response & body for ChangePassword
	reqBody = []byte(fmt.Sprintf(`{
		"id": "%v",
		"old_password": "string",
		"new_password": "password",
		"new_password_conf": "password"
	}`, userData.ID.String()))

	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+userData.ID.String()+"/change-password", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Check 403 response & body for ChangePassword
	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+userData.ID.String()+"/change-password", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.user.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Check 200 response & body for ChangePassword
	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+userData.ID.String()+"/change-password", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: userData.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.successBody, respBody)

	// * Check 401 response for ChangePassword
	reqBody = []byte(fmt.Sprintf(`{
		"id": "%v",
		"old_password": "password",
		"new_password": "coolPassword",
		"new_password_conf": "notCoolPassword"
	}`, userData.ID.String()))

	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+userData.ID.String()+"/change-password", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: userData.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// USER LOGIN
	// * Check 401 response & cookie for UserLogin
	badInputs = []string{
		`{"username": "TestUser", "password": "fakePassword"}`,
		`{"username": "FakeUser", "password": "fakePassword"}`,
	}

	for _, input := range badInputs {
		req, err = http.NewRequest(http.MethodPost, suite.server.URL+"/login", bytes.NewBufferString(input))
		require.NoError(t, err)

		resp, err = http.DefaultTransport.RoundTrip(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	}

	// * Check 200 response, body, & cookie for UserLogin
	reqBody = []byte(`{
		"username": "TestUser",
		"password": "password"
	}`)

	req, err = http.NewRequest(http.MethodPost, suite.server.URL+"/login", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, expectedBody, respBody)
}

func (suite *UsersTestSuite) TestUpdateUser() {
	t := suite.T()

	// ID of the manually added user "user"
	userID := "372bcfb3-6b1d-4925-9f3d-c5ec683a4294"

	reqBody := []byte(`{
		"first_name": "Tester",
		"last_name": "Testerson",
		"role": "business"
	}`)

	// * Check 401 response & body
	req, err := http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+userID, bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, body)

	// * Check 403 response & body
	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+userID, bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.business.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, resp.StatusCode)

	body, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, body)

	// * Check 200 response
	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+userID, bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.user.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// * Check 400 response
	badInputs := []string{
		`{"city_name": "Kansas City", "state_code": "KY"}`,
		`{"city_name": "Arlington"}`,
		`{"state_code": "WI"}`,
		`{"email": "real-email"}`,
	}

	for _, input := range badInputs {
		req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+userID, bytes.NewBufferString(input))
		require.NoError(t, err)

		req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.user.Token})

		resp, err = http.DefaultTransport.RoundTrip(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	}
}

func (suite *UsersTestSuite) TestPromoteToAdmin() {
	t := suite.T()

	// ID of the manually added user "user"
	userID := "372bcfb3-6b1d-4925-9f3d-c5ec683a4294"

	// * Check 401 response & body
	req, err := http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+userID+"/promote", nil)
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, body)

	// * Check 403 response & body
	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+userID+"/promote", nil)
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.business.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, resp.StatusCode)

	body, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, body)

	// * Check 200 response & body
	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+userID+"/promote", nil)
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assert.Contains(t, string(body), "admin")

	// * Check 400 response
	userID = "fgdsd67uh-87yg-ghj" // invalid uuid

	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/users/"+userID+"/promote", nil)
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func (suite *UsersTestSuite) TestLogout() {
	t := suite.T()

	// * Check 200 response & body for logged-in user
	req, err := http.NewRequest(http.MethodDelete, suite.server.URL+"/logout", nil)
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.user.Token})

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	expectedBody := []byte(`{
		"access_token": "<ignore-diff>",
		"message": "success"
	}`)

	assertjson.Equal(t, expectedBody, body)

	// * Check 200 response & body for logged-out user
	req, err = http.NewRequest(http.MethodDelete, suite.server.URL+"/logout", nil)
	require.NoError(t, err)

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	expectedBody = []byte(`{
		"access_token": "",
		"message": "success"
	}`)

	assertjson.Equal(t, expectedBody, body)
}

func (suite *UsersTestSuite) TestDeleteUser() {
	t := suite.T()

	// ID of the manually added user "admin"
	userID := "e6473137-f4ef-46cc-a5e5-96ccb9d41043"

	// * Check 401 response & body
	req, err := http.NewRequest(http.MethodDelete, suite.server.URL+"/users/"+userID, nil)
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, body)

	// * Check 403 response & body
	req, err = http.NewRequest(http.MethodDelete, suite.server.URL+"/users/"+userID, nil)
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.user.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, resp.StatusCode)

	body, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, body)

	// * Check 200 response & body
	req, err = http.NewRequest(http.MethodDelete, suite.server.URL+"/users/"+userID, nil)
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	expectedBody := []byte(`{
		"access_token": "<ignore-diff>",
		"message": "success"
	}`)

	assertjson.Equal(t, expectedBody, body)
}
