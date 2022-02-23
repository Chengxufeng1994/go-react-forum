package router

import (
	"github.com/Chengxufeng1994/go-react-forum/controller"
	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct {
}

func (ar AuthorityRouter) InitAuthorityRouter(routerGroup *gin.RouterGroup) {
	authorityGroup := routerGroup.Group("/auth")
	authorityController := new(controller.AuthorityController)
	{
		authorityGroup.GET("/register", authorityController.Register)
		authorityGroup.GET("/login", authorityController.Login)
		authorityGroup.GET("/logout", authorityController.Logout)
	}

}
