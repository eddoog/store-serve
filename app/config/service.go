package config

import (
	"github.com/eddoog/store-serve/service/auth"
)

type Service struct {
	AuthService auth.IAuthService
}

func InitService(
	repository *Repository,
) *Service {
	return &Service{
		AuthService: auth.InitAuthService(repository.AuthRepository),
	}
}
