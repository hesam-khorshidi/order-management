package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"order-management/common"
	"os"
	"strings"
)

func main() {
	common.LoadEnv()

	// Getting engine configurations
	ginPort := common.GetEnv("PORT", "8080")
	ginMode := common.GetEnv("GIN_MODE", "debug")

	// Setting up log file
	logFilePath := common.GetEnv("LOG_FILE_PATH", "server.log")
	logFile, err := os.Create(logFilePath)
	if err != nil {
		log.Printf("Failed to create log file: %s", err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// Setting up engine
	gin.SetMode(ginMode)
	engine := gin.Default()

	// Setting up cors configuration
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(common.GetEnv("ALLOW_ORIGIN_DOMAINS", "http://127.0.0.1"), ","),
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Needs to change if server is behind a proxy server (for security purpose)
	if err := engine.SetTrustedProxies(nil); err != nil {
		log.Fatalf("unable to set trusted proxies: %s", err.Error())
	}

	// Starting engine
	if err := engine.Run(":" + ginPort); err != nil {
		log.Fatalf("Error starting engine: %s", err.Error())
	}
}
