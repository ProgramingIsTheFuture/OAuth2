package main

import (
	"math/rand"

	"github.com/golang-jwt/jwt/v4"
)

type OAuthGenerator struct {
	secretKey string
}

// Generate a random string for secret key
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func NewOAuthGenerator() *OAuthGenerator {
	return &OAuthGenerator{secretKey: randomString(16)}
}

// Same as NewOAuthGenerator but you can define your custom secret key
func NewOAuthGeneratorWithParams(secret string) *OAuthGenerator {
	return &OAuthGenerator{secretKey: secret}
}

// GenerateAccessToken
// Must pass claims type map[string]interface{}
// Must contain a field: "iat", "nbf", "exp"
// Takes the claims and generate a new token with it
func (oa OAuthGenerator) GenerateToken(claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims(claims))
	return token.SignedString([]byte(oa.secretKey))
}

func (oa OAuthGenerator) Parse(refresh string) (map[string]interface{}, error) {

	var claims = jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refresh, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(oa.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}
