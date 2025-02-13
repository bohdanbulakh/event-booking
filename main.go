package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	server := gin.Default()
	port := os.Getenv("PORT")

	server.Run(":" + port)
}
