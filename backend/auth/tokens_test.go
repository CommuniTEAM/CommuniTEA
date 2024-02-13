package auth_test

import (
	"testing"

	"github.com/CommuniTEAM/CommuniTEA/auth"
	"github.com/stretchr/testify/require"
)

func TestNewAuthenticator(t *testing.T) {
	auth, err := auth.NewAuthenticator()
	require.NoError(t, err)

	if err != nil {
		require.Nil(t, auth)
	}
	require.NotNil(t, auth)
}
