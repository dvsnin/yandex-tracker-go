package tracker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BasicUsers []BasicUser

// User
// Basic user structure in Yandex.Tracker
type BasicUser struct {
	Self    string
	ID      string
	Display string
}

// Id
// Get user id
func (u *BasicUser) Id() string {
	if u != nil {
		return u.ID
	}

	return ""
}

// Name
// Get username
func (u *BasicUser) Name() string {
	if u != nil {
		return u.Display
	}

	return ""
}

type Users []User

// User structure in Yandex.Tracker
// https://cloud.yandex.ru/en/docs/tracker/get-user-info
type User struct {
	// Address of the API resource with information about the user account
	Self string `json:"self"`

	// Unique ID of the user Tracker account
	UID int `json:"uid"`

	// User's login
	Login string `json:"login"`

	// Unique ID of the user Tracker account
	TrackerUid int `json:"trackerUid"`

	// Unique ID of the user account in the Yandex 360 for Business organization and Yandex ID
	PassportUid int `json:"passportUid"`

	// User unique ID in Yandex Cloud Organization
	CloudUid string `json:"cloudUid"`

	// User's first name
	FirstName string `json:"firstName"`

	// User's last name
	LastName string `json:"lastName"`

	// Displayed user name
	Display string `json:"display"`

	// User email address
	Email string `json:"email"`

	// Service parameter
	External bool `json:"external"`

	// Flag indicating whether the user has full access to Tracker:
	// true: Full access
	// false: Read-only access
	HasLicense bool `json:"hasLicense"`

	// User status in the organization:
	// true: User is dismissed
	// false: User is a current employee
	Dismissed bool `json:"dismissed"`

	// Service parameter
	UseNewFilters bool `json:"useNewFilters"`

	// Flag indicating whether user notifications are forced disabled:
	// true: Disabled
	// false: Enabled
	DisableNotifications bool `json:"disableNotifications"`

	// Date and time of the user's first authentication, in the YYYY-MM-DDThh:mm:ss.sss±hhmm format
	FirstLoginDate string `json:"firstLoginDate"`

	// Date and time of the user's last authentication, in the YYYY-MM-DDThh:mm:ss.sss±hhmm format
	LastLoginDate string `json:"lastLoginDate"`

	// Method of adding a user:
	// true: By sending an invitation by email
	// false: By other means
	WelcomeMailSent bool `json:"welcomeMailSent"`
}

func (t *trackerClient) Myself() (*User, error) {
	request := t.client.R().SetHeaders(t.headers)
	resp, err := request.Get(baseUrl + "/v2/myself")
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d, message=%s, headers=%s", resp.StatusCode(), string(resp.Body()), t.headers)
	}

	result := new(User)
	if err := json.Unmarshal(resp.Body(), result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}
