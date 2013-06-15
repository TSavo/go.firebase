// Package firebase implements the functions to manipulate Firebase
package firebase

import (
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

// Firebase struct
type firebaseRoot struct {
	base_uri   string
	auth_token string
}

/*
Class Methods
*/

// Construct a firebase type
func New(url string) *firebaseRoot {
	return &firebaseRoot{base_uri: url}
}

/*
Instance Methods
Firebase Package - REST Interface
*/

// Writes and returns the data to the firebase endpoint
// Example: firebase.Set('/users/info', [struct]) => []byte(data), error
func (f *firebaseRoot) set(path string, v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	json := bytes.NewBuffer(b)

	// TODO: Construct from Base URI
	// create a put request
	req, err := http.NewRequest("PUT", "http://example.com", json)
	if err != nil {
		return nil, err
	}

	// send JSON to firebase
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad HTTP Response: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// // Returns the data at a path
// func Get(path string) string {

// }

// // Writes the data, returns the key name of the data added
// // Example: firebase.Push('users', [struct]) => {'id':'-INOQPH-aV_psbk3ZXEX'}
// func Push(path string) string {

// }

// // Deletes the data and returns true or false
// func Delete(path string) bool {

// }

// // Updates the data at path and returns the data. Does not delete omitted children.
// func Update(path string) string {

// }
