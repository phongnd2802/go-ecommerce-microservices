-- name: CreateUserVerify :one
INSERT INTO "user"."user_verify" (
    "verify_otp",
    "verify_key",
    "verify_key_hash"
) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUserVerifyByKeyHash :one
SELECT *
FROM "user"."user_verify"
WHERE "verify_key_hash" = $1;

-- name: UpdateUserVerify :one
UPDATE "user"."user_verify"
SET 
    is_verified = COALESCE(sqlc.narg(is_verified), is_verified),
    is_deleted = COALESCE(sqlc.narg(is_deleted), is_deleted)
WHERE verify_key_hash = sqlc.arg(verify_key_hash)
RETURNING *;
