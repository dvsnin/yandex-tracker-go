package tracker

type BasicQueue struct {
	// Address of the API resource with information about the queue.
	Self string `json:"self"`

	// Queue ID.
	ID string `json:"id"`

	// Queue key.
	Key string `json:"key"`

	// Queue name displayed.
	Display string `json:"display"`
}
