package util

import "time"

// Credentials username/password retrieved from yaml config file
type Credentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// ErrorInfo generic failed http response
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details []struct {
		Message string `json:"message"`
		Detail  string `json:"detail"`
	} `json:"details"`
}

// AuthInfo response to get access token
type AuthInfo struct {
	AccessToken    string `json:"access_token"`
	ExpiresIn      int    `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	ExpirationTime time.Time
}

// RefreshInfo response to refresh access token
type RefreshInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// Geolocation geolocation with latitude/longitude coordinates
type Geolocation struct {
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	DisplayAddress string `json:"display_address"`
}

// WebhookSubscription the registered webhook url
type WebhookSubscription struct {
	URL string `json:"url"`
}

// QuotesList quotes
type QuotesList struct {
	ID           string `json:"id"`
	Availability struct {
		Vehicles struct {
			Classes []string `json:"classes"`
			Tags    []string `json:"tags"`
			Types   []string `json:"types"`
		} `json:"vehicles"`
	} `json:"availability"`
	Quotes []struct {
		ID    string `json:"id"`
		Price struct {
			CurrencyCode string `json:"currency_code"`
			High         int    `json:"high"`
			Low          int    `json:"low"`
			Net          struct {
				High int `json:"high"`
				Low  int `json:"low"`
			} `json:"net"`
		} `json:"price"`
		PickUpType string `json:"pick_up_type"`
		QuoteType  string `json:"quote_type"`
		Source     string `json:"source"`
		Fleet      struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Rating      struct {
				Count int `json:"count"`
				Score int `json:"score"`
			} `json:"rating"`
			LogoURL            string   `json:"logo_url"`
			TermsConditionsURL string   `json:"terms_conditions_url"`
			PhoneNumber        string   `json:"phone_number"`
			Capabilities       []string `json:"capabilities"`
		} `json:"fleet"`
		Vehicle struct {
			QTA struct {
				HighMinutes int `json:"high_minutes"`
				LowMinutes  int `json:"low_minutes"`
			} `json:"qta"`
			Class             string   `json:"class"`
			Type              string   `json:"type"`
			PassengerCapacity int      `json:"passenger_capacity"`
			LuggageCapacity   int      `json:"luggage_capacity"`
			Tags              []string `json:"tags"`
		} `json:"vehicle"`
	} `json:"quotes"`
	Status   string `json:"status"`
	Validity int    `json:"validity"`
}

// BookingDetails details of a booking
type BookingDetails struct {
	ID         string `json:"id"`
	Passengers struct {
		AdditionalPassengers int `json:"additional_passengers"`
		PassengerDetails     []struct {
			FirstName   string `json:"first_name"`
			LastName    string `json:"last_name"`
			Email       string `json:"email"`
			PhoneNumber string `json:"phone_number"`
			Locale      string `json:"locale"`
		} `json:"passenger_details"`
		Luggage struct {
			Total int `json:"total"`
		} `json:"luggage"`
	} `json:"passengers"`
	PartnerTravellerID string `json:"partner_traveller_id"`
	Status             string `json:"status"`
	StateDetails       string `json:"state_details"`
	Origin             struct {
		DisplayAddress string `json:"display_address"`
		Position       struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"position"`
		PlaceID  string `json:"place_id"`
		PoiType  string `json:"poi_type"`
		Timezone string `json:"timezone"`
	} `json:"origin"`
	Destination struct {
		DisplayAddress string `json:"display_address"`
		Position       struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"position"`
		PlaceID  string `json:"place_id"`
		PoiType  string `json:"poi_type"`
		Timezone string `json:"timezone"`
	} `json:"destination"`
	DateScheduled time.Time `json:"date_scheduled"`
	Quote         struct {
		Type            string `json:"type"`
		Total           int    `json:"total"`
		Currency        string `json:"currency"`
		GratuityPercent int    `json:"gratuity_percent"`
		Breakdown       []struct {
			Value       int    `json:"value"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"breakdown"`
		VehicleClass      string `json:"vehicle_class"`
		QtaHighMinutes    int    `json:"qta_high_minutes"`
		QtaLowMinutes     int    `json:"qta_low_minutes"`
		VehicleAttributes struct {
			PassengerCapacity int  `json:"passenger_capacity"`
			LuggageCapacity   int  `json:"luggage_capacity"`
			Hybrid            bool `json:"hybrid"`
			Electric          bool `json:"electric"`
			ChildSeat         bool `json:"child_seat"`
		} `json:"vehicle_attributes"`
		Source    string `json:"source"`
		HighPrice int    `json:"high_price"`
		LowPrice  int    `json:"low_price"`
	} `json:"quote"`
	Fare struct {
		Total           int    `json:"total"`
		Currency        string `json:"currency"`
		GratuityPercent int    `json:"gratuity_percent"`
		Breakdown       []struct {
			Value       int    `json:"value"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"breakdown"`
	} `json:"fare"`
	ExternalTripID string `json:"external_trip_id"`
	DisplayTripID  string `json:"display_trip_id"`
	FleetInfo      struct {
		FleetID            string `json:"fleet_id"`
		Name               string `json:"name"`
		LogoURL            string `json:"logo_url"`
		Description        string `json:"description"`
		PhoneNumber        string `json:"phone_number"`
		TermsConditionsURL string `json:"terms_conditions_url"`
		Email              string `json:"email"`
	} `json:"fleet_info"`
	Vehicle struct {
		VehicleClass        string `json:"vehicle_class"`
		Description         string `json:"description"`
		VehicleLicensePlate string `json:"vehicle_license_plate"`
		Driver              struct {
			FirstName     string `json:"first_name"`
			LastName      string `json:"last_name"`
			PhoneNumber   string `json:"phone_number"`
			PhotoURL      string `json:"photo_url"`
			LicenseNumber string `json:"license_number"`
		} `json:"driver"`
		Attributes struct {
			PassengerCapacity int  `json:"passenger_capacity"`
			LuggageCapacity   int  `json:"luggage_capacity"`
			Hybrid            bool `json:"hybrid"`
			Electric          bool `json:"electric"`
			ChildSeat         bool `json:"child_seat"`
		} `json:"attributes"`
	} `json:"vehicle"`
	PartnerTripID string `json:"partner_trip_id"`
	Comments      string `json:"comments"`
	FlightNumber  string `json:"flight_number"`
	TrainNumber   string `json:"train_number"`
	DateBooked    string `json:"date_booked"`
	MeetingPoint  struct {
		Position struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"position"`
		Type         string `json:"type"`
		Instructions string `json:"instructions"`
		Note         string `json:"note"`
	} `json:"meeting_point"`
	Agent struct {
		UserID           string `json:"user_id"`
		UserName         string `json:"user_name"`
		OrganisationID   string `json:"organisation_id"`
		OrganisationName string `json:"organisation_name"`
	} `json:"agent"`
	CostCenterReference string `json:"cost_center_reference"`
	CancelledBy         struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		ID        string `json:"id"`
		Email     string `json:"email"`
	} `json:"cancelled_by"`
	FollowCode string `json:"follow_code"`
	Meta       struct {
		AdditionalProp string `json:"additionalProp"`
	} `json:"meta"`
	TrainTime time.Time `json:"train_time"`
}

// WebhookSecretKey shared secret when setting up webhook
const WebhookSecretKey = "Jal019lafj0192QtYbNzmMAsL"

const (
	// GetAccessTokenURL url to get access token
	GetAccessTokenURL = "https://rest.sandbox.karhoo.com/v1/auth/token"
	// RefreshAccessTokenURL url to refresh access token
	RefreshAccessTokenURL = "https://rest.sandbox.karhoo.com/v1/auth/refresh"
	// GetQuotesURL url to get quotes
	GetQuotesURL = "https://rest.sandbox.karhoo.com/v2/quotes/"
	// RetrieveQuoteList url to retrieve quote list
	RetrieveQuoteList = "https://rest.sandbox.karhoo.com/v2/quotes/"
	// BookingURL url to make a booking
	BookingURL = "https://rest.sandbox.karhoo.com/v1/bookings/"
	// GetBookingDetailsURL url to get booking details
	GetBookingDetailsURL = "https://rest.sandbox.karhoo.com/v1/bookings/"
	// CancelBookingURL url to cancel booking
	CancelBookingURL = "https://rest.sandbox.karhoo.com/v1/bookings/%s/cancel/"
	// RegisterWebhookURL url to register webhook
	RegisterWebhookURL = "https://rest.sandbox.karhoo.com/v1/webhooks/"
	// ReturnSubscriptionURL url to return current webhook subscription for user
	ReturnSubscriptionURL = "https://rest.sandbox.karhoo.com/v1/webhooks/"
)

// CancelBookingReasons reason strings to cancel booking
var CancelBookingReasons = []string{
	"OTHER_USER_REASON",
	"DRIVER_DIDNT_SHOW_UP",
	"ETA_TOO_LONG",
	"DRIVER_IS_LATE",
	"CAN_NOT_FIND_VEHICLE",
	"NOT_NEEDED_ANYMORE",
	"ASKED_BY_DRIVER_TO_CANCEL",
	"FOUND_BETTER_PRICE",
	"NOT_CLEAR_MEETING_INSTRUCTIONS",
	"COULD_NOT_CONTACT_CARRIER",
}
