package auth

import (
	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/domains/models"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	CheckUserExist(string) (bool, error)
	Register(params entities.UserRegister) error
	GetUserByEmail(string) (models.User, error)
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
