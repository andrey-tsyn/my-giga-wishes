package model

type username struct {
	data string
}

func NewUsername(usernameStr string) (username, error) {
	// TODO: name validation
	return username{data: usernameStr}, nil
}

func (u *username) String() string {
	return u.data
}

type email struct {
	data string
}

func (e *email) String() string {
	return e.data
}

func NewEmail(emailStr string) (email, error) {
	// TODO: email validation
	return email{data: emailStr}, nil
}
