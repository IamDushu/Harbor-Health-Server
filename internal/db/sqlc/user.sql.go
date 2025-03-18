// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    user_id,
    email, 
    first_name, 
    last_name,
    phone_number
) VALUES (
   $1, $2, $3, $4, $5
) RETURNING user_id, email, first_name, last_name, phone_number, is_onboarded, created_at
`

type CreateUserParams struct {
	UserID      uuid.UUID `json:"user_id"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.UserID,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.IsOnboarded,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT user_id, email, first_name, last_name, phone_number, is_onboarded, created_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.IsOnboarded,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users 
SET first_name = $1, last_name = $2, phone_number = $3, is_onboarded = $4
WHERE email = $5
RETURNING user_id, email, first_name, last_name, phone_number, is_onboarded, created_at
`

type UpdateUserParams struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	IsOnboarded bool   `json:"is_onboarded"`
	Email       string `json:"email"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.IsOnboarded,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.IsOnboarded,
		&i.CreatedAt,
	)
	return i, err
}
