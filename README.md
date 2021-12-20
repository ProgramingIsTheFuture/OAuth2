# OAUTH2

OAuth2 simple to use implementation in golang

- Basic functionalities:
	- Generate JWT tokens for authentication
	- Interacts with your database to store the valid tokens 
	- Verify if the token given is valid or invalid
	- Error handlers for OAuth2 lib

```golang
// first we need to initialize the generator

generator := oauth2.NewOAuthGenerator()

// We can create the oauth passing the generator and the refresh exp time and access exp time
oauth := oauth2.NewOAuth2(
	generator,
	time.Now().Add(time.Hour*24*7).Unix(),
	time.Now().Add(time.Minute*1).Unix(),
)

// Now we have access to three methods

// Generate the refresh and the access token
// First parameter is the claims of the tokens - there is no need to pass "exp"
// and the responseWritter to respond with json
// Returns an error
oauth.AllTokens(map[string]interface{}, http.ResponseWritter)

// Generate the access token 
// First parameter is the refresh token
// Responds json 
// Returns an error
oauth.AccessTokens(string, http.ResponseWriter)

// Parse a token given
// Receives the token
// Returns the claims of that specific token and an error
oauth.Parse(string)

```

```json
// Json response types

{
	"refresh": "<Token>",
	"access": "<Token>",
}

// or

{
	"access": "<Token>",
}
```

