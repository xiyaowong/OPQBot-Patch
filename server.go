package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(AuthorizeMiddleware())

	router.NoRoute(BotProxyHandler)

	addr := fmt.Sprintf(":%d", ServerPort)
	fmt.Printf("Running on => 0.0.0.0%s\n", addr)
	go SendMsg()
	router.Run(addr)
}
