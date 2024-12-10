package config

import (
	"github.com/eddoog/store-serve/repository/auth"
	"gorm.io/gorm"
)

type Repository struct {
	AuthRepository auth.IAuthRepository
}

func InitRepository(db *gorm.DB) *Repository {
	return &Repository{
		AuthRepository: auth.InitAuthRepository(db),
	}
}
