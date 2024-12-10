package user

import (
	"github.com/eddoog/store-serve/domains/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByID(userID uint) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
