package user

import (
	"github.com/eddoog/store-serve/domains/models"
	"github.com/eddoog/store-serve/repository/user"
)

type IUserService interface {
	GetProfile(userID uint) (*models.User, error)
}

type UserService struct {
	UserRepository user.IUserRepository
}

func InitUserService(userRepository user.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepository,
	}
}
