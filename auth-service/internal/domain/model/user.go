package model

type User struct {
	Id       int64
	Username username
	Email    email
	PassHash []byte
	Roles    []string
	IsActive bool
}
