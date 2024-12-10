package auth

import (
	"errors"

	"github.com/eddoog/store-serve/domains/entities"
	"github.com/eddoog/store-serve/domains/models"
)

func (a *AuthService) Login(params entities.UserLogin) (models.User, error) {
	user, err := a.AuthRepository.GetUserByEmail(params.Email)

	if err != nil {
		return models.User{}, errors.New("email or password is wrong")
	}

	if !CheckPasswordHash(params.Password, user.Password) {
		return models.User{}, errors.New("email or password is wrong")
	}

	return user, nil

}

func (a *AuthService) Register(params entities.UserRegister) error {
	exists, err := a.AuthRepository.CheckUserExist(params.Email)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("email already in use")
	}

	params.Password, _ = HashPassword(params.Password)

	return a.AuthRepository.Register(params)
}
