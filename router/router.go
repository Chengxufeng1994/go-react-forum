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

	authRouter := new(AuthRouter)
	userRouter := new(UserRouter)
	postRouter := new(PostRouter)

	v1 := router.Group("/api/v1")
	{
		authRouter.InitAuthRouter(v1)
		userRouter.InitUserRouter(v1)
		postRouter.InitPostRouter(v1)
	}

	return router
}
