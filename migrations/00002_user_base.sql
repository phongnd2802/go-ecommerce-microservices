-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user"."user_base" (
    "user_id" bigserial PRIMARY KEY,
    "user_email" varchar NOT NULL,
    "user_password" varchar NOT NULL,
    "user_salt" varchar NOT NULL,
    "user_login_time" timestamp DEFAULT NULL,
    "user_logout_time" timestamp DEFAULT NULL,
    "user_login_ip" varchar DEFAULT NULL,
    "user_created_at" timestamp NOT NULL DEFAULT (now()),
    "user_updated_at" timestamp NOT NULL DEFAULT (now()),
    UNIQUE ("user_email")
);

ALTER TABLE "user"."user_base" ADD FOREIGN KEY ("user_email") REFERENCES "user"."user_verify" ("verify_key");

CREATE INDEX on "user"."user_base" ("user_email");

CREATE OR REPLACE FUNCTION "user".user_base_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.user_updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER update_user_base_updated_at
BEFORE UPDATE ON "user".user_base
FOR EACH ROW 
EXECUTE FUNCTION "user".user_base_updated_at();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_user_base_updated_at ON "user".user_base;

DROP FUNCTION IF EXISTS "user".user_base_updated_at();

DROP TABLE IF EXISTS "user".user_base;
-- +goose StatementEnd
