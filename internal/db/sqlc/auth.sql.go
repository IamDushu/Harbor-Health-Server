// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: auth.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createVerifyRecord = `-- name: CreateVerifyRecord :one
INSERT INTO email_verification (
    verification_id,
    email,
    token,
    hashed_otp,
    purpose,
    attempts,
    expires_at,
    valid
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING verification_id, email, token, hashed_otp, purpose, attempts, expires_at, valid, created_at
`

type CreateVerifyRecordParams struct {
	VerificationID uuid.UUID `json:"verification_id"`
	Email          string    `json:"email"`
	Token          string    `json:"token"`
	HashedOtp      string    `json:"hashed_otp"`
	Purpose        string    `json:"purpose"`
	Attempts       int32     `json:"attempts"`
	ExpiresAt      time.Time `json:"expires_at"`
	Valid          bool      `json:"valid"`
}

func (q *Queries) CreateVerifyRecord(ctx context.Context, arg CreateVerifyRecordParams) (EmailVerification, error) {
	row := q.db.QueryRowContext(ctx, createVerifyRecord,
		arg.VerificationID,
		arg.Email,
		arg.Token,
		arg.HashedOtp,
		arg.Purpose,
		arg.Attempts,
		arg.ExpiresAt,
		arg.Valid,
	)
	var i EmailVerification
	err := row.Scan(
		&i.VerificationID,
		&i.Email,
		&i.Token,
		&i.HashedOtp,
		&i.Purpose,
		&i.Attempts,
		&i.ExpiresAt,
		&i.Valid,
		&i.CreatedAt,
	)
	return i, err
}

const getVerifyRecord = `-- name: GetVerifyRecord :one
SELECT verification_id, email, token, hashed_otp, purpose, attempts, expires_at, valid, created_at FROM email_verification
WHERE email = $1
AND valid = TRUE
AND purpose = $2
`

type GetVerifyRecordParams struct {
	Email   string `json:"email"`
	Purpose string `json:"purpose"`
}

func (q *Queries) GetVerifyRecord(ctx context.Context, arg GetVerifyRecordParams) (EmailVerification, error) {
	row := q.db.QueryRowContext(ctx, getVerifyRecord, arg.Email, arg.Purpose)
	var i EmailVerification
	err := row.Scan(
		&i.VerificationID,
		&i.Email,
		&i.Token,
		&i.HashedOtp,
		&i.Purpose,
		&i.Attempts,
		&i.ExpiresAt,
		&i.Valid,
		&i.CreatedAt,
	)
	return i, err
}

const getVerifyRecordOnToken = `-- name: GetVerifyRecordOnToken :one
SELECT verification_id, email, token, hashed_otp, purpose, attempts, expires_at, valid, created_at FROM email_verification
WHERE token = $1
`

func (q *Queries) GetVerifyRecordOnToken(ctx context.Context, token string) (EmailVerification, error) {
	row := q.db.QueryRowContext(ctx, getVerifyRecordOnToken, token)
	var i EmailVerification
	err := row.Scan(
		&i.VerificationID,
		&i.Email,
		&i.Token,
		&i.HashedOtp,
		&i.Purpose,
		&i.Attempts,
		&i.ExpiresAt,
		&i.Valid,
		&i.CreatedAt,
	)
	return i, err
}

const updateVerifyRecordAttempt = `-- name: UpdateVerifyRecordAttempt :one
UPDATE email_verification
SET attempts = attempts + 1
WHERE verification_id = $1
RETURNING verification_id, email, token, hashed_otp, purpose, attempts, expires_at, valid, created_at
`

func (q *Queries) UpdateVerifyRecordAttempt(ctx context.Context, verificationID uuid.UUID) (EmailVerification, error) {
	row := q.db.QueryRowContext(ctx, updateVerifyRecordAttempt, verificationID)
	var i EmailVerification
	err := row.Scan(
		&i.VerificationID,
		&i.Email,
		&i.Token,
		&i.HashedOtp,
		&i.Purpose,
		&i.Attempts,
		&i.ExpiresAt,
		&i.Valid,
		&i.CreatedAt,
	)
	return i, err
}

const updateVerifyRecordInvalid = `-- name: UpdateVerifyRecordInvalid :one
UPDATE email_verification
SET valid = FALSE
WHERE verification_id = $1
RETURNING verification_id, email, token, hashed_otp, purpose, attempts, expires_at, valid, created_at
`

func (q *Queries) UpdateVerifyRecordInvalid(ctx context.Context, verificationID uuid.UUID) (EmailVerification, error) {
	row := q.db.QueryRowContext(ctx, updateVerifyRecordInvalid, verificationID)
	var i EmailVerification
	err := row.Scan(
		&i.VerificationID,
		&i.Email,
		&i.Token,
		&i.HashedOtp,
		&i.Purpose,
		&i.Attempts,
		&i.ExpiresAt,
		&i.Valid,
		&i.CreatedAt,
	)
	return i, err
}
