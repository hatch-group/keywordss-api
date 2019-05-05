package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hatch-group/keywordss-api/api/model"
	"github.com/jmoiron/sqlx"
)

type User struct {
	DB *sqlx.DB
}

func (u *User) UserSignUp(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)

	tx, err := u.DB.Beginx()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "start transaction fail",
		})
	}

	result, err := user.SignUp(tx)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "signup error",
		})
	}
	if err := tx.Commit(); err != nil {
		c.JSON(500, gin.H{
			"message": "tx commit error",
		})
	}
	user.ID, err = result.LastInsertId()

	c.JSON(200, gin.H{
		"message": "投稿に成功しました",
	})
}

func (u *User) UserSignIn(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)

	result, err := user.SignIn(u.DB)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Signin error",
		})
	}
	c.JSON(200, gin.H{
		"name":   result.Name,
		"status": "ログインに成功しました",
	})
}
