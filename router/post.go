package router

import (
	"github.com/Chengxufeng1994/go-react-forum/controller"
	"github.com/Chengxufeng1994/go-react-forum/middlewares"
	"github.com/gin-gonic/gin"
)

type PostRouter struct {
}

func (pr *PostRouter) InitPostRouter(routerGroup *gin.RouterGroup) {
	postController := new(controller.PostController)
	postGroup := routerGroup.Group("/posts")
	{
		postGroup.POST("/", middlewares.AuthMiddleware(), postController.CreatePost)
		postGroup.GET("/all", postController.GetPosts)
		postGroup.GET("/:id", postController.GetPost)
	}
}
