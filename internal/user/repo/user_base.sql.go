// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user_base.sql

package repo

import (
	"context"
)

const checkUserBaseExists = `-- name: CheckUserBaseExists :one
SELECT COUNT(*)
FROM "user"."user_base"
WHERE "user_email" = $1
`

func (q *Queries) CheckUserBaseExists(ctx context.Context, userEmail string) (int64, error) {
	row := q.db.QueryRow(ctx, checkUserBaseExists, userEmail)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUserBase = `-- name: CreateUserBase :one
INSERT INTO "user".user_base (
    user_email,
    user_password,
    user_salt
) VALUES ($1, $2, $3) RETURNING user_id, user_email, user_password, user_salt, user_login_time, user_logout_time, user_login_ip, user_created_at, user_updated_at
`

type CreateUserBaseParams struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserSalt     string `json:"user_salt"`
}

func (q *Queries) CreateUserBase(ctx context.Context, arg CreateUserBaseParams) (UserUserBase, error) {
	row := q.db.QueryRow(ctx, createUserBase, arg.UserEmail, arg.UserPassword, arg.UserSalt)
	var i UserUserBase
	err := row.Scan(
		&i.UserID,
		&i.UserEmail,
		&i.UserPassword,
		&i.UserSalt,
		&i.UserLoginTime,
		&i.UserLogoutTime,
		&i.UserLoginIp,
		&i.UserCreatedAt,
		&i.UserUpdatedAt,
	)
	return i, err
}
