package tracker

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Basic Issue structure in Yandex.Tracker
type BasicIssue struct {
	// Address of the API resource with information about the issue.
	Self string `json:"self"`

	// Issue ID.
	ID string `json:"id"`

	// Issue key.
	Key string `json:"key"`

	// Issue name displayed.
	Display string `json:"display"`
}

// Issue structure in Yandex.Tracker
// https://cloud.yandex.ru/en/docs/tracker/concepts/issues/get-issue
type Issue struct {
	// Address of the API resource with information about the issue.
	Self string `json:"self"`

	// Issue ID.
	ID string `json:"id"`

	// Issue key.
	Key string `json:"key"`

	// Issue version. Each change to the issue parameters increases its version number.
	Version int `json:"version"`

	// Date and time when the last comment was added.
	LastCommentUpdatedAt string `json:"lastCommentUpdatedAt"`

	// Issue name.
	Summary string `json:"summary"`

	// Object with information about the parent issue.
	Parent *BasicIssue `json:"parent"`

	// Array with information about alternative issue keys.
	Aliases []string `json:"aliases"`

	// 	Object with information about the employee who edited the issue last.
	UpdatedBy *BasicUser `json:"updatedBy"`

	// Issue description.
	Description string `json:"description"`

	// Array of objects with information about the sprint.
	Sprint []*BasicSprint `json:"sprint"`

	// Object with information about the issue type.
	Type *IssueType `json:"type"`

	// Object with information about the priority.
	Priority *BasicPriority `json:"priority"`

	// Issue creation date and time.
	CreatedAt string `json:"createdAt"`

	// Array of objects with information about issue followers.
	Followers []*BasicUser `json:"followers"`

	// Object with information about the user who created the issue.
	CreatedBy *BasicUser `json:"createdBy"`

	// Number of votes for the issue.
	Votes int `json:"votes"`

	// Object with information about the issue's assignee.
	Assignee *BasicUser `json:"assignee"`

	// Object with information about the issue queue.
	Queue *BasicQueue `json:"queue"`

	// Date and time when the issue was last updated.
	UpdatedAt string `json:"updatedAt"`

	// Object with information about the issue status.
	Status *BasicStatus `json:"status"`

	// Object with information about the previous status of the issue.
	PreviousStatus *BasicStatus `json:"previousStatus"`

	// Favorite issue flag:
	// true: Issue added to favorites by the user.
	// false: Issue not added to favorites.
	Favorite bool `json:"favorite"`
}

// https://cloud.yandex.ru/en/docs/tracker/concepts/issues/create-issue
type CreateIssueOptions struct {
	// Issue name. Required.
	Summary *string `json:"summary,omitempty"`

	// Queue in which to create the issue. Required.
	// Can be set as an object, a string (if the queue key is provided), or a number (if the queue ID is provided).
	Queue interface{} `json:"queue,omitempty"`

	// Parent issue.
	// Object or string.
	Parent interface{} `json:"parent,omitempty"`

	// Issue description.
	Description *string `json:"description,omitempty"`

	// Block with information about sprints.
	// Array of objects or strings.
	Sprint *[]interface{} `json:"sprint,omitempty"`

	// Issue type.
	// Can be set as an object, a string (if the issue type key is provided), or a number (if the issue type ID is provided).
	Type interface{} `json:"type,omitempty"`

	// Issue priority.
	// Can be set as an object, a string (if the priority key is provided), or a number (if the priority ID is provided).
	Priority interface{} `json:"priority,omitempty"`

	// IDs or usernames of issue followers.
	// Array of objects, numbers, or strings.
	Followers *[]interface{} `json:"followers,omitempty"`

	// ID or username of issue assignee.
	// Object, number, or string.
	Assignee interface{} `json:"assignee,omitempty"`

	// ID or username of issue author.
	// Object, number, or string.
	Author interface{} `json:"author,omitempty"`

	// Field with a unique value that disables creation of duplicate issues.
	// If you try to create an issue with the same value of this parameter again, no duplicate will be created and the response will contain an error with code 409.
	Unique *string `json:"unique,omitempty"`

	// List of attachment IDs.
	// Array of strings
	AttachmentIDs *[]string `json:"attachmentIds,omitempty"`
}

type ListOptions struct {
	// Additional fields to be included into the response:
	// transitions: Workflow transitions between statuses
	// attachments: Attached files
	Expand string

	// Number of issues per response page. The default value is 50. To set up additional response output parameters, use pagination.
	// https://cloud.yandex.ru/en/docs/tracker/common-format#displaying-results
	PerPage int
}

type FindIssuesOptions struct {
	// Queue
	Queue *string `json:"queue,omitempty"`

	// List of issue keys
	// String or Array of strings
	Keys interface{} `json:"keys,omitempty"`

	// Issue filtering parameters. The parameter can specify any field and value to filter by.
	Filter map[string]interface{} `json:"filter,omitempty"`

	// Filter using the query language
	// https://cloud.yandex.ru/en/docs/tracker/user/query-filter
	Query *string `json:"query,omitempty"`
}

func (t *trackerClient) CreateIssue(opts *CreateIssueOptions) (*Issue, *resty.Response, error) {
	req := t.NewRequest(resty.MethodPost, "/v2/issues/", opts)
	result := new(Issue)
	resp, err := t.Do(req, result)
	if err != nil {
		return nil, nil, fmt.Errorf("request: %w", err)
	}
	return result, resp, nil
}

func (t *trackerClient) FindIssues(opts *FindIssuesOptions, listOpts *ListOptions) ([]*Issue, *resty.Response, error) {
	req := t.NewRequest(resty.MethodPost, "/v2/issues/_search", opts)
	// TODO:
	if listOpts.Expand != "" {
		req.SetQueryParam("expand", listOpts.Expand)
	}
	if listOpts.PerPage > 0 {
		req.SetQueryParam("perPage", fmt.Sprint(listOpts.PerPage))
	}
	var result []*Issue
	resp, err := t.Do(req, &result)
	if err != nil {
		return nil, nil, fmt.Errorf("request: %w", err)
	}

	return result, resp, nil
}

func (t *trackerClient) GetIssue(issueKey string) (*Issue, *resty.Response, error) {
	req := t.NewRequest(resty.MethodGet, "/v2/issues/"+issueKey, nil)
	result := new(Issue)
	resp, err := t.Do(req, result)
	if err != nil {
		return nil, nil, fmt.Errorf("request: %w", err)
	}

	return result, resp, err
}
