// services/auth.go
package services

import (
	"auth-simple/models"
	"auth-simple/utils"
	"errors"

	"gorm.io/gorm"
)

func Register(db *gorm.DB, username, email string) (string, error) {
	pw, err := utils.GenerateRandomPassword()
	if err != nil {
		return "", err
	}
	hashed, err := utils.HashPassword(pw)
	if err != nil {
		return "", err
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: hashed,
	}

	if err := db.Create(&user).Error; err != nil {
		return "", err
	}
	// "Send" password back (simulate email)
	return pw, nil
}

func Login(db *gorm.DB, username, pw string) (*models.User, error) {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	if !utils.CheckPassword(user.Password, pw) {
		return nil, errors.New("invalid password")
	}
	return &user, nil
}

func ChangePassword(db *gorm.DB, username, oldPw, newPw string) error {
	user, err := Login(db, username, oldPw)
	if err != nil {
		return err
	}
	hashed, err := utils.HashPassword(newPw)
	if err != nil {
		return err
	}
	user.Password = hashed
	return db.Save(user).Error
}
