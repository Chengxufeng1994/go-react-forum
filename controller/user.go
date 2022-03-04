package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Chengxufeng1994/go-react-forum/global"
	"github.com/Chengxufeng1994/go-react-forum/model"
	"github.com/Chengxufeng1994/go-react-forum/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	err = user.BeforeSave()
	if err != nil {
		errList["Before_save"] = "Hash password error"
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

	fmt.Println("requestBody: ", requestBody)
	updatedUser := model.User{}
	if requestBody["current_password"] == "" && requestBody["new_password"] != "" {
		errList["Empty_current"] = "Please Provide current password"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	if requestBody["current_password"] != "" && requestBody["new_password"] == "" {
		errList["Empty_new"] = "Please Provide new password"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	if requestBody["current_password"] != "" && requestBody["new_password"] != "" {
		// also check if the new password
		if len(requestBody["new_password"]) < 6 {
			errList["Invalid_password"] = "Password should be atleast 6 characters"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return
		}
		// if they do, check that the former password is correct
		err = util.VerifyPwd(formerUser.Password, requestBody["current_password"])
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			errList["Password_mismatch"] = "The password not correct"
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  errList,
			})
			return
		}
		// update both the password and the email
		updatedUser.Username = formerUser.Username // remember, you cannot update the username
		updatedUser.Email = requestBody["email"]
		updatedUser.Password = requestBody["new_password"]
	}

	updatedUser.Username = formerUser.Username
	updatedUser.Email = requestBody["email"]
	updatedUser.Prepare()
	_, err = updatedUser.UpdateUser(global.GRF_DB, uint32(iUserId))
	if err != nil {
		formattedError := util.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": updatedUser,
	})
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
