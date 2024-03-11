package repository

import "github.com/andrey-tsyn/my-giga-wishes/auth-service/internal/domain/model"

type UserRepository interface {
	CreateUser(user model.User) error
	GetUserByUsername(username string) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(user model.User) error
}
