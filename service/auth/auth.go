package auth

import (
	"errors"

	"github.com/eddoog/store-serve/domains/entities"
)

func (a *AuthService) Login() {

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
