package user

import "github.com/eddoog/store-serve/domains/models"

func (u *UserService) GetProfile(userID uint) (*models.User, error) {
	user, err := u.UserRepository.GetUserByID(userID)

	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}
