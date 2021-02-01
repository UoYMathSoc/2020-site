package models

import (
	"database/sql"
	"fmt"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/structs"
)

type SessionStore struct {
	querier database.Querier
}

func NewSessionStore(querier database.Querier) *SessionStore {
	return &SessionStore{querier: querier}
}

func NewSessionFromConfig(db structs.Database) *SessionStore {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Name)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return &SessionStore{querier: database.New(conn)}
}
