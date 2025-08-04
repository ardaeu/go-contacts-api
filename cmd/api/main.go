package main

import (
	"github.com/ardaeu/go-contacts-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handler.PingHandler)

	r.Run()
}
