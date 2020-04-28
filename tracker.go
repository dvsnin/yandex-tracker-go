package tracker

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	ticketUrl      = "https://api.tracker.yandex.net/v2/issues/"
	ticketComments = "/comments"
)

type Tracker struct {
	request *resty.Request
}

func New(token string, xOrgID string) *Tracker {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": token,
		"X-Org-Id":      xOrgID,
	}
	return &Tracker{
		request: resty.New().R().SetHeaders(headers),
	}
}

// Get Yandex.Tracker ticket by ticket key
func (t *Tracker) GetTicket(ticketKey string) (ticket Ticket, err error) {
	resp, err := t.request.Get(ticketUrl + ticketKey)
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

	if resp.StatusCode() != 200 {
		return ticket, fmt.Errorf(string(resp.Body()))
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
		Patch(ticketUrl + ticketKey)
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

	if resp.StatusCode() != 200 {
		return ticket, fmt.Errorf(string(resp.Body()))
	}

	err = json.Unmarshal(resp.Body(), &ticket)
	if err != nil {
		return
	}
	return
}

// Get Yandex.Tracker ticket comments by ticket key
func (t *Tracker) GetTicketComments(ticketKey string) (comments TicketComments, err error) {
	resp, err := t.request.Get(ticketUrl + ticketKey + ticketComments)
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

	if resp.StatusCode() != 200 {
		return comments, fmt.Errorf(string(resp.Body()))
	}

	err = json.Unmarshal(resp.Body(), &comments)
	if err != nil {
		return
	}

	return
}
