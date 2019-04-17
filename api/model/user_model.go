package model

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func (u *User) SignUp(tx *sqlx.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	insert into users (name) 
	values(?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(u.Name)
}

func (u *User) SignIn(db *sqlx.DB) (user User, err error) {
	if err := db.Select(&user, "SELECT * FROM stories where name = ?", u.Name); err != nil {
		return User{}, err
	}
	return user, nil
}
