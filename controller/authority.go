package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthorityController struct{}

func (ac AuthorityController) Register(c *gin.Context) {
	c.String(http.StatusOK, "Register!")
}

func (ac AuthorityController) Login(c *gin.Context) {
	c.String(http.StatusOK, "Login!")
}

func (ac AuthorityController) Logout(c *gin.Context) {
	c.String(http.StatusOK, "Logout!")
}
