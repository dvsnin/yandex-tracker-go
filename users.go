package tracker

type User struct {
	Self    string
	ID      string
	Display string
}

type Users []User

func (u *User) Id() string {
	if u != nil {
		return u.ID
	}

	return ""
}

func (u *User) Name() string {
	if u != nil {
		return u.Display
	}

	return ""
}
