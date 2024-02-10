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

func (suite *LocationsTestSuite) TestGetAllCities() {
	t := suite.T()

	req, err := http.NewRequest(http.MethodGet, suite.server.URL+"/locations/cities", nil)
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func (suite *LocationsTestSuite) TestGetCity() {
	t := suite.T()

	// ID of the manually added Seattle entry
	cityID := "4c33e0bc-3d43-4e77-aed0-b7aff09bb689"

	// * Check 200 response & body
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

	// * Check 404 response
	cityID = "4c33e0bc-3d43-4e77-aed0-b7aff09bb600" // a uuid not in cities table

	req, err = http.NewRequest(http.MethodGet, suite.server.URL+"/locations/cities/"+cityID, nil)
	require.NoError(t, err)

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	// * Check 400 response
	cityID = "4c33e0bc-3d43-4e77-aed0-b7af" // not a valid uuid

	req, err = http.NewRequest(http.MethodGet, suite.server.URL+"/locations/cities/"+cityID, nil)
	require.NoError(t, err)

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func (suite *LocationsTestSuite) TestGetCitiesInState() {
	t := suite.T()

	testState := "WA"

	// * Check 200 response & body
	req, err := http.NewRequest(http.MethodGet, suite.server.URL+"/locations/states/"+testState+"/cities", nil)
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	expectedBody, err := os.ReadFile("_testdata/locations/get_cities_in_state.json")
	require.NoError(t, err)

	assertjson.Equal(t, expectedBody, body)

	// * Check 400 response
	badInputs := []string{"WAA", "W", "12", "VR"}

	for _, input := range badInputs {
		req, err = http.NewRequest(http.MethodGet, suite.server.URL+"/locations/states/"+input+"/cities", nil)
		require.NoError(t, err)

		resp, err = http.DefaultTransport.RoundTrip(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	}
}

func (suite *LocationsTestSuite) TestCreateCity() {
	t := suite.T()

	// * Check 401 response & body
	reqBody := []byte(`{
		"city_name": "Chicago",
		"state_code": "IL"
	}`)
	req, err := http.NewRequest(http.MethodPost, suite.server.URL+"/locations/cities", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Check 403 response & body
	req, err = http.NewRequest(http.MethodPost, suite.server.URL+"/locations/cities", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.user.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Check 200 response & body
	req, err = http.NewRequest(http.MethodPost, suite.server.URL+"/locations/cities", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	expectedBody := []byte(`{
		"id": "<ignore-diff>",
		"name": "Chicago",
		"state": "IL"
	}`)

	assertjson.Equal(t, expectedBody, respBody)

	// * Check 409 response & body
	req, err = http.NewRequest(http.MethodPost, suite.server.URL+"/locations/cities", bytes.NewBuffer(reqBody))
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusConflict, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Check 400 response
	badInputs := []string{
		`{"name": "Chicago", "state": "ILLLL"}`,
		`{"name": "Chicago", "state": "15"}`,
		`{"state": "IL"}`,
		`{"name": "Chicago"}`,
		`{"name": "", "state": ""}`,
	}

	for _, input := range badInputs {
		req, err = http.NewRequest(http.MethodPost, suite.server.URL+"/locations/cities", bytes.NewBufferString(input))
		require.NoError(t, err)

		req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

		resp, err = http.DefaultTransport.RoundTrip(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	}
}

func (suite *LocationsTestSuite) TestUpdateCity() {
	t := suite.T()

	// ID of the manually added New York entry
	cityID := "6937755c-7e87-4226-9692-36d3019be32a"
	reqBody := `{"name": "New York"}`

	// * Check 401 response & body
	req, err := http.NewRequest(http.MethodPut, suite.server.URL+"/locations/cities/"+cityID, bytes.NewBufferString(reqBody))
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Check 403 response & body
	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/locations/cities/"+cityID, bytes.NewBufferString(reqBody))
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.user.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Check 200 response & body
	reqBody = `{"name": "Buffalo"}`

	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/locations/cities/"+cityID, bytes.NewBufferString(reqBody))
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	expectedBody := []byte(`{
		"id": "6937755c-7e87-4226-9692-36d3019be32a",
		"name": "Buffalo",
		"state": "NY"
	}`)

	assertjson.Equal(t, expectedBody, respBody)

	// * Check 409 response & body
	badInputs := []string{
		`{"name": "Buffalo"}`,
		`{"name": "buffalo"}`,
	}
	for _, input := range badInputs {
		req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/locations/cities/"+cityID, bytes.NewBufferString(input))
		require.NoError(t, err)

		req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

		resp, err = http.DefaultTransport.RoundTrip(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusConflict, resp.StatusCode)

		respBody, err = io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.NoError(t, resp.Body.Close())

		assertjson.Equal(t, suite.errBody, respBody)
	}

	// * Check 404 response & body
	badInput := "4c33e0bc-3d43-4e77-aed0-b7aff0900000" // fake uuid

	req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/locations/cities/"+badInput, bytes.NewBufferString(reqBody))
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
	badInputs = []string{
		`{"name": ""}`,
		`{}`,
		`{"state": "MT"}`,
	}
	for _, input := range badInputs {
		req, err = http.NewRequest(http.MethodPut, suite.server.URL+"/locations/cities/"+cityID, bytes.NewBufferString(input))
		require.NoError(t, err)

		req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

		resp, err = http.DefaultTransport.RoundTrip(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	}
}

func (suite *LocationsTestSuite) TestDeleteCity() {
	t := suite.T()

	// ID of the manually added Kansas City entry
	cityID := "07eca16a-8ee1-4c1a-831e-cb984a851bf3"

	// * Test 401 response & body
	req, err := http.NewRequest(http.MethodDelete, suite.server.URL+"/locations/cities/"+cityID, nil)
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Test 403 response & body
	req, err = http.NewRequest(http.MethodDelete, suite.server.URL+"/locations/cities/"+cityID, nil)
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.user.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)

	// * Test 400 response
	badInput := "lkashdflkj"

	req, err = http.NewRequest(http.MethodDelete, suite.server.URL+"/locations/cities/"+badInput, nil)
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// * Test 200 response & body
	req, err = http.NewRequest(http.MethodDelete, suite.server.URL+"/locations/cities/"+cityID, nil)
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.successBody, respBody)

	// * Test 409 response & body
	// ID of city with foreign key conflict: Seattle
	cityID = "4c33e0bc-3d43-4e77-aed0-b7aff09bb689"

	req, err = http.NewRequest(http.MethodDelete, suite.server.URL+"/locations/cities/"+cityID, nil)
	require.NoError(t, err)

	req.AddCookie(&http.Cookie{Name: "bearer-token", Value: suite.authTokens.admin.Token})

	resp, err = http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusConflict, resp.StatusCode)

	respBody, err = io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	assertjson.Equal(t, suite.errBody, respBody)
}
