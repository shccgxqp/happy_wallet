-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2024-05-17T08:53:47.772Z

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "is_email_verified" boolean NOT NULL DEFAULT false,
  "password_change_at" timestamp NOT NULL DEFAULT '0001-01-01',
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "verify_emails" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "expired_at" timestamp NOT NULL DEFAULT (now() + interval '15 minutes')
);

CREATE TABLE "teams" (
  "id" bigserial PRIMARY KEY,
  "owner" bigserial NOT NULL,
  "team_name" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "team_members" jsonb,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "team_members" (
  "id" bigserial PRIMARY KEY,
  "team_id" bigint,
  "member_name" varchar NOT NULL,
  "linked_user_id" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "expenses" (
  "id" bigserial PRIMARY KEY,
  "team_id" bigint,
  "goal" varchar NOT NULL,
  "amount" decimal(10,2) NOT NULL,
  "currency" varchar NOT NULL,
  "sharing_method" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "expense_details" (
  "id" bigserial PRIMARY KEY,
  "expense_id" bigint,
  "member_id" bigint,
  "actual_amount" decimal(10,2) DEFAULT 0,
  "shared_amount" decimal(10,2) DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "teams" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "team_members" ADD FOREIGN KEY ("team_id") REFERENCES "teams" ("id") ON DELETE CASCADE;

ALTER TABLE "team_members" ADD FOREIGN KEY ("linked_user_id") REFERENCES "users" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("team_id") REFERENCES "teams" ("id") ON DELETE CASCADE;

ALTER TABLE "expense_details" ADD FOREIGN KEY ("member_id") REFERENCES "team_members" ("id");

ALTER TABLE "expense_details" ADD FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id") ON DELETE CASCADE;

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
