package firebase

import (
	"testing"
)

const firebase_url = "https://go-firebase.firebaseio.com/"

func TestNewFirebase(t *testing.T) {
	// Create new firebase instance
	firebaseRoot := New(firebase_url)

	if firebaseRoot.base_url != firebase_url {
		t.Errorf("Expected base_url: %q, Actual: %q", firebase_url, firebaseRoot.base_url)
	}

	// Add auth token
	firebaseRoot.auth_token = "auth"
	if firebaseRoot.auth_token != "auth" {
		t.Errorf("Expected auth_token: %q, Actual: %q", "auth", firebaseRoot.auth_token)
	}
}

func TestBuildURL(t *testing.T) {
	firebaseRoot := New(firebase_url)
	url := firebaseRoot.BuildURL("/messages")
	if url != firebase_url+"messages" {
		t.Errorf("Expected URL: %q, Actual: %q", firebase_url+"messages", url)
	}
}

type Message struct {
	Name string
	Body string
}

func TestSetObject(t *testing.T) {
	// Create new firebase instance
	firebaseRoot := New(firebase_url)
	msg := Message{"testing", "1..2..3"}

	_, err := firebaseRoot.Set("/messages/1", msg)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestGetObject(t *testing.T) {
	firebaseRoot := New(firebase_url)
	_, err := firebaseRoot.Get("/messages")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestPushObject(t *testing.T) {
	firebaseRoot := New(firebase_url)
	msg := Message{"testing", "1..2..3"}
	_, err := firebaseRoot.Push("/messages", msg)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestUpdateObject(t *testing.T) {
	firebaseRoot := New(firebase_url)
	msg := Message{"testing", "1..2..3"}
	_, err := firebaseRoot.Update("/messages", msg)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestDeleteObject(t *testing.T) {
	firebaseRoot := New(firebase_url)
	_, err := firebaseRoot.Delete("/messages")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
