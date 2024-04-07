package service

import (
	"github.com/andrey-tsyn/my-giga-wishes/auth-service/internal/domain/repository"
	"github.com/andrey-tsyn/my-giga-wishes/auth-service/pkg/domain/model"
)

type AuthService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userReposiotory repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userReposiotory,
	}
}

func (a *AuthService) CreateUser(user model.User, password string) (*model.User, error) {
	// TODO: Not implemented
	return nil, nil
}
