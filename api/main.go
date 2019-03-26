package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Users []User

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		dburl := os.Getenv("MYSQL_URL")
		db, err := sqlx.Connect("mysql", dburl)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "mysql connect error",
			})
		}
		rows, err := db.Queryx("SELECT * FROM users")
		if err != nil {
			c.JSON(500, gin.H{
				"message": "db select error",
			})
		}
		var user User
		var users Users
		for rows.Next() {
			err := rows.StructScan(&user)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "user bind error",
				})
			}
			users = append(users, user)
		}
		c.JSON(200, gin.H{
			"users": users,
		})
	})
	r.Run(":8080")
}
