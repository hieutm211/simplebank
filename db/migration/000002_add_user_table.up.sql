CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz DEFAULT '00001-01-01 00:00:00Z',
  "created_at" timestamptz DEFAULT (now())
);

ALTER TABLE "accounts" ADD CONSTRAINT "accounts_owner_currency_key" UNIQUE("owner", "currency");
ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
