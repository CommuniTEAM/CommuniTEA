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

type EventsTestSuite struct {
	TestSuite
}

func TestEventsTestSuite(t *testing.T) {
	suite.Run(t, new(EventsTestSuite))
}

func (suite *EventsTestSuite) TestGetAllEvents() {
	t := suite.T()

	eventId := "e6e8e3e3-3e3e-4e3e-8e3e-3e3e3e3e3e3e"

	req, err := http.NewRequest("GET", suite.server.URL+"/events/"+eventId, nil)
	require.NoError(t, err)

	resp, err := http.DefaultTransport.RoundTrip(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	expectedBody, err := os.ReadFile("_testdata/events/get_event.json")
	require.NoError(t, err)

	assertjson.Equal(t, expectedBody, body)
}
