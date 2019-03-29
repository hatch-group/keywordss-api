package model

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Story struct {
	ID         int64      `db:"id"`
	Title      string     `db:"title"`
	Body       string     `db:"body"`
	PostedTime *time.Time `db:"posted_time"`
	Keywords   string     `db:"keywords"`
	UserId     int64      `db:"user_id"`
}

func StoriesAll(db *sqlx.DB) (stories []Story, err error) {
	if err := db.Select(&stories, "SELECT * FROM stories"); err != nil {
		return nil, err
	}
	return stories, nil
}
