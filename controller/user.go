package controller

import (
	"encoding/json"
	"github.com/Chengxufeng1994/go-react-forum/global"
	"github.com/Chengxufeng1994/go-react-forum/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UserController struct {
}

func (uc UserController) CreateUser(c *gin.Context) {
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

	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"statusCode": http.StatusUnprocessableEntity,
			"error":      errList,
		})
		return
	}
	user.Prepare()
	userCreated, err := user.SaveUser(global.GRF_DB)
	if err != nil {
		errList["Unmarshal_error"] = err.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"statusCode": http.StatusUnprocessableEntity,
			"error":      errList,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"statusCode": http.StatusCreated,
		"response":   userCreated,
	})
}

func (uc UserController) GetUser(c *gin.Context) {
	errList := map[string]string{}
	userId := c.Param("id")
	iUserId, err := strconv.ParseInt(userId, 10, 32)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	user := model.User{}
	userGotten, err := user.FindUserById(global.GRF_DB, uint32(iUserId))
	if err != nil {
		errList["no_user"] = "User not found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": userGotten,
	})
}

func (uc UserController) UpdateUser(c *gin.Context) {
	errList := map[string]string{}
	userId := c.Param("id")
	iUserId, err := strconv.ParseInt(userId, 10, 32)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	// Start processing the request
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	requestBody := map[string]string{}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	formerUser := model.User{}
	db := global.GRF_DB
	result := db.Debug().Model(model.User{}).Where("id = ?", iUserId).Take(&formerUser)
	if result.Error != nil {
		errList["User_invalid"] = "The user is does not exist"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

}

func (uc UserController) DeleteUser(c *gin.Context) {
	errList := map[string]string{}
	userId := c.Param("id")
	iUserId, err := strconv.ParseInt(userId, 10, 32)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	user := model.User{}
	_, err = user.DeleteUser(global.GRF_DB, uint32(iUserId))
	if err != nil {
		errList["Other_error"] = "Please try again later"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": "User deleted",
	})
}
