package tracker

type BasicSprint struct {
	// Address of the API resource with information about the sprint.
	Self string `json:"self"`

	// Sprint ID.
	ID string `json:"id"`

	// Sprint name displayed.
	Display string `json:"display"`
}
