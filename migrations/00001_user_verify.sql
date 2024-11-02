-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS "user";

CREATE TABLE "user"."user_verify" (
    "verify_id" bigserial PRIMARY KEY,
    "verify_otp" varchar NOT NULL,
    "verify_key" varchar NOT NULL,
    "verify_key_hash" varchar NOT NULL,
    "is_verified" boolean DEFAULT false,
    "is_deleted" boolean DEFAULT false,
    "verify_created_at" timestamp NOT NULL DEFAULT (now()),
    "verify_updated_at" timestamp NOT NULL DEFAULT (now()),
    UNIQUE ("verify_key")
);

CREATE INDEX on "user"."user_verify" ("verify_otp");

CREATE OR REPLACE FUNCTION "user".user_verify_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.verify_updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER update_user_verify_updated_at
BEFORE UPDATE ON "user".user_verify
FOR EACH ROW 
EXECUTE FUNCTION "user".user_verify_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_user_verify_updated_at ON "user".user_verify;

DROP FUNCTION IF EXISTS "user".user_verify_updated_at();

DROP TABLE IF EXISTS "user"."user_verify";
-- +goose StatementEnd
