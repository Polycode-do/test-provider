package auth

// `LoginRequest` is the request body for the login endpoint.
// @property {string} Username - The username of the user.
// @property {string} Password - The password of the user.
// @property {string} GrantType - This is the type of grant you are requesting. Will always be `implicit`.
type LoginRequest struct {
	Username  string `json:"identity"`
	Password  string `json:"secret"`
	GrantType string `json:"grantType"`
}