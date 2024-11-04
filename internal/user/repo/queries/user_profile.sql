-- name: CreateUserProfile :one
INSERT INTO "user".user_profile (
    user_id,
    user_email,
    user_nickname
) VALUES ($1, $2, $3) RETURNING *;
