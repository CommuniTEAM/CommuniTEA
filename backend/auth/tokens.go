package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var signingKey = generateSigningKey() //nolint: gochecknoglobals // key needs to be randomly generated once at app start

type TokenCookie struct {
	Token string `cookie:"bearer-token,httponly,secure,samesite=strict,path=/,max-age:3600" json:"access_token"`
}

type TokenData struct {
	TokenCookie
	TokenType string    `json:"token_type"`
	ExpiresIn int       `json:"expires_in"`
	ID        uuid.UUID `json:"id"`
	Role      string    `json:"role"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Location  uuid.UUID `json:"location"`
}

// ValidateJWT takes a signed JWT, verifies it against the key, then parses
// the enclosed data and returns it. Returns nil if the token is invalid.
func ValidateJWT(token string) map[string]interface{} {
	verifiedToken, err := jwt.Parse([]byte(token), jwt.WithKey(jwa.RS256, signingKey))
	if err != nil {
		return nil
	}

	return verifiedToken.PrivateClaims()
}

// GenerateNewJWT takes a struct of a validated user's information and appends
// a signed JWT with their information enclosed to the struct. The JWT is valid
// for one hour.
func GenerateNewJWT(tokenData *TokenData, expired bool) (*TokenData, error) {
	// Instantiate JWT builder
	jwtBuilder := jwt.NewBuilder()

	jwtBuilder.Issuer("communitea.life")

	if !expired {
		// Modify the JWT with the user's information (and issue/expiry time)
		jwtBuilder.Claim(`id`, tokenData.ID)
		jwtBuilder.Claim(`role`, tokenData.Role)
		jwtBuilder.Claim(`username`, tokenData.Username)
		jwtBuilder.Claim(`first_name`, tokenData.FirstName)
		jwtBuilder.Claim(`last_name`, tokenData.LastName)
		jwtBuilder.Claim(`location`, tokenData.Location)
		jwtBuilder.IssuedAt(time.Now())
		jwtBuilder.Expiration(time.Now().Add(time.Hour))
	} else {
		jwtBuilder.Expiration(time.Now().Add(-time.Hour))
	}

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

// generateSigningKey creates a new cryptographically secure RSA private key.
// It will shut down the app in the event of an error so that the program
// does not continue unsecured.
func generateSigningKey() *rsa.PrivateKey {
	const bits = 2048

	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Fatal(fmt.Errorf("could not generate signing key: %w", err))
	}

	return key
}
