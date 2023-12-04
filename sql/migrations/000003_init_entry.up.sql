CREATE TABLE "entries" (
  "id" varchar PRIMARY KEY,
  "account_id" varchar NOT NULL,
  "transaction_type" varchar NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

CREATE INDEX ON "entries" ("account_id");