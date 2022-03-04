package main

import (
	"github.com/Chengxufeng1994/go-react-forum/config"
	"github.com/Chengxufeng1994/go-react-forum/database"
	"github.com/Chengxufeng1994/go-react-forum/global"
	"github.com/Chengxufeng1994/go-react-forum/router"
	"github.com/Chengxufeng1994/go-react-forum/server"
	"log"
)

func main() {
	global.GRF_VP = config.Init("development")

	username := config.GetConfig().GetString("mysql.username")
	password := config.GetConfig().GetString("mysql.password")
	host := config.GetConfig().GetString("mysql.host")
	port := config.GetConfig().GetString("mysql.port")
	dbname := config.GetConfig().GetString("mysql.dbname")
	global.GRF_DB = database.Init(username, password, host, port, dbname)

	if global.GRF_DB != nil {
		database.RegisterTables(global.GRF_DB)
	}

	r := router.Init()
	serverPort := config.GetConfig().GetString("server.port")
	s := server.Init(serverPort, r)

	log.Fatalf(s.ListenAndServe().Error())
}
