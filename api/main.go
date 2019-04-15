package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hatch-group/keywordss-api/api/controller"
	"github.com/jmoiron/sqlx"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		story := &controller.Story{}
		dburl := os.Getenv("MYSQL_URL")
		db, err := sqlx.Connect("mysql", dburl)
		if err != nil {
			fmt.Println("mysql connect error")
		}
		story.DB = db
		api.GET("/story", story.IndexGet)
		api.GET("/story/:id", story.ShowItem)
		api.POST("/story", story.Post)
		api.PUT("/story", story.Edit)
		api.DELETE("/story", story.Delete)

		api.GET("/user/:user_id/story", story.IndexMyPost)
	}

	r.Run(":8080")
}
