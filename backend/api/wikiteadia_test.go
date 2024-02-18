package api_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/swaggest/assertjson"
)

type WikiteadiaTestSuite struct {
	TestSuite
}

func TestWikiteadiaTestSuite(t *testing.T) {
	suite.Run(t, new(WikiteadiaTestSuite))
}

func (suite *WikiteadiaTestSuite) TestCreateTea() {
	t := suite.T()

	reqBody := []byte(`{
		"name": "Earl Grey",
  		"img_url": "https://images.unsplash.com/photo-1605618826115-fb9e775cfb40?q=80&w=1779&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
  		"description": "It is a black tea mix",
  		"brew_time": "20 minutes",
  		"brew_temp": 175,
  		"published": false
	}`)

	// * Check 401 response & body

	req, err := http.NewRequest(http.MethodPost, suite.server.URL+"/teas", bytes.NewBuffer(reqBody))

	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)

	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)

	require.NoError(t, err)

	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Check 400 response

	badInput := []byte(
		`{"name": 123,
		"img_url": "https://images.unsplash.com/photo-1605618826115-fb9e775cfb40?q=80&w=1779&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"description": "It is a black tea mix from England",
		"brew_time": "5 minutes",
		"brew_temp": "175",
		"published": false}`)

	// for _, input := range badInputs {
	req, err = http.NewRequest(http.MethodPost, suite.server.URL+"/teas", bytes.NewBuffer(badInput))

	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.user.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)

	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// * Check 200 response & body
	reqBody = []byte(`{
		"name": "Earl Grey",
  		"img_url": "https://images.unsplash.com/photo-1605618826115-fb9e775cfb40?q=80&w=1779&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
  		"description": "It is a black tea mix",
  		"brew_time": "20 minutes",
  		"brew_temp": 175,
  		"published": false
	}`)

	req, err = http.NewRequest(http.MethodPost, suite.server.URL+"/teas", bytes.NewBuffer(reqBody))

	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.user.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)

	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)

	require.NoError(t, err)

	require.NoError(t, resp.Body.Close())

	expectedBody := []byte(`{
		"id": "<ignore-diff>",
		"name": "Earl Grey",
		"img_url": "https://images.unsplash.com/photo-1605618826115-fb9e775cfb40?q=80&w=1779&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"description": "It is a black tea mix",
		"brew_time": "20 minutes",
		"brew_temp": 175,
		"published": false
		}`)

	assertjson.Equal(t, expectedBody, respBody)
}

func (suite *WikiteadiaTestSuite) TestGetAllUnpublishedTeas() {
	t := suite.T()

	req, err := http.NewRequest(http.MethodGet, suite.server.URL+"/teas/false", nil)

	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)

	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func (suite *WikiteadiaTestSuite) TestGetAllPublishedTeas() {
	t := suite.T()

	req, err := http.NewRequest(http.MethodGet, suite.server.URL+"/teas/true", nil)

	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)

	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func (suite *WikiteadiaTestSuite) TestUpdateTea() {
	t := suite.T()

	teaID := "c64ff5ab-7323-4142-9077-aea320c3c4cc"

	reqBody := `{
		"name": "Earl Grey",
  		"img_url": "https://images.unsplash.com/photo-1605618826115-fb9e775cfb40?q=80&w=1779&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
  		"description": "It is a black tea mix",
  		"brew_time": "20 minutes",
  		"brew_temp": 175,
  		"published": false
	}`

	// * Check 401 response & body

	req, err := http.NewRequest(http.MethodPut, suite.server.URL+"/teas/"+teaID, bytes.NewBufferString(reqBody))

	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)

	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)

	require.NoError(t, err)

	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Check 403 response & body

	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/teas/"+teaID, bytes.NewBufferString(reqBody))

	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.user.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)

	require.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)

	require.NoError(t, err)

	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Check 404 response & body

	badInput := "a64ff0aa-7003-4142-9077-aea320c3c4bb" // fake uuid

	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/teas/"+badInput, bytes.NewBufferString(reqBody))

	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)

	require.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)

	require.NoError(t, err)

	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Check 400 response

	badInputs := []string{
		`{"name": "Earl Grey",
		"img_url": "https://images.unsplash.com/photo-1605618826115-fb9e775cfb40?q=80&w=1779&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"description": "It is a black tea mix from England",
		"brew_time": 5,
		"brew_temp": "175F",
		"published": true}`,
	}

	for _, input := range badInputs {
		req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/teas/"+teaID, bytes.NewBufferString(input))

		require.NoError(t, err)

		req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

		resp, err = http.DefaultTransport.RoundTrip(req)

		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	}

	// * Check 200 response & body

	reqBody = `{"id": "c64ff5ab-7323-4142-9077-aea320c3c4cc",
    "name": "Earl Grey",
    "img_url": "https://images.unsplash.com/photo-1605618826115-fb9e775cfb40?q=80&w=1779&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
    "description": "It is a black tea mix from England",
    "brew_time": "5 minutes",
    "brew_temp": 175,
    "published": true}`

	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/teas/"+teaID, bytes.NewBufferString(reqBody))

	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)

	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)

	require.NoError(t, err)

	require.NoError(t, resp.Body.Close())

	expectedBody := []byte(`{
  "id": "c64ff5ab-7323-4142-9077-aea320c3c4cc",
  "name": "Earl Grey",
  "img_url": "https://images.unsplash.com/photo-1605618826115-fb9e775cfb40?q=80&w=1779&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
  "description": "It is a black tea mix from England",
  "brew_time": "5 minutes",
  "brew_temp": 175,
  "published": true
}`)

	assertjson.Equal(t, expectedBody, respBody)
}
