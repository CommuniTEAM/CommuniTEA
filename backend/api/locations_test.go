package api_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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
	t.Logf("%v", string(body))
}
