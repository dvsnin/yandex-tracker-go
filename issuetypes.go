package tracker

// IssueType
// https://cloud.yandex.ru/en/docs/tracker/concepts/issues/get-issue#type
type IssueType struct {
	// Address of the API resource with information about the issue type.
	Self string `json:"self"`

	// ID of the issue type.
	ID string `json:"id"`

	// Key of the issue type.
	Key string `json:"key"`

	// Issue type name displayed.
	Display string `json:"display"`
}
