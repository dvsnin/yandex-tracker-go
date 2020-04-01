package tracker

type Ticket map[string]interface{}

// Get ticket author
func (t Ticket) CreatedBy() User {
	if createdBy, ok := t["createdBy"].(map[string]interface{}); ok {
		return User{
			Self:    toString(createdBy["self"]),
			ID:      toString(createdBy["id"]),
			Display: toString(createdBy["display"]),
		}
	}

	return User{}
}

// Get ticket assignee
func (t Ticket) Assignee() User {
	if assignee, ok := t["assignee"].(map[string]interface{}); ok {
		return User{
			Self:    toString(assignee["self"]),
			ID:      toString(assignee["id"]),
			Display: toString(assignee["display"]),
		}
	}

	return User{}
}

// Get ticket followers
func (t Ticket) Followers() Users {
	if followers, ok := t["followers"].([]interface{}); ok {
		users := make(Users, len(followers))
		for i := range followers {
			users[i] = User{
				Self:    toString(followers[i].(map[string]interface{})["self"]),
				ID:      toString(followers[i].(map[string]interface{})["id"]),
				Display: toString(followers[i].(map[string]interface{})["display"]),
			}
		}
		return users
	}

	return Users{}
}

// Get ticket summary
func (t Ticket) Summary() string {
	return t.GetField("summary")
}

// Get ticket key
func (t Ticket) Key() string {
	return t.GetField("key")
}

// Get ticket description
func (t Ticket) Description() string {
	return t.GetField("description")
}

// Get ticket status
func (t Ticket) Status() string {
	if status, ok := t["status"].(map[string]interface{}); ok {
		if display, ok := status["display"]; ok {
			return toString(display)
		}
	}
	return ""
}

// Get ticket custom field slack message id
func (t Ticket) SlackMessageID() string {
	return t.GetField("slackMessageID")
}

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
