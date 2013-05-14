package firebase

import "testing"

func TestNewFirebase(t *testing.T) {
	// Create new firebase instance
	firebaseRoot := New("http://example.com")

	if firebaseRoot.base_uri != "http://example.com" {
		t.Errorf("Expected base_uri: %q, Actual: %q", "http://example.com", firebaseRoot.base_uri)
	}

	// Add auth token
	firebaseRoot.auth_token = "auth"
	if firebaseRoot.auth_token != "auth" {
		t.Errorf("Expected auth_token: %q, Actual: %q", "auth", firebaseRoot.auth_token)
	}
}

type Message struct {
	Name string
	Body string
}

func TestSetObject(t *testing.T) {

	// Create new firebase instance
	firebaseRoot := New("http://example.com")
	msg := Message{"testing", "1..2..3"}

	_, err := firebaseRoot.set("/users", msg)
	if err != nil {
		t.Error("Error seting object")
	}
}
