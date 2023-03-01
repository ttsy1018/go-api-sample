package repositories

import (
	"myapp/db"
	"myapp/models"
)

type UserRepository struct{}

// メールアドレスからユーザを取得
func (ur *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// メールアドレスからユーザリストを取得
func (ur *UserRepository) GetAllByEmail(email string) ([]models.User, error) {
	var users []models.User
	err := db.DB.Where("email = ?", email).Find(&users).Error // ※Findではレコードなしの場合でもエラーは返さない
	if err != nil {
		return users, err
	}

	return users, nil
}

// ユーザ作成
func (ul UserRepository) CreateUser(user *models.User) error {
	if err := db.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
