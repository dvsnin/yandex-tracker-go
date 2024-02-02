package tracker

// Status
// https://cloud.yandex.ru/en/docs/tracker/concepts/issues/get-issue#status
type BasicStatus struct {
	// Address of the API resource with information about the status.
	Self string `json:"self"`

	// Status ID.
	ID string `json:"id"`

	// Status key.
	Key string `json:"key"`

	// Status name displayed.
	Display string `json:"display"`
}
