package postgres

import (
	model "github.com/andrey-tsyn/my-giga-wishes/auth-service/pkg/domain/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{db: db}
}

func (u *UserRepository) CreateUser(user model.User) error {
	// TODO: Not implemented
	panic("not implemented")
}

func (u *UserRepository) GetUserByUsername(username string) (model.User, error) {
	// TODO: Not implemented
	panic("not implemented")
}

func (u *UserRepository) GetUserByEmail(email string) (model.User, error) {
	// TODO: Not implemented
	panic("not implemented")
}

func (u *UserRepository) UpdateUser(user model.User) error {
	// TODO: Not implemented
	panic("not implemented")
}

func (u *UserRepository) DeleteUser(user model.User) error {
	// TODO: Not implemented
	panic("not implemented")
}
