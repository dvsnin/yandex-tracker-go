package tracker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

func (api *Client) createRequest(httpMethod string, fields map[string]string) *http.Request {
	data, _ := json.Marshal(fields)
	body := bytes.NewBufferString(string(data))
	req, _ := http.NewRequest(httpMethod, api.url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", api.token)
	req.Header.Add("X-Org-Id", api.xOrgID)

	return req
}

func (api *Client) parseURL(req *http.Request, apiMethod string) (err error) {
	req.URL, err = url.Parse(api.url)
	if err != nil {
		return
	}
	req.URL.Path = path.Join(req.URL.Path, apiMethod)

	return
}

// StatusCodeError represents an http response error.
// type httpStatusCode interface { HTTPStatusCode() int } to handle it.
type statusCodeError struct {
	Code   int
	Status string
}

func (t statusCodeError) Error() string {
	return fmt.Sprintf("tracker server error: %s", t.Status)
}

func checkStatusCode(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		return statusCodeError{Code: resp.StatusCode, Status: resp.Status}
	}

	return nil
}

func (api *Client) do(req *http.Request, ticketKey string) (*http.Response, error) {
	err := api.parseURL(req, ticketKey)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = checkStatusCode(resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
