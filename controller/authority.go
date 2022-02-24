package controller

import (
	"github.com/Chengxufeng1994/go-react-forum/dao"
	"github.com/Chengxufeng1994/go-react-forum/model"
	"github.com/Chengxufeng1994/go-react-forum/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type LoginRequest struct {
	Email    string
	Password string
}

type AuthorityController struct{}

func (ac AuthorityController) Register(c *gin.Context) {
	json := map[string]string{}
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Register Failed",
			"result":  err.Error(),
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(json["password"]), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user := &model.User{}
	user.Username = json["username"]
	user.Email = json["email"]
	user.Password = string(hash)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = dao.Register(user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Register Failed",
			"result":  err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Register Successfully",
		})
		return
	}
}

func (ac AuthorityController) Login(c *gin.Context) {
	var loginRequest LoginRequest
	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login Failed",
			"result":  err.Error(),
		})
		return
	}

	if loginRequest.Email == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login Failed",
			"result":  "Email required",
		})
		return
	}
	if loginRequest.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login Failed",
			"result":  "Password required",
		})
		return
	}

	user := dao.FindUserByEmail(loginRequest.Email)
	if user == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login Failed",
			"result":  "User not found",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login Failed",
			"result":  "Password wrong",
		})
		return
	}

	sess := &model.Session{}
	uuid := util.CreateUUID()
	sess.SessionID = uuid
	sess.Username = user.Username
	sess.UserID = user.ID
	err = dao.CreateSession(sess)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login Failed",
			"result":  "Create session failed",
		})
	}
	c.SetCookie("user", uuid, 60*60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully",
	})
}

func (ac AuthorityController) Logout(c *gin.Context) {
	uuid, _ := c.Cookie("user")
	if uuid != "" {
		dao.DeleteSession(uuid)
		c.SetCookie("user", "", -1, "/", "localhost", false, true)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Successfully",
	})
}
