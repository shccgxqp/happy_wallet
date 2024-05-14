CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
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

ALTER TABLE "teams" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "team_members" ADD FOREIGN KEY ("team_id") REFERENCES "teams" ("id") ON DELETE CASCADE;

ALTER TABLE "team_members" ADD FOREIGN KEY ("linked_user_id") REFERENCES "users" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("team_id") REFERENCES "teams" ("id") ON DELETE CASCADE;

ALTER TABLE "expense_details" ADD FOREIGN KEY ("member_id") REFERENCES "team_members" ("id");

ALTER TABLE "expense_details" ADD FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id") ON DELETE CASCADE;
