// Package firebase implements the functions to manipulate Firebase
package firebase

// Returns
func New(base_uri, auth_token string) firebase {
	return firebase{base_uri, auth_token}
}

// firebase
type firebase struct {
	base_uri   string
	auth_token string
}
