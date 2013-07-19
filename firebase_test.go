package firebase

import (
	"testing"
	"time"
)

const firebase_url = "https://go-firebase-test.firebaseio.com/"

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
	firebaseRoot := New("https://your-firebase.firebaseio.com/")
	url := firebaseRoot.BuildURL("/messages")
	if url != "https://your-firebase.firebaseio.com/messages.json" {
		t.Errorf("Expected URL: %q, Actual: %q", "https://your-firebase.firebaseio.com/messages.json", url)
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

	body, err := firebaseRoot.Set("1", msg)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("%q", body)

	time.Sleep(time.Second)
}

func TestGetObject(t *testing.T) {
	firebaseRoot := New(firebase_url)
	body, err := firebaseRoot.Get("1")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("%q", body)

	time.Sleep(time.Second)
}

func TestPushObject(t *testing.T) {
	firebaseRoot := New(firebase_url)
	msg := Message{"testing", "1..2..3"}
	body, err := firebaseRoot.Push("/", msg)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("%q", body)

	time.Sleep(time.Second)
}

func TestUpdateObject(t *testing.T) {
	firebaseRoot := New(firebase_url)
	msg := Message{"testing", "1..2..3"}
	body, err := firebaseRoot.Update("1", msg)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("%q", body)

	time.Sleep(time.Second)
}

func TestDeleteObject(t *testing.T) {
	firebaseRoot := New(firebase_url)
	body, err := firebaseRoot.Delete("1")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("%t", body)

	time.Sleep(time.Second)
}
