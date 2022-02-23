package router

import (
	"github.com/Chengxufeng1994/go-react-forum/controller"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controller.HealthController)
	// Health Check
	router.GET("/health", health.Status)

	authorityRouter := new(AuthorityRouter)

	apiGroup := router.Group("api")
	publicGroup := apiGroup.Group("")
	{
		authorityRouter.InitAuthorityRouter(publicGroup)
	}
	// privateGroup := apiGroup.Group("")

	return router
}
