package auth

import (
	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/domains/models"
)

func (a *AuthRepository) CheckUserExist(email string) (bool, error) {
	var count int64

	err := a.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (a *AuthRepository) Register(params entities.UserRegister) error {
	user := models.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: params.Password,
	}

	return a.db.Create(&user).Error
}

func (a *AuthRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := a.db.Where("email = ?", email).First(&user).Error

	return user, err
}
