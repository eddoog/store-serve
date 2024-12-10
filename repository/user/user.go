package user

import "github.com/eddoog/store-serve/domains/models"

func (u *UserRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User

	if err := u.db.First(&user, userID).Error; err != nil {
		return &user, err
	}

	return &user, nil
}
