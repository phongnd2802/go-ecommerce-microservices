-- name: CheckUserBaseExists :one
SELECT COUNT(*)
FROM "user"."user_base"
WHERE "user_email" = $1;

-- name: CreateUserBase :one
INSERT INTO "user".user_base (
    user_email,
    user_password,
    user_salt
) VALUES ($1, $2, $3) RETURNING *;