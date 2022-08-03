package auth

import "time"

type AuthResponse struct {
	Metadata struct{} `json:"metadata"`
	Data     AuthData `json:"data"`
}

type AuthData struct {
	AccessToken string    `json:"accessToken"`
	TokenType   string    `json:"tokenType"`
	ExpiresAt   time.Time `json:"expiresAt"`
}
