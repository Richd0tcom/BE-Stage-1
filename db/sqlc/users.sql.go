// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    name
) VALUES (
    $1
) RETURNING id, name
`

func (q *Queries) CreateUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, name)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteUserById = `-- name: DeleteUserById :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserById, id)
	return err
}

const deleteUserByName = `-- name: DeleteUserByName :exec
DELETE FROM users
WHERE name like $1
`

func (q *Queries) DeleteUserByName(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteUserByName, name)
	return err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name FROM users where id = $1 limit 1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT id, name FROM users where name LIKE $1 limit 1
`

func (q *Queries) GetUserByName(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, name)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const updateUserById = `-- name: UpdateUserById :one
UPDATE users SET name = $2 WHERE id = $1 RETURNING id, name
`

type UpdateUserByIdParams struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (q *Queries) UpdateUserById(ctx context.Context, arg UpdateUserByIdParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserById, arg.ID, arg.Name)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const updateUserByName = `-- name: UpdateUserByName :one
UPDATE users SET name = $2 WHERE name LIKE $1 RETURNING id, name
`

type UpdateUserByNameParams struct {
	Name   string `json:"name"`
	Name_2 string `json:"name_2"`
}

func (q *Queries) UpdateUserByName(ctx context.Context, arg UpdateUserByNameParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserByName, arg.Name, arg.Name_2)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
