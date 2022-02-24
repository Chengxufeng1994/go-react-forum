package dao

import (
	"github.com/Chengxufeng1994/go-react-forum/database"
	"github.com/Chengxufeng1994/go-react-forum/model"
)

func Register(user *model.User) error {
	db := database.GetDB()
	result := db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func FindUserByEmail(email string) *model.User {
	var user *model.User
	db := database.GetDB()
	db.Where("email = ?", email).First(&user)

	return user
}
