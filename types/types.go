/*
Package types is to defined data schema in communication
*/
package types

// RegisterData is the data received from user who is registerring
type RegisterData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthDataResponse is the response to users who succeed sign up or sign in
type AuthDataResponse struct {
	AccessToken string `json:"accessToken"`
}

// TokenPayload is the data contained in tokens
type TokenPayload struct {
	Uid         string
	TokenType   string
	Exp         float64
	RefeshToken string
}
