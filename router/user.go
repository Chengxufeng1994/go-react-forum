package router

import (
	"github.com/Chengxufeng1994/go-react-forum/controller"
	"github.com/Chengxufeng1994/go-react-forum/middlewares"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur UserRouter) InitUserRouter(routerGroup *gin.RouterGroup) {
	userController := new(controller.UserController)
	userGroup := routerGroup.Group("/user")
	{
		userGroup.POST("/", userController.CreateUser)
		userGroup.GET("/:id", userController.GetUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", middlewares.AuthMiddleware(), userController.DeleteUser)
	}
}
