package auth

// `Credentials` is a struct that contains the credentials of a user.
// @property {string} Username - The username of the user you want to authenticate.
// @property {string} Password - The password for the user.
type Credentials struct {
	Username string
	Password string
}
