// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"
)

type User struct {
	ID        int32
	Username  string
	Password  string
	Email     string
	Age       int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
