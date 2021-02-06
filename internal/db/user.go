package db

import "github.com/gocraft/dbr/v2"

type User struct {
	ID             string `db:"id"`
	Name           string `db:"name"`
	HashedPassword string `db:"hashed_password"`
	Email          string `db:"email"`
}

func Insert(runner dbr.SessionRunner, user User) (*User, error) {

	_, err := runner.InsertInto("users").Columns("id", "name", "hashed_password", "email").
		Record(user).Exec()

	if err != nil {
		return nil, err
	}

	return &user, err
}
