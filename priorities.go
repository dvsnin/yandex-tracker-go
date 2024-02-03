package tracker

// BasicPriority
// https://cloud.yandex.ru/en/docs/tracker/concepts/issues/get-issue#priority
type BasicPriority struct {
	// Address of the API resource with information about the priority.
	Self string `json:"self"`

	// Priority ID.
	ID string `json:"id"`

	// Priority key.
	Key string `json:"key"`

	// Priority name displayed.
	Display string `json:"display"`
}

// Priority
// https://cloud.yandex.ru/en/docs/tracker/concepts/issues/get-priorities
type Priority struct {
	// Address of the API resource with information about the priority.
	Self string `json:"self"`

	// Priority ID.
	ID int `json:"id"`

	// Priority key.
	Key string `json:"key"`

	// Priority version.
	Version int `json:"version"`

	// Priority name displayed.
	// If localized=false is provided in the request, this parameter duplicates the name in other languages.
	Name string `json:"name"`

	// Priority weight. This parameter affects the order of priority display in the interface.
	Order int `json:"order"`
}
