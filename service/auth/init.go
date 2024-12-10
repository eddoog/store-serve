package auth

import (
	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/repository/auth"
)

type IAuthService interface {
	Login()
	Register(params entities.UserRegister) error
}

type AuthService struct {
	AuthRepository auth.IAuthRepository
}

func InitAuthService(
	authRepository auth.IAuthRepository,
) IAuthService {
	return &AuthService{
		AuthRepository: authRepository,
	}
}
