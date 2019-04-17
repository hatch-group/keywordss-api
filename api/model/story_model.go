package model

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type Story struct {
	ID         int64      `json:"id" db:"id"`
	Title      string     `json:"title" db:"title"`
	Body       string     `json:"body" db:"body"`
	PostedTime *time.Time `json:"posted_time" db:"posted_time"`
	Keywords   string     `json:"keywords" db:"keywords"`
	UserId     int64      `json:"user_id" db:"user_id"`
}

func StoriesAll(db *sqlx.DB) (stories []Story, err error) {
	if err := db.Select(&stories, "SELECT * FROM stories"); err != nil {
		return nil, err
	}
	return stories, nil
}

func Show(db *sqlx.DB, id int) (story Story, err error) {
	if err := db.Get(&story, "SELECT * FROM stories where id=?", id); err != nil {
		return Story{}, err
	}
	return story, nil
}

func (s *Story) Insert(tx *sqlx.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	insert into stories (title, body, posted_time, keywords, user_id) 
	values(?, ?, ?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(s.Title, s.Body, s.PostedTime, s.Keywords, s.UserId)
}

func (s *Story) Delete(tx *sqlx.Tx, id int) (sql.Result, error) {
	stmt, err := tx.Prepare(`delete from stories where id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}

func (s *Story) Edit(tx *sqlx.Tx, id int) (sql.Result, error) {
	stmt, err := tx.Prepare(`update stories set title=?, body=? where id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(s.Title, s.Body, id)
}

func IndexMyPost(db *sqlx.DB, id int) (stories []Story, err error) {
	if err := db.Select(&stories, "SELECT * FROM stories where user_id = ?", id); err != nil {
		return nil, err
	}
	return stories, nil
}
