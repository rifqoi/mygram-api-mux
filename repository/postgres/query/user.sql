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

-- name: GetUserById :one
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
