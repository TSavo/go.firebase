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
	if url != "https://go-firebase.firebaseio.com/messages" {
		t.Errorf("Expected URL: %q, Actual: %q", "https://go-firebase.firebaseio.com/messages", url)
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

	body, err := firebaseRoot.Set("/messages/1", msg)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("%q", body)
}

func TestGetObject(t *testing.T) {
	firebaseRoot := New(firebase_url)
	body, err := firebaseRoot.Get("/messages/1")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("%q", body)
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
