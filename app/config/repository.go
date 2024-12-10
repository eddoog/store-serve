package config

import (
	"github.com/eddoog/store-serve/repository/auth"
	"github.com/eddoog/store-serve/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	AuthRepository auth.IAuthRepository
	UserRepository user.IUserRepository
}

func InitRepository(db *gorm.DB) *Repository {
	return &Repository{
		AuthRepository: auth.InitAuthRepository(db),
		UserRepository: user.InitUserRepository(db),
	}
}
