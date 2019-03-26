package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	user := os.Getenv("MYSQL_USER")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": user,
		})
	})
	r.Run(":8080")
}
