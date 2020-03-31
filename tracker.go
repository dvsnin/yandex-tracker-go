package tracker

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

const (
	TRACKER_URL = "https://api.tracker.yandex.net/v2/issues/"
)

type Tracker struct {
	token   string
	xOrgID  string
	url     string
	request *resty.Request
}

func New(token string, xOrgID string) *Tracker {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": token,
		"X-Org-Id":      xOrgID,
	}
	return &Tracker{
		token:   token,
		xOrgID:  xOrgID,
		url:     TRACKER_URL,
		request: resty.New().R().SetHeaders(headers),
	}
}

// Get Yandex.Tracker ticket by ticket key
func (t *Tracker) GetTicket(ticketKey string) (ticket Ticket, err error) {
	resp, err := t.request.Get(TRACKER_URL + ticketKey)
	if resp != nil {
		defer func() {
			if err := resp.RawBody().Close(); err != nil {
				return
			}
		}()
	}
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Body(), &ticket)
	if err != nil {
		return
	}
	return
}

// Patch Yandex.Tracker ticket by ticket key
func (t *Tracker) PatchTicket(ticketKey string, body map[string]string) (ticket Ticket, err error) {
	resp, err := t.request.
		SetBody(body).
		Patch(TRACKER_URL + ticketKey)
	if resp != nil {
		defer func() {
			if err := resp.RawBody().Close(); err != nil {
				return
			}
		}()
	}
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Body(), &ticket)
	if err != nil {
		return
	}
	return
}
