// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"
)

type Account struct {
	ID          int32
	UserID      int32
	CategoryID  int32
	Title       string
	Type        string
	Value       int32
	Date        time.Time
	Description string
	CreatedAt   time.Time
}

type Category struct {
	ID          int32
	UserID      int32
	Title       string
	Type        string
	Description string
	CreatedAt   time.Time
}

type User struct {
	ID        int32
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
}
