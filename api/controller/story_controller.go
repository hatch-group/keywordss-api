package controller

import (
	"strconv"
	"time"

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
		c.JSON(500, gin.H{
			"message": "failed to get a;; index",
		})
	}
	total := len(stories)
	c.JSON(200, gin.H{
		"total":   total,
		"stories": stories,
	})
}

func (s *Story) ShowItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message":"failed conversion string to int",
		})
	}
	story, err := model.Show(s.DB, id)
	if err != nil {
		c.JSON(500, gin.H{
			"message":"failed model connection",
		})
	}
	c.JSON(200, story)
}

func (s *Story) Post(c *gin.Context) {
	var story model.Story
	c.BindJSON(&story)

	time := time.Now()
	story.PostedTime = &time

	tx, err := s.DB.Beginx()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "start transaction fail",
		})
	}

	result, err := story.Insert(tx)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "insert error",
		})
	}
	if err := tx.Commit(); err != nil {
		c.JSON(500, gin.H{
			"message": "tx commit error",
		})
	}
	story.ID, err = result.LastInsertId()

	c.JSON(200, gin.H{
		"message": "ポストが完了しました",
	})
}

func (s *Story) Edit(c *gin.Context) {
	var story model.Story
	c.BindJSON(&story)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed conversion string to int",
		})
	}

	tx, err := s.DB.Beginx()
	if err != nil {
		c.JSON(500, gin.H{
			"message":  "start transaction fail",
		})
	}

	_, err = story.Edit(tx, id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "edit error"),
		})
	}

	if err := tx.Commit(); err != nil {
		c.JSON(500, gin.H{
			"message": "tx commit error",
		})
	}

	c.JSON(200, gin.H{
		"message": "編集が完了しました！",
	})
}

func (s *Story) Delete(c *gin.Context) {
	var story model.Story
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message":  "failed conversion string to int",
		})
	}

	tx, err := s.DB.Beginx()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "start transaction fail",
		})
	}

	_, err = story.Delete(tx, id)
	if err != nil {
		c.JSON(500, gin.H{
			"message":  "delete error",
		})
	}

	if err := tx.Commit(); err != nil {
		c.JSON(500, gin.H{
			"message":  "tx commit error",
		})
	}

	c.JSON(200, gin.H{
		"message": "消去が完了しました！",
	})
}

func (s *Story) IndexMyPost(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message":"failed conversion string to int",
		})
	}
	stories, err := model.IndexMyPost(s.DB, user_id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed model connection",
		})
	}
	c.JSON(200, stories)
}
