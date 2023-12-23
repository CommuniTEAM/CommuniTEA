package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var signingKey = generateSigningKey() //nolint: gochecknoglobals // key needs to be randomly generated once at app start

type TokenData struct {
	Token     string    `cookie:"auth_bearer,httponly,secure,samesite=strict,path=/,max-age:3600" json:"access_token"`
	TokenType string    `json:"token_type"`
	ExpiresIn int       `json:"expires_in"`
	ID        uuid.UUID `json:"id"`
	Role      string    `json:"role"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Location  uuid.UUID `json:"location"`
}

// GenerateNewJWT takes a struct of a validated user's information and appends a signed JWT
// with their information enclosed to the struct. The JWT is valid for one hour.
func GenerateNewJWT(tokenData *TokenData) (*TokenData, error) {
	// Instantiate JWT builder
	jwtBuilder := jwt.NewBuilder()

	// Modify the JWT with the user's information (and issue/expiry time)
	jwtBuilder.Issuer("communitea.life")
	jwtBuilder.Claim(`id`, tokenData.ID)
	jwtBuilder.Claim(`role`, tokenData.Role)
	jwtBuilder.Claim(`username`, tokenData.Username)
	jwtBuilder.Claim(`first_name`, tokenData.FirstName)
	jwtBuilder.Claim(`last_name`, tokenData.LastName)
	jwtBuilder.Claim(`location`, tokenData.Location)
	jwtBuilder.IssuedAt(time.Now())
	jwtBuilder.Expiration(time.Now().Add(time.Hour))

	// Create the JWT
	token, err := jwtBuilder.Build()
	if err != nil {
		return nil, fmt.Errorf("could not build new token: %w", err)
	}

	// Sign the JWT with the .env signing key and a secure algorithm
	signedToken, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, signingKey))
	if err != nil {
		return nil, fmt.Errorf("could not sign token: %w", err)
	}

	tokenData.Token = string(signedToken)

	return tokenData, nil
}

func generateSigningKey() *rsa.PrivateKey {
	const bits = 2048

	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(fmt.Errorf("could not generate signing key: %w", err))
	}
	return key
}
