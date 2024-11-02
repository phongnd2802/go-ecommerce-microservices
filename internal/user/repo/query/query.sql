-- name: CreateUserVerify :one
INSERT INTO "user"."user_verify" (
    "verify_otp",
    "verify_key",
    "verify_key_hash"
) VALUES ($1, $2, $3) RETURNING *;

-- name: CheckUserBaseExists :one
SELECT COUNT(*)
FROM "user"."user_base"
WHERE "user_email" = $1;
