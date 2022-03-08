package controller

import (
	"encoding/json"
	"github.com/Chengxufeng1994/go-react-forum/auth"
	"github.com/Chengxufeng1994/go-react-forum/global"
	"github.com/Chengxufeng1994/go-react-forum/model"
	"github.com/Chengxufeng1994/go-react-forum/util"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type PostController struct {
}

type PostForm struct {
}

func (pc PostController) CreatePost(c *gin.Context) {
	errList := map[string]string{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_Body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"statusCode": http.StatusUnprocessableEntity,
			"error":      errList,
		})
		return
	}

	post := model.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"statusCode": http.StatusUnprocessableEntity,
			"error":      errList,
		})
		return
	}

	uid, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}

	user := model.User{}
	_, err = user.FindUserById(global.GRF_DB, uid)
	if err != nil {
		errList["Unauthorized"] = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  errList,
		})
		return
	}
	post.AuthorID = uid
	post.Prepare()
	postCreated, err := post.SavePost(global.GRF_DB)
	if err != nil {
		errList := util.FormatError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusCreated,
		"response": postCreated,
	})
}

func (pc PostController) GetPosts(c *gin.Context) {}

func (pc PostController) GetPost(c *gin.Context) {}
