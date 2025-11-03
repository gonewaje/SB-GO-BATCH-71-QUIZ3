package repository

import (
	"database/sql"
	"errors"
	"library/structs"
)

func GetUserByUsername(db *sql.DB, username string) (structs.User, error) {
	var u structs.User
	err := db.QueryRow(`SELECT id, username, password FROM users WHERE username=$1`, username).
		Scan(&u.ID, &u.Username, &u.Password)
	if err == sql.ErrNoRows {
		return u, errors.New("user not found")
	}
	return u, err
}
