-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user"."user_profile" (
    "user_id" bigint PRIMARY KEY,
    "user_email" varchar NOT NULL,
    "user_nickname" varchar NOT NULL,
    "user_avatar" varchar DEFAULT NULL,
    "user_mobile" varchar DEFAULT NULL,
    "user_gender" boolean DEFAULT NULL,
    "user_birthday" date DEFAULT NULL,
    "user_created_at" timestamp NOT NULL DEFAULT (now()),
    "user_updated_at" timestamp NOT NULL DEFAULT (now()),
    UNIQUE ("user_email")
);

ALTER TABLE "user".user_profile ADD FOREIGN KEY ("user_email") REFERENCES "user"."user_base" ("user_email");
ALTER TABLE "user".user_profile ADD FOREIGN KEY ("user_id") REFERENCES "user"."user_base" ("user_id");


CREATE INDEX on "user".user_profile ("user_mobile");
CREATE INDEX on "user".user_profile ("user_email");

CREATE OR REPLACE FUNCTION "user".user_profile_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.user_updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER update_user_profile_updated_at
BEFORE UPDATE ON "user".user_profile
FOR EACH ROW 
EXECUTE FUNCTION "user".user_profile_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_user_profile_updated_at ON "user".user_profile;
DROP FUNCTION IF EXISTS "user".update_user_profile_updated_at;

DROP TABLE IF EXISTS "user".user_profile;
-- +goose StatementEnd
