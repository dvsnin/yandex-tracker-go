package tracker

type Ticket map[string]interface{}

// CreatedBy
// Get ticket author
func (t Ticket) CreatedBy() BasicUser {
	if createdBy, ok := t["createdBy"].(map[string]interface{}); ok {
		return BasicUser{
			Self:    toString(createdBy["self"]),
			ID:      toString(createdBy["id"]),
			Display: toString(createdBy["display"]),
		}
	}

	return BasicUser{}
}

// Assignee
// Get ticket assignee
func (t Ticket) Assignee() BasicUser {
	if assignee, ok := t["assignee"].(map[string]interface{}); ok {
		return BasicUser{
			Self:    toString(assignee["self"]),
			ID:      toString(assignee["id"]),
			Display: toString(assignee["display"]),
		}
	}

	return BasicUser{}
}

// Followers
// Get ticket followers
func (t Ticket) Followers() BasicUsers {
	if followers, ok := t["followers"].([]interface{}); ok {
		users := make(BasicUsers, len(followers))
		for i := range followers {
			users[i] = BasicUser{
				Self:    toString(followers[i].(map[string]interface{})["self"]),
				ID:      toString(followers[i].(map[string]interface{})["id"]),
				Display: toString(followers[i].(map[string]interface{})["display"]),
			}
		}
		return users
	}

	return BasicUsers{}
}

// Summary
// Get ticket summary
func (t Ticket) Summary() string {
	return t.GetField("summary")
}

// Key
// Get ticket key
func (t Ticket) Key() string {
	return t.GetField("key")
}

// Description
// Get ticket description
func (t Ticket) Description() string {
	return t.GetField("description")
}

// Status
// Get ticket status
func (t Ticket) Status() string {
	if status, ok := t["status"].(map[string]interface{}); ok {
		if display, ok := status["display"]; ok {
			return toString(display)
		}
	}
	return ""
}

// SlackMessageID
// Get ticket custom field slack message id
func (t Ticket) SlackMessageID() string {
	return t.GetField("slackMessageID")
}

// GetField
// Get any custom ticket field
func (t Ticket) GetField(field string) string {
	if key, ok := t[field]; ok {
		return toString(key)
	}
	return ""
}

func toString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	default:
		return ""
	}
}
