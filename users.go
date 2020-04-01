package tracker

// Basic user structure in Yandex.Tracker
type User struct {
	Self    string
	ID      string
	Display string
}

type Users []User

// Get user id
func (u *User) Id() string {
	if u != nil {
		return u.ID
	}

	return ""
}

// Get user name
func (u *User) Name() string {
	if u != nil {
		return u.Display
	}

	return ""
}
