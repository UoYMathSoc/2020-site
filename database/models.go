// Code generated by sqlc. DO NOT EDIT.

package database

import (
	"database/sql"
	"time"
)

type Event struct {
	ID          int32
	Name        string
	StartTime   time.Time
	EndTime     sql.NullTime
	Location    sql.NullString
	Description sql.NullString
}

type Position struct {
	ID          int32
	Name        sql.NullString
	Alias       string
	Ordering    int16
	Description sql.NullString
	Executive   sql.NullBool
}

type Post struct {
	ID      int32
	Date    time.Time
	Title   string
	Body    sql.NullString
	EventID int32
	Link    sql.NullString
}

type User struct {
	ID       int32
	Username string
	Name     string
	Bio      sql.NullString
}

type UsersCommittee struct {
	ID          int32
	UserID      int32
	CommitteeID int32
	FromDate    time.Time
	TillDate    sql.NullTime
	Outgoing    bool
}

type UsersPass struct {
	ID       int32
	Password string
}
