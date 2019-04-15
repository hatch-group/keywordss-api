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
		c.String(500, "全件取得に失敗しておりますぞ")
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
		c.String(500, "failed strvonv")
	}
	story, err := model.Show(s.DB, id)
	c.JSON(200, story)
}

func (s *Story) Post(c *gin.Context) {
	var story model.Story
	c.BindJSON(&story)

	time := time.Now()
	story.PostedTime = &time

	tx, err := s.DB.Beginx()
	if err != nil {
		c.String(500, "failed strvonv")
	}

	result, err := story.Insert(tx)
	if err != nil {
		c.String(500, "insert error")
	}
	if err := tx.Commit(); err != nil {
		c.String(500, "tx commit error")
	}
	story.ID, err = result.LastInsertId()

	c.String(200, "ポストが完了しました")
}

func (s *Story) Edit(c *gin.Context) {
	var story model.Story
	c.BindJSON(&story)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(500, "failed strvonv")
	}

	tx, err := s.DB.Beginx()
	if err != nil {
		c.String(500, "start transaction fail")
	}

	_, err = story.Edit(tx, id)
	if err != nil {
		c.String(500, "edit error")
	}

	if err := tx.Commit(); err != nil {
		c.String(500, "tx commit error")
	}

	c.String(200, "編集が完了しました！")
}

func (s *Story) Delete(c *gin.Context) {
	var story model.Story
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(500, "failed strvonv")
	}

	tx, err := s.DB.Beginx()
	if err != nil {
		c.String(500, "start transaction fail")
	}

	_, err = story.Delete(tx, id)
	if err != nil {
		c.String(500, "delete error")
	}

	if err := tx.Commit(); err != nil {
		c.String(500, "tx commit error")
	}

	c.String(200, "消去が完了しました！")
}

func (s *Story) IndexMyPost(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.String(500, "failed strvonv")
	}
	stories, err := model.IndexMyPost(s.DB, user_id)
	c.JSON(200, stories)
}
