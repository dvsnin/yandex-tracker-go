package tracker

const (
	ApiUrl = "https://api.tracker.yandex.net/"
)

func New(token string, xOrgID string) *Client {
	return &Client{
		token:  token,
		xOrgID: xOrgID,
		url:    ApiUrl,
	}
}

type Client struct {
	token  string
	xOrgID string
	url    string
}
