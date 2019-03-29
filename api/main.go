package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hatch-group/keywordss-api/api/controller"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Users []User

func main() {
	r := gin.Default()

	router := r.Group("/api")
	{
		story := &controller.Story{}
		dburl := os.Getenv("MYSQL_URL")
		db, err := sqlx.Connect("mysql", dburl)
		if err != nil {
			fmt.Println("mysql connect error")
		}
		story.DB = db
		router.GET("/story", story.IndexGet)
	}

	r.Run(":8080")
}
