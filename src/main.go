package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err:= godotenv.Load()
	if err != nil {
		panic("not possible to load env variables")
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	serverPort := os.Getenv("SERVER_PORT")
	r.Run(serverPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
