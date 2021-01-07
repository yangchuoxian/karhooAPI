package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"karhooAPIs.com/util"
)

func main() {
	// ****************************** get access token
	authInfo, err := getAccessToken()
	if err != nil {
		log.Fatal(err)
	}
	// ****************************** refresh access token if necessary
	err = refreshAccessTokenIfExpired(authInfo)
	if err != nil {
		log.Fatal(err)
	}
	// ****************************** register karhoo webhook
	err = registerWebhook(authInfo, "http://karhoo-webhooks.piizu.com/webhook", util.WebhookSecretKey)
	if err != nil {
		log.Fatal(err)
	}

	subscriptions, err := getRegisteredWebhookURLs(authInfo)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("********************* REGISTERED WEBHOOK: ")
	util.PrintStruct(subscriptions)
	// ****************************** request quotes
	origin := util.Geolocation{
		Latitude:       "50.037933",
		Longitude:      "8.562152",
		DisplayAddress: "Frankfurt Airport",
	}
	destination := util.Geolocation{
		Latitude:       "51.037933",
		Longitude:      "8.910231",
		DisplayAddress: "Some place nearby",
	}
	quotesList, err := getQuotes(authInfo, origin, destination, "")
	if err != nil {
		log.Fatal(err)
	}
	// ****************************** retrieve quote list
	retrievedQuoteList, err := retrieveQuoteList(authInfo, quotesList.ID)
	if err != nil {
		log.Fatal(err)
	}
	if len(retrievedQuoteList.Quotes) == 0 {
		log.Fatal("failed to request quotes")
	}
	log.Println("********************* RETRIEVED QUOTE LIST: ")
	util.PrintStruct(retrievedQuoteList)
	// ****************************** select the first quote from quote list and make a booking
	bookingResults, err := bookATrip(authInfo, map[string]interface{}{
		"quote_id": retrievedQuoteList.Quotes[0].ID,
		"passengers": map[string]interface{}{
			"passenger_details": []map[string]interface{}{{
				"first_name":   "Chuoxian",
				"last_name":    "Yang",
				"phone_number": "+15005550006",
			}},
			"luggage": map[string]interface{}{
				"total": 1,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("********************* REQUESTED A BOOKING: ")
	util.PrintStruct(bookingResults)
	// ****************************** get booking details of previous book request
	bookingDetails, err := getBookingDetails(authInfo, bookingResults.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("********************* GOT BOOKING DETAILS: ")
	util.PrintStruct(bookingDetails)
	// ****************************** cancel booking
	err = cancelBooking(authInfo, bookingDetails.ID, util.CancelBookingReasons[0])
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("********************* BOOKING %s CANCELED SUCCESSFULLY", bookingDetails.ID)
}

func getAccessToken() (*util.AuthInfo, error) {
	cred, err := util.RetrieveCredentials()
	if err != nil {
		return nil, err
	}
	res, err := util.PostRequest(util.GetAccessTokenURL, nil, map[string]interface{}{
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
		res, err := util.PostRequest(util.RefreshAccessTokenURL, nil, map[string]interface{}{
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

func getQuotes(a *util.AuthInfo, origin util.Geolocation, destination util.Geolocation, pickupTime string) (*util.QuotesList, error) {
	res, err := util.PostRequest(util.GetQuotesURL, a, map[string]interface{}{
		"origin":               origin,
		"destination":          destination,
		"local_time_of_pickup": "",
	})
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()
	if res.StatusCode == http.StatusCreated {
		var quotesList *util.QuotesList
		err = decoder.Decode(&quotesList)
		if err != nil {
			return nil, err
		}
		return quotesList, nil
	}
	// get quote failed with code and error message
	var e *util.ErrorInfo
	err = decoder.Decode(&e)
	if err != nil {
		return nil, err
	}
	return nil, errors.New(e.Message)
}

func retrieveQuoteList(a *util.AuthInfo, quoteListID string) (*util.QuotesList, error) {
	res, err := util.GetRequest(util.RetrieveQuoteList+quoteListID, a)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		var quotesList *util.QuotesList
		err = decoder.Decode(&quotesList)
		if err != nil {
			return nil, err
		}
		return quotesList, nil
	}
	// retrieve quote list failed with code and error message
	var e *util.ErrorInfo
	err = decoder.Decode(&e)
	if err != nil {
		return nil, err
	}
	return nil, errors.New(e.Message)
}

func bookATrip(a *util.AuthInfo, params map[string]interface{}) (*util.BookingDetails, error) {
	res, err := util.PostRequest(util.BookingURL, a, params)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()
	if res.StatusCode == http.StatusCreated {
		var bookingResponse *util.BookingDetails
		err = decoder.Decode(&bookingResponse)
		if err != nil {
			return nil, err
		}
		return bookingResponse, nil
	}
	// book trip failed with code and error message
	var e *util.ErrorInfo
	err = decoder.Decode(&e)
	if err != nil {
		return nil, err
	}
	return nil, errors.New(e.Message)
}

func getBookingDetails(a *util.AuthInfo, bookingID string) (*util.BookingDetails, error) {
	res, err := util.GetRequest(util.GetBookingDetailsURL+bookingID, a)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		var bookingDetails *util.BookingDetails
		err = decoder.Decode(&bookingDetails)
		if err != nil {
			return nil, err
		}
		return bookingDetails, nil
	}
	// get booking details failed with code and error message
	var e *util.ErrorInfo
	err = decoder.Decode(&e)
	if err != nil {
		return nil, err
	}
	return nil, errors.New(e.Message)
}

func cancelBooking(a *util.AuthInfo, bookingID, cancelReason string) error {
	cancelBookingURL := fmt.Sprintf(util.CancelBookingURL, bookingID)
	res, err := util.PostRequest(cancelBookingURL, a, map[string]interface{}{
		"reason": cancelReason,
	})
	if err != nil {
		return err
	}
	if res.StatusCode == http.StatusNoContent {
		return nil
	}
	// cancel booking failed with code and error message
	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()
	var e *util.ErrorInfo
	err = decoder.Decode(&e)
	if err != nil {
		return err
	}
	return errors.New(e.Message)
}

func registerWebhook(a *util.AuthInfo, url, sharedSecret string) error {
	res, err := util.PostRequest(util.RegisterWebhookURL, a, map[string]interface{}{
		"url":           url,
		"shared_secret": sharedSecret,
	})
	if err != nil {
		return err
	}
	if res.StatusCode == http.StatusCreated {
		return nil
	}
	// register webhook failed with code and error message
	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()
	var e *util.ErrorInfo
	err = decoder.Decode(&e)
	if err != nil {
		return err
	}
	return errors.New(e.Message)
}

func getRegisteredWebhookURLs(a *util.AuthInfo) (*util.WebhookSubscription, error) {
	res, err := util.GetRequest(util.ReturnSubscriptionURL, a)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		var sub *util.WebhookSubscription
		err = decoder.Decode(&sub)
		if err != nil {
			return nil, err
		}
		return sub, nil
	}
	// get registered webhook urls failed with code and error message
	var e *util.ErrorInfo
	err = decoder.Decode(&e)
	if err != nil {
		return nil, err
	}
	return nil, errors.New(e.Message)
}
