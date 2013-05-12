package firebase

import "testing"

func TestNewFirebase(t *testing.T) {
	base_uri := "http://example.com"
	auth := "auth_token"
	firebase := New(base_uri, auth)
	if firebase.base_uri != base_uri {
		t.Errorf("Expected base_uri: %q, Actual: %q", base_uri, firebase.base_uri)
	}
	if firebase.auth_token != auth {
		t.Errorf("Expected auth_token: %q, Actual: %q", auth, firebase.auth_token)
	}
}
