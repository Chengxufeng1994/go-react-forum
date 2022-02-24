package dao

import (
	"github.com/Chengxufeng1994/go-react-forum/database"
	"github.com/Chengxufeng1994/go-react-forum/model"
)

func CreateSession(session *model.Session) error {
	db := database.GetDB()
	result := db.Create(&session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteSession(uuid string) error {
	db := database.GetDB()
	db.Where("session_id = ?", uuid).Delete(&model.Session{})

	return nil
}
