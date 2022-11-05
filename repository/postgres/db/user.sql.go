// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package db

import (
	"context"
)

const getUserById = `-- name: GetUserById :one
SELECT
    id,
    username,
    password,
    email,
    age
FROM
  users
WHERE
  id = $1
LIMIT 1
`

type GetUserByIdRow struct {
	ID       int32
	Username string
	Password string
	Email    string
	Age      int32
}

func (q *Queries) GetUserById(ctx context.Context, id int32) (GetUserByIdRow, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i GetUserByIdRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Age,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO users (
    username,
    password,
    email,
    age,
    created_at,
    updated_at
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    current_timestamp,
    current_timestamp
)
`

type InsertUserParams struct {
	Username string
	Password string
	Email    string
	Age      int32
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.db.Exec(ctx, insertUser,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.Age,
	)
	return err
}