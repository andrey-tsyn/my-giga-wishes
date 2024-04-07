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
	data       string
	isVerified bool
}

func (e *email) String() string {
	return e.data
}

func (e *email) GetEmail() string {
	return e.data
}

func (e *email) IsVerified() bool {
	return e.isVerified
}

func NewEmail(emailStr string, isVerified bool) (email, error) {
	// TODO: email validation
	return email{data: emailStr, isVerified: isVerified}, nil
}
