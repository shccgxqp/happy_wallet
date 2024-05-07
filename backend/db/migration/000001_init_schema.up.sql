CREATE TABLE "e_group" (
  "id" bigserial PRIMARY KEY,
  "code" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "expenses" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "amount" float NOT NULL,
  "user_id" bigint NOT NULL,
  "group_code" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "share" (
  "id" bigserial PRIMARY KEY,
  "expense_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "share_type" varchar NOT NULL,
  "share_value" float,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "e_group" ("code");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "expenses" ("user_id");

CREATE INDEX ON "expenses" ("group_code");

CREATE INDEX ON "expenses" ("name", "amount");

CREATE INDEX ON "share" ("expense_id");

CREATE INDEX ON "share" ("user_id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("group_code") REFERENCES "e_group" ("code");

ALTER TABLE "share" ADD FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id");

ALTER TABLE "share" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");