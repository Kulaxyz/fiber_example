package repository

import (
	"github.com/Kulaxyz/partner/database"
	model "github.com/Kulaxyz/partner/models"
)

func GetUserByEmail(e string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Email: e}).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(u string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Username: u}).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
