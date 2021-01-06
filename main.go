package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"karhooAPIs.com/util"
)

func main() {
	authInfo, err := getAccessToken()
	if err != nil {
		log.Fatal(err)
	}
	err = refreshAccessTokenIfExpired(authInfo)
	if err != nil {
		log.Fatal(err)
	}
}

func getAccessToken() (*util.AuthInfo, error) {
	cred, err := util.RetrieveCredentials()
	if err != nil {
		return nil, err
	}
	res, err := util.PostRequest(util.GetAccessTokenURL, map[string]interface{}{
		"username": cred.Username,
		"password": cred.Password,
	})
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()
	if res.StatusCode == http.StatusCreated {
		var a *util.AuthInfo
		err = decoder.Decode(&a)
		if err != nil {
			return nil, err
		}
		// calculate expiration time
		now := time.Now()
		a.ExpirationTime = now.Add(time.Second * time.Duration(a.ExpiresIn))
		return a, nil
	}
	// authentication failed with code and error message
	var e *util.ErrorInfo
	err = decoder.Decode(&e)
	if err != nil {
		return nil, err
	}
	return nil, errors.New(e.Message)
}

func refreshAccessTokenIfExpired(a *util.AuthInfo) error {
	// check if access token is already expired
	if a.ExpirationTime.Before(time.Now()) {
		res, err := util.PostRequest(util.RefreshAccessTokenURL, map[string]interface{}{
			"refresh_token": a.RefreshToken,
		})
		if err != nil {
			return err
		}
		decoder := json.NewDecoder(res.Body)
		defer res.Body.Close()
		if res.StatusCode == http.StatusCreated {
			var r *util.RefreshInfo
			err = decoder.Decode(&r)
			if err != nil {
				return err
			}
			// refresh access token succeeded, update access token and expiration time
			a.AccessToken = r.AccessToken
			a.ExpiresIn = r.ExpiresIn
			a.ExpirationTime = time.Now().Add(time.Second * time.Duration(r.ExpiresIn))
			return nil
		}
		// refresh access token failed with code and error message
		var e *util.ErrorInfo
		err = decoder.Decode(&e)
		if err != nil {
			return err
		}
		return errors.New(e.Message)
	}
	return nil
}
