-- name: InsertUser :exec
INSERT INTO users (
    username,
    password,
    email,
    age,
    created_at,
    updated_at
)
VALUES (
    @username,
    @password,
    @email,
    @age,
    current_timestamp,
    current_timestamp
);

-- name: FindUserById :one
SELECT
    id,
    username,
    password,
    email,
    age
FROM
  users
WHERE
  id = @id
LIMIT 1;

-- name: FindUserByEmail :one
SELECT
    id,
    username,
    password,
    email,
    age
FROM
  users
WHERE
    email = @email
LIMIT 1;

-- name: FindUserByID :one
SELECT
    id,
    username,
    password,
    email,
    age
FROM
  users
WHERE
    id = @id
LIMIT 1;

-- name: DeleteUserByEmail :exec
DELETE FROM users
WHERE email = @email;

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = @id; 

-- name: UpdateUserByID :one
UPDATE users 
SET 
    email = COALESCE(sqlc.narg(email), email),
    username = COALESCE(sqlc.narg(username), username),
    password = COALESCE(sqlc.narg(password), password),
    age = COALESCE(sqlc.narg(age), age)
WHERE id = @id
RETURNING *;
