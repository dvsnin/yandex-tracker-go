package tracker

type TicketComments []TicketComment
type TicketComment map[string]interface{}

// GetLast
// Return last comment
func (t TicketComments) GetLast() TicketComment {
	countComments := len(t)
	if countComments == 0 {
		return TicketComment{}
	}

	return t[len(t)-1]
}

// CreatedBy
// Get comment author
func (t TicketComment) CreatedBy() BasicUser {
	if createdBy, ok := t["createdBy"].(map[string]interface{}); ok {
		return BasicUser{
			Self:    toString(createdBy["self"]),
			ID:      toString(createdBy["id"]),
			Display: toString(createdBy["display"]),
		}
	}

	return BasicUser{}
}

// Text
// Get comment text
func (t TicketComment) Text() string {
	return t.GetField("text")
}

// Summonees
// Get comment author
func (t TicketComment) Summonees() BasicUsers {
	if summonees, ok := t["summonees"].([]interface{}); ok {
		users := make(BasicUsers, len(summonees))
		for i := range summonees {
			users[i] = BasicUser{
				Self:    toString(summonees[i].(map[string]interface{})["self"]),
				ID:      toString(summonees[i].(map[string]interface{})["id"]),
				Display: toString(summonees[i].(map[string]interface{})["display"]),
			}
		}
		return users
	}

	return BasicUsers{}
}

// GetField
// Get any custom ticket field
func (t TicketComment) GetField(field string) string {
	if key, ok := t[field]; ok {
		return toString(key)
	}

	return ""
}
