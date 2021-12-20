package main

import (
	"encoding/json"
	"net/http"
)

type Generator interface {
	GenerateToken(map[string]interface{}) (string, error)
	Parse(string) (map[string]interface{}, error)
}

type OAuth2 struct {
	Generator  Generator
	RefreshExp int64
	AccessExp  int64
}

func NewOAuth2(generator Generator, refreshExp int64, accessExp int64) *OAuth2 {
	return &OAuth2{generator, refreshExp, accessExp}
}

// Generate refresh and access token
// Responds json with the access and refresh tokens
func (oauth2 *OAuth2) AllTokens(tokensData map[string]interface{}, w http.ResponseWriter) error {
	tokensData["exp"] = oauth2.RefreshExp
	refresh, err := oauth2.Generator.GenerateToken(tokensData)
	if err != nil {
		return err
	}

	tokensData["exp"] = oauth2.AccessExp
	access, err := oauth2.Generator.GenerateToken(tokensData)
	if err != nil {
		return err
	}

	resp, err := json.Marshal(map[string]string{"access": access, "refresh": refresh})
	if err != nil {
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)

	return nil
}

func (oauth2 *OAuth2) AccessTokens(refreshToken string, w http.ResponseWriter) error {
	claims, err := oauth2.Generator.Parse(refreshToken)
	if err != nil {
		return err
	}

	claims["exp"] = oauth2.AccessExp

	token, err := oauth2.Generator.GenerateToken(claims)
	if err != nil {
		return err
	}

	resp, err := json.Marshal(map[string]string{"access": token})
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)

	return nil
}

func (oauth2 *OAuth2) Parse(token string) (map[string]interface{}, error) {
	return oauth2.Generator.Parse(token)
}
