package repositories

import (
	"myapp/models"
	"myapp/db"
)

type UserRepository struct {}

func (ur *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}