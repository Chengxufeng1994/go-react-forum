package model

import (
	"errors"
	"github.com/Chengxufeng1994/go-react-forum/util"
	"gorm.io/gorm"
	"html"
	"net/mail"
	"strings"
	"time"
)

type User struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `gorm:"datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"datetime;default:CURRENT_TIMESTAMP" json:"updated_at"`
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

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	result := db.Create(&u)
	if result.Error != nil {
		return &User{}, result.Error
	}

	return u, nil
}

func (u *User) FindUserById(db *gorm.DB, uid uint32) (*User, error) {
	result := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u)
	if result.Error != nil {
		return &User{}, result.Error
	}

	return u, nil
}

func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	updates := make(map[string]interface{})
	if u.Password != "" {
		err := u.BeforeSave()
		if err != nil {
			panic(err)
		}
		updates["password"] = u.Password
	}
	updates["email"] = u.Email
	updates["updated_at"] = time.Now()

	result := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Updates(updates)
	if result.Error != nil {
		return &User{}, result.Error
	}

	result2 := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u)
	if result2.Error != nil {
		return &User{}, result.Error
	}

	return u, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	result := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
	if result.Error != nil {
		return -1, result.Error
	}

	return db.RowsAffected, nil
}

func (u *User) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error
	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			err = errors.New("required email")
			errorMessages["Required_email"] = err.Error()
		}
		if u.Email != "" {
			if _, err = mail.ParseAddress(u.Email); err != nil {
				err = errors.New("invalid email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	default:
		if u.Username == "" {
			err = errors.New("required username")
			errorMessages["Required_username"] = err.Error()
		}
		if u.Password == "" {
			err = errors.New("required password")
			errorMessages["Required_password"] = err.Error()
		}
		if u.Password != "" && len(u.Password) < 6 {
			err = errors.New("password should be atleast 6 characters")
			errorMessages["Invalid_password"] = err.Error()
		}
		if u.Email == "" {
			err = errors.New("required email")
			errorMessages["Required_email"] = err.Error()
		}
		if u.Email != "" {
			if _, err = mail.ParseAddress(u.Email); err != nil {
				err = errors.New("invalid email")
				errorMessages["Invalid_email"] = err.Error()
			}
		}

	}

	return errorMessages
}
