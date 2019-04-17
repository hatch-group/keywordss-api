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

	tx, err := s.DB.Beginx()
	if err != nil {
		c.String(500, "start transaction fail")
	}

	result, err := user.SignUp(tx)
	if err != nil {
		c.String(500, "signup error")
	}
	if err := tx.Commit(); err != nil {
		c.String(500, "tx commit error")
	}
	user.ID, err = result.LastInsertId()

	c.String(200, "登録が完了しました")
}

func (u *User) UserSignIn(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)

	result, err := user.SignIn(s.DB)
	if err != nil {
		c.String(500, "Signin error")
	}
	c.JSON(200, gin.H{
		"name": result.Name,
		"status": "ログインに成功しました"
	})
}
