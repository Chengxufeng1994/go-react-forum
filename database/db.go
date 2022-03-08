package database

import (
	"fmt"
	"github.com/Chengxufeng1994/go-react-forum/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(user, pwd, host, port, dbname string) *gorm.DB {
	var err error
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, port, dbname)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Fatal connect database: %s \n", err))
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}

func RegisterTables(db *gorm.DB) {
	err := db.Debug().AutoMigrate(
		&model.User{},
		&model.Post{},
	)
	if err != nil {
		panic(fmt.Errorf("register table failed: %#v", err))
	}

	fmt.Println("register table success")
}
