package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"log"
	"time"

	db "github.com/CommuniTEAM/CommuniTEA/db/sqlc"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Authenticator struct {
	signingKey *rsa.PrivateKey
}

type TokenCookie struct {
	Token string `cookie:"bearer-token,httponly,secure,samesite=strict,path=/,max-age:3600" json:"access_token"`
}

type TokenData struct {
	TokenCookie
	TokenType string           `json:"token_type"`
	ExpiresIn int              `json:"expires_in"`
	ID        uuid.UUID        `json:"id"`
	Role      string           `json:"role"`
	Username  string           `json:"username"`
	FirstName string           `json:"first_name,omitempty"`
	LastName  string           `json:"last_name,omitempty"`
	Email     string           `json:"email,omitempty"`
	Location  db.LocationsCity `json:"location"`
}

// ValidateJWT takes a signed JWT, verifies it against the key, then parses
// the enclosed data and returns it. Returns nil if the token is invalid.
func (key *Authenticator) ValidateJWT(token string) *TokenData {
	verifiedToken, err := jwt.Parse([]byte(token), jwt.WithKey(jwa.RS256, key.signingKey))
	if err != nil {
		return nil
	}

	jsonbody, err := json.Marshal(verifiedToken.PrivateClaims())
	if err != nil {
		log.Printf("could not marshal token claims: %v", err)
		return nil
	}

	var userData TokenData
	err = json.Unmarshal(jsonbody, &userData)
	if err != nil {
		log.Printf("could not unmarshal token data: %v", err)
		return nil
	}

	return &userData
}

// GenerateNewJWT takes a struct of a validated user's information and appends
// a signed JWT with their information enclosed to the struct. The JWT is valid
// for one hour.
func (key *Authenticator) GenerateNewJWT(tokenData *TokenData, expired bool) (*TokenData, error) {
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
		jwtBuilder.Claim(`email`, tokenData.Email)
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
	signedToken, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, key.signingKey))
	if err != nil {
		return nil, fmt.Errorf("could not sign token: %w", err)
	}

	tokenData.Token = string(signedToken)
	tokenData.TokenType = "bearer"
	return tokenData, nil
}

// NewAuthenticator instantiates a new authenticator with a randomly generated,
// cryptographically secure RSA private key.
func NewAuthenticator() (*Authenticator, error) {
	const bits = 2048

	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, fmt.Errorf("could not generate signing key: %w", err)
	}

	return &Authenticator{signingKey: key}, nil
}
