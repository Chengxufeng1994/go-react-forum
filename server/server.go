package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server interface {
	ListenAndServe() error
}

func Init(address string, router *gin.Engine) Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
