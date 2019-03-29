package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hatch-group/keywordss-api/api/model"
	"github.com/jmoiron/sqlx"
)

type Story struct {
	DB *sqlx.DB
}

func (s *Story) IndexGet(c *gin.Context) {
	stories, err := model.StoriesAll(s.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"message": "何でえ",
		})
	}
	total := len(stories)
	c.JSON(200, gin.H{
		"total":   total,
		"stories": stories,
	})
}
