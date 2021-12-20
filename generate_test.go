package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	generator := NewOAuthGenerator()

	auth := NewOAuth2(generator, time.Now().Add(time.Hour*24*7).Unix(), time.Now().Add(time.Minute*1).Unix())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := auth.AllTokens(map[string]interface{}{"username": "aaa"}, w)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%s", err.Error())))
		}

		token, _ := auth.Generator.GenerateToken(map[string]interface{}{"username": "aaa", "exp": auth.AccessExp})
		fmt.Println(token)
		fmt.Println(auth.Parse(token))

	})

	http.ListenAndServe(":8000", nil)
}
