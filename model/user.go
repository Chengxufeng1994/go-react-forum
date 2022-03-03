package model

import (
	"github.com/Chengxufeng1994/go-react-forum/database"
	"github.com/Chengxufeng1994/go-react-forum/util"
	"html"
	"strings"
	"time"
)

type User struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := util.Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) SaveUser() (*User, error) {
	db := database.GetDB()
	result := db.Create(&u)
	if result.Error != nil {
		return &User{}, result.Error
	}

	return u, nil
}

func (u *User) FindUserById(uid uint32) (*User, error) {
	db := database.GetDB()
	result := db.Debug().Model(User{}).Where("id = ?", uid).Take(&u)
	if result.Error != nil {
		return &User{}, result.Error
	}

	return u, nil
}

func (u *User) DeleteUser(uid uint32) (int64, error) {
	db := database.GetDB()
	result := db.Debug().Model(User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
	if result.Error != nil {
		return -1, result.Error
	}

	return db.RowsAffected, nil
}
