package auth

// `LoginResponse` is the response body for the login endpoint.
// @property {string} AccessToken - The access token that you'll use to make requests to the API.
// @property {string} TokenType - This is the type of token that is returned. Will always be `bearer`.
// @property {string} ExpiresAt - The time at which the access token will expire.
type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresAt   string `json:"expiresAt"`
}
