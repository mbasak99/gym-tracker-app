-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
    name,
    dob,
    joined
) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUserName :exec
UPDATE users
SET name = $2
WHERE id = $1;

-- name: UpdateUserDOB :exec
UPDATE users
SET dob = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
