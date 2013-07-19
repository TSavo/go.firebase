package firebase

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
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
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{""}`)
	}))
	defer ts.Close()

	// Create new firebase instance
	firebaseRoot := New(ts.URL)
	msg := Message{"testing", "1..2..3"}

	body, err := firebaseRoot.Set("1", msg)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("%q", body)
}

func TestGetObject(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{""}`)
	}))
	defer ts.Close()

	firebaseRoot := New(ts.URL)
	body, err := firebaseRoot.Get("1")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("%q", body)
}

func TestPushObject(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{""}`)
	}))
	defer ts.Close()

	firebaseRoot := New(ts.URL)
	msg := Message{"testing", "1..2..3"}
	body, err := firebaseRoot.Push("/", msg)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("%q", body)
}

func TestUpdateObject(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{""}`)
	}))
	defer ts.Close()

	firebaseRoot := New(ts.URL)
	msg := Message{"testing", "1..2..3"}
	body, err := firebaseRoot.Update("1", msg)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("%q", body)
}

func TestDeleteObject(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{""}`)
	}))
	defer ts.Close()

	firebaseRoot := New(firebase_url)
	body, err := firebaseRoot.Delete("1")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("%t", body)
}
