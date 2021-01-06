package util

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// GetProjectRoot get project root path
func GetProjectRoot() string {
	return os.Getenv("GOPROJECTROOT")
}

// RetrieveCredentials retrieves username and password from yaml config file
func RetrieveCredentials() (*Credentials, error) {
	f, err := os.Open(GetProjectRoot() + "/cred.sandbox.yml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var c *Credentials
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// PostRequest generic http post request
func PostRequest(url string, authInfo *AuthInfo, postData map[string]interface{}) (*http.Response, error) {
	postBody, err := json.Marshal(postData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// add authentication information to header if needed
	if authInfo != nil {
		req.Header.Add("Authorization", "Bearer "+authInfo.AccessToken)
	}

	return http.DefaultClient.Do(req)
}

// GetRequest generic http get request
func GetRequest(url string, authInfo *AuthInfo) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// add authentication information to header if needed
	if authInfo != nil {
		req.Header.Add("Authorization", "Bearer "+authInfo.AccessToken)
	}

	return http.DefaultClient.Do(req)
}

// PrintStruct prints out a struct in a json human readable format
func PrintStruct(s interface{}) {
	m, err := json.MarshalIndent(s, "", "	")
	if err == nil {
		log.Println(string(m))
	}
}
