package tracker

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	ticketUrl      = "https://api.tracker.yandex.net/v2/issues/"
	ticketComments = "/comments"
)

type Tracker struct {
	headers map[string]string
	client  *resty.Client
}

func New(token string, xOrgID string) *Tracker {
	return &Tracker{
		client: resty.New(),
		headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": token,
			"X-Org-Id":      xOrgID,
		},
	}
}

// GetTicket
// Get Yandex.Tracker ticket by ticket key
func (t *Tracker) GetTicket(ticketKey string) (ticket Ticket, err error) {
	request := t.client.R().SetHeaders(t.headers)
	resp, err := request.Get(ticketUrl + ticketKey)
	if err != nil {
		return
	}
	defer resp.RawBody().Close()

	if resp.StatusCode() != http.StatusOK {
		return ticket, fmt.Errorf(string(resp.Body()))
	}

	if err := json.Unmarshal(resp.Body(), &ticket); err != nil {
		return
	}

	return
}

// PatchTicket
// Patch Yandex.Tracker ticket by ticket key
func (t *Tracker) PatchTicket(ticketKey string, body map[string]string) (ticket Ticket, err error) {
	request := t.client.R().SetHeaders(t.headers)
	resp, err := request.
		SetBody(body).
		Patch(ticketUrl + ticketKey)
	if err != nil {
		return
	}
	defer resp.RawBody().Close()

	if resp.StatusCode() != http.StatusOK {
		return ticket, fmt.Errorf(string(resp.Body()))
	}

	if err := json.Unmarshal(resp.Body(), &ticket); err != nil {
		return
	}
	return
}

// GetTicketComments
// Get Yandex.Tracker ticket comments by ticket key
func (t *Tracker) GetTicketComments(ticketKey string) (comments TicketComments, err error) {
	request := t.client.R().SetHeaders(t.headers)
	resp, err := request.Get(ticketUrl + ticketKey + ticketComments)
	if err != nil {
		return
	}
	defer resp.RawBody().Close()

	if resp.StatusCode() != http.StatusOK {
		return comments, fmt.Errorf(string(resp.Body()))
	}

	if err := json.Unmarshal(resp.Body(), &comments); err != nil {
		return
	}

	return
}
