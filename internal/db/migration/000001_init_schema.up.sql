CREATE TABLE "members" (
  "member_id" uuid PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "phone_number" varchar,
  "gender" varchar,
  "insurance" uuid,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "providers" (
  "provider_id" uuid PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "specialization" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "insurance" (
  "insurer_id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "visits" (
  "visit_id" uuid PRIMARY KEY,
  "provider_id" uuid,
  "members_id" uuid,
  "scheduled_at" timestamptz NOT NULL,
  "completed_at" timestamptz,
  "status" varchar NOT NULL,
  "notes" TEXT,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE INDEX "idx_members_email" ON "members" ("email");

CREATE INDEX "idx_visits_provider_id" ON "visits" ("provider_id");

CREATE INDEX "idx_visits_member_id" ON "visits" ("members_id");

ALTER TABLE "members" ADD FOREIGN KEY ("insurance") REFERENCES "insurance" ("insurer_id");

ALTER TABLE "visits" ADD FOREIGN KEY ("provider_id") REFERENCES "providers" ("provider_id");

ALTER TABLE "visits" ADD FOREIGN KEY ("members_id") REFERENCES "members" ("member_id");
