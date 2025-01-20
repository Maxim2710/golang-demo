package main

import (
	"github.com/gin-gonic/gin"
	"golang-demo/internal/app"
	"log"
)

func main() {
	gin.SetMode(gin.DebugMode)

	//gin.SetTrustedProxies([]string{"127.0.0.1"})

	if err := app.Run(); err != nil {
		log.Fatalf("Application terminated with error: %v", err)
	}
}
