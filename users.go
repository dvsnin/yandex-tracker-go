package tracker

type Users []User

// User
// Basic user structure in Yandex.Tracker
type User struct {
	Self    string
	ID      string
	Display string
}

// Id
// Get user id
func (u *User) Id() string {
	if u != nil {
		return u.ID
	}

	return ""
}

// Name
// Get username
func (u *User) Name() string {
	if u != nil {
		return u.Display
	}

	return ""
}
