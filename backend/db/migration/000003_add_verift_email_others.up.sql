CREATE TABLE "verify_emails" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "expired_at" timestamp NOT NULL DEFAULT (now() + interval '15 minutes')
);

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users" ADD COLUMN "is_email_verified" boolean NOT NULL DEFAULT false;

ALTER TABLE "users" ADD COLUMN "password_change_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z');

ALTER TABLE "sessions" ADD COLUMN "username" varchar NOT NULL;