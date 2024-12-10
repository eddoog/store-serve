package auth

import (
	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/domains/models"
	"github.com/eddoog/store-serve/repository/auth"
)

type IAuthService interface {
	Login(params entities.UserLogin) (models.User, error)
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
