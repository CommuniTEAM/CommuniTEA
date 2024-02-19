package auth_test

import (
	"testing"

	"github.com/CommuniTEAM/CommuniTEA/auth"
	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
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

func TestGenerateJWT(t *testing.T) {
	userData := auth.TokenData{
		ExpiresIn: 3600,
		ID:        uuid.MustParse("372bcfb3-6b1d-4925-9f3d-c5ec683a4294"),
		Role:      "user",
		Username:  "user",
		Location: db.LocationsCity{
			ID:    uuid.MustParse("4c33e0bc-3d43-4e77-aed0-b7aff09bb689"),
			Name:  "Seattle",
			State: "WA",
		}}

	authenticator, err := auth.NewAuthenticator()
	require.NoError(t, err)

	output, err := authenticator.GenerateNewJWT(&userData, false)
	require.NoError(t, err)

	require.Equal(t, "bearer", output.TokenType)
	require.Equal(t, "user", output.Username)

	validToken := output.Token

	output, err = authenticator.GenerateNewJWT(&userData, true)
	require.NoError(t, err)

	require.NotEqual(t, validToken, output.Token)
}

func TestValidateJWT(t *testing.T) {
	userData := auth.TokenData{
		ExpiresIn: 3600,
		ID:        uuid.MustParse("372bcfb3-6b1d-4925-9f3d-c5ec683a4294"),
		Role:      "user",
		Username:  "user",
		Location: db.LocationsCity{
			ID:    uuid.MustParse("4c33e0bc-3d43-4e77-aed0-b7aff09bb689"),
			Name:  "Seattle",
			State: "WA",
		}}

	authenticator, err := auth.NewAuthenticator()
	require.NoError(t, err)

	validToken, err := authenticator.GenerateNewJWT(&userData, false)
	require.NoError(t, err)

	output := authenticator.ValidateJWT(validToken.Token)
	require.NotNil(t, output)

	require.Equal(t, userData.Username, output.Username)

	invalidToken, err := authenticator.GenerateNewJWT(&userData, true)
	require.NoError(t, err)

	output = authenticator.ValidateJWT(invalidToken.Token)
	require.Nil(t, output)
}
