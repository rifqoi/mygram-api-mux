// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const deleteUserByEmail = `-- name: DeleteUserByEmail :exec
DELETE FROM users
WHERE email = $1
`

func (q *Queries) DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, deleteUserByEmail, email)
	return err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUserByID, id)
	return err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT
    id,
    username,
    password,
    email,
    age
FROM
  users
WHERE
    email = $1
LIMIT 1
`

type FindUserByEmailRow struct {
	ID       int32
	Username string
	Password string
	Email    string
	Age      int32
}

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (FindUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, findUserByEmail, email)
	var i FindUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Age,
	)
	return i, err
}

const findUserByID = `-- name: FindUserByID :one
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

type FindUserByIDRow struct {
	ID       int32
	Username string
	Password string
	Email    string
	Age      int32
}

func (q *Queries) FindUserByID(ctx context.Context, id int32) (FindUserByIDRow, error) {
	row := q.db.QueryRow(ctx, findUserByID, id)
	var i FindUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Age,
	)
	return i, err
}

const findUserById = `-- name: FindUserById :one
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

type FindUserByIdRow struct {
	ID       int32
	Username string
	Password string
	Email    string
	Age      int32
}

func (q *Queries) FindUserById(ctx context.Context, id int32) (FindUserByIdRow, error) {
	row := q.db.QueryRow(ctx, findUserById, id)
	var i FindUserByIdRow
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

const updateUserByID = `-- name: UpdateUserByID :one
UPDATE users 
SET 
    email = COALESCE($1, email),
    username = COALESCE($2, username),
    password = COALESCE($3, password),
    age = COALESCE($4, age)
WHERE id = $5
RETURNING id, username, password, email, age, created_at, updated_at
`

type UpdateUserByIDParams struct {
	Email    sql.NullString
	Username sql.NullString
	Password sql.NullString
	Age      sql.NullInt32
	ID       int32
}

func (q *Queries) UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserByID,
		arg.Email,
		arg.Username,
		arg.Password,
		arg.Age,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Age,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
