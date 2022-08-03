package auth

type AuthRequest struct {
	Username  string `json:"identity"`
	Password  string `json:"secret"`
	GrantType string `json:"grantType"`
}