package auth

import (
	"github.com/eddoog/store-serve/domains/entities"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	CheckUserExist(string) (bool, error)
	Register(params entities.UserRegister) error
}

type AuthRepository struct {
	db *gorm.DB
}

func InitAuthRepository(
	db *gorm.DB,
) IAuthRepository {
	return &AuthRepository{
		db: db,
	}
}
