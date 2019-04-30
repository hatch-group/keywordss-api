package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hatch-group/keywordss-api/api/controller"
	"github.com/jmoiron/sqlx"
)

// type Server struct {
// 	db *sqlx.DB
// }

// func (s *Server) init() {
// 	dburl := os.Getenv("MYSQL_URL")
// 	db, err := sqlx.Connect("mysql", dburl)
// 	if err != nil {
// 		fmt.Println("mysql connect error")
// 	}
// 	s.db = db
// }

// func (s *Server) route() {
// 	r := gin.Default()

// 	api := r.Group("/api")
// 	{
// 		story := &controller.Story{}
// 		user := &controller.User{}
// 		story.DB = s.db
// 		user.DB = s.db
// 		api.GET("/stories", story.IndexGet)
// 		api.GET("/stories/:id", story.ShowItem)
// 		api.POST("/stories", story.Post)
// 		api.PUT("/stories", story.Edit)
// 		api.DELETE("/stories", story.Delete)

// 		api.GET("/user/:user_id/stories", story.IndexMyPost)
// 		api.POST("/users/signup", user.UserSignUp)
// 		api.POST("/users/signin", user.UserSignIn)
// 	}

// 	r.Run(":8080")
// }

func main() {
	r := gin.Default()
	dburl := os.Getenv("MYSQL_URL")
	db, err := sqlx.Connect("mysql", dburl)
	if err != nil {
		fmt.Println("mysql connect error")
	}

	api := r.Group("/api")
	{
		story := &controller.Story{DB: db}
		user := &controller.User{DB: db}

		api.GET("/stories", story.IndexGet)
		api.GET("/stories/:id", story.ShowItem)
		api.POST("/stories", story.Post)
		api.PUT("/stories", story.Edit)
		api.DELETE("/stories", story.Delete)

		api.GET("/user/:user_id/stories", story.IndexMyPost)
		api.POST("/users/signup", user.UserSignUp)
		api.POST("/users/signin", user.UserSignIn)
	}

	r.Run(":8080")
}
