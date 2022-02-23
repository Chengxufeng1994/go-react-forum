package main

import (
	"github.com/Chengxufeng1994/go-react-forum/config"
	"github.com/Chengxufeng1994/go-react-forum/database"
	"github.com/Chengxufeng1994/go-react-forum/router"
	"github.com/Chengxufeng1994/go-react-forum/server"
	"log"
)

func main() {
	config.Init("development")

	username := config.GetConfig().GetString("mysql.username")
	password := config.GetConfig().GetString("mysql.password")
	host := config.GetConfig().GetString("mysql.host")
	port := config.GetConfig().GetString("mysql.port")
	dbname := config.GetConfig().GetString("mysql.database-name")
	database.Init(username, password, host, port, dbname)

	router := router.Init()

	serverPort := config.GetConfig().GetString("server.port")
	server := server.Init(serverPort, router)

	log.Fatalf(server.ListenAndServe().Error())
}
