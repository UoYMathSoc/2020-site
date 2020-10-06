package models

import (
	"database/sql"
	"fmt"

	"github.com/UoYMathSoc/2020-site/database"
	"github.com/UoYMathSoc/2020-site/structs"
)

type Session struct {
	querier database.Querier
}

func NewSession(querier database.Querier) *Session {
	return &Session{querier: querier}
}

func NewSessionFromConfig(db structs.Database) *Session {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Name)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return &Session{querier: database.New(conn)}
}
