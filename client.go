package tracker

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

var (
	_ Client = (*trackerClient)(nil)
)

const (
	ticketUrl      = "https://api.tracker.yandex.net/v2/issues/"
	ticketComments = "/comments"
)

type Client interface {
	// GetTicket - get Yandex.Tracker ticket by ticket key
	GetTicket(ticketKey string) (ticket Ticket, err error)
	// PatchTicket - patch Yandex.Tracker ticket by ticket key
	PatchTicket(ticketKey string, body map[string]string) (ticket Ticket, err error)
	// GetTicketComments - get Yandex.Tracker ticket comments by ticket key
	GetTicketComments(ticketKey string) (comments TicketComments, err error)
}

func New(token, xOrgID string) Client {
	return &trackerClient{
		client: resty.New(),
		headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": token,
			"X-Org-Id":      xOrgID,
		},
	}
}

type trackerClient struct {
	headers map[string]string
	client  *resty.Client
}

func (t *trackerClient) GetTicket(ticketKey string) (Ticket, error) {
	request := t.client.R().SetHeaders(t.headers)
	resp, err := request.Get(ticketUrl + ticketKey)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d, message=%s", resp.StatusCode(), string(resp.Body()))
	}

	var result Ticket
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}

func (t *trackerClient) PatchTicket(ticketKey string, body map[string]string) (Ticket, error) {
	request := t.client.R().SetHeaders(t.headers)
	resp, err := request.
		SetBody(body).
		Patch(ticketUrl + ticketKey)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d, message=%s", resp.StatusCode(), string(resp.Body()))
	}

	var result Ticket
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}

func (t *trackerClient) GetTicketComments(ticketKey string) (TicketComments, error) {
	request := t.client.R().SetHeaders(t.headers)
	resp, err := request.Get(ticketUrl + ticketKey + ticketComments)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d, message=%s", resp.StatusCode(), string(resp.Body()))
	}

	var result TicketComments
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}
