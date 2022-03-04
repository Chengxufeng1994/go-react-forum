package controller

import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/Chengxufeng1994/go-react-forum/auth"
	"github.com/Chengxufeng1994/go-react-forum/global"
	"github.com/Chengxufeng1994/go-react-forum/model"
	"github.com/Chengxufeng1994/go-react-forum/util"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthController struct{}

func (ac AuthController) Register(c *gin.Context) {
	var err error
	var errList = make(map[string]interface{})
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	user.Prepare()
	errorMessages := user.Validate("")
	if len(errorMessages) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
	}

	err = user.BeforeSave()
	if err != nil {
		errList["Before_save"] = "Hash password error"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"statusCode": http.StatusUnprocessableEntity,
			"error":      errList,
		})
		return
	}
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

func (ac AuthController) Login(c *gin.Context) {
	var loginRequest LoginRequest
	var err error
	// var errList = make(map[string]interface{})
	err = c.BindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  "unable get request",
		})
	}

	user := model.User{}
	user.Email = loginRequest.Email
	user.Password = loginRequest.Password
	errorMessages := user.Validate("login")
	if len(errorMessages) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errorMessages,
		})
		return
	}

	loginUser := model.User{}
	db := global.GRF_DB
	result := db.Debug().Model(&model.User{}).Where("email = ?", loginRequest.Email).Take(&loginUser)
	if result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  fmt.Sprintf("error at the getting the user: %#v", err.Error()),
		})
	}

	err = util.VerifyPwd(loginUser.Password, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  fmt.Sprintf("error at the hashing the password: %#v", err.Error()),
		})
	}

	token, err := auth.CreateToken(uint32(loginUser.ID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  fmt.Sprintf("error at the creating the token: %#v", err.Error()),
		})
		return
	}

	var userData = make(map[string]interface{})
	userData["id"] = loginUser.ID
	userData["username"] = loginUser.Username
	userData["email"] = loginUser.Email

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"user":   userData,
		"token":  token,
	})
}
