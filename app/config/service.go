package config

import (
	"github.com/eddoog/store-serve/service/auth"
	"github.com/eddoog/store-serve/service/user"
)

type Service struct {
	AuthService auth.IAuthService
	UserService user.IUserService
}

func InitService(
	repository *Repository,
) *Service {
	return &Service{
		AuthService: auth.InitAuthService(repository.AuthRepository),
		UserService: user.InitUserService(repository.UserRepository),
	}
}
