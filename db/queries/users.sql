-- name: CreateUser :one
INSERT INTO users (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users where id = $1 limit 1;


-- name: GetUserByName :one
SELECT * FROM users where name LIKE $1 limit 1;

-- name: UpdateUserByName :one
UPDATE users SET name = $2 WHERE name LIKE $1 RETURNING *;

-- name: UpdateUserById :one
UPDATE users SET name = $2 WHERE id = $1 RETURNING *;

-- name: DeleteUserById :exec
DELETE FROM users
WHERE id = $1;

-- name: DeleteUserByName :exec
DELETE FROM users
WHERE name like $1;