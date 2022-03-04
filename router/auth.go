package router

import (
	"github.com/Chengxufeng1994/go-react-forum/controller"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (ar AuthRouter) InitAuthRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/auth")
	authController := new(controller.AuthController)
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
	}

}
