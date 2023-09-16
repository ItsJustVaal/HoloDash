package main

import (
	"log"

	"github.com/ItsJustVaal/Holodash/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Establish Server
// Routes
// Handlers

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	server := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	server.GET("/", routes.GetAllChannels)
	server.POST("/", routes.AddChannelToDB)

	server.Use(cors.New(config))
	server.Run()
}
