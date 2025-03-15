CREATE TABLE "users" (
  "user_id" uuid PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "phone_number" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "members" (
  "member_id" uuid PRIMARY KEY,
  "user_id" uuid UNIQUE NOT NULL,
  "gender" varchar NOT NULL,
  "insurance" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "providers" (
  "provider_id" uuid PRIMARY KEY,
  "user_id" uuid UNIQUE NOT NULL,
  "specialization" varchar NOT NULL,
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

CREATE TABLE "email_verification" (
  "verification_id" uuid PRIMARY KEY,
  "email" varchar NOT NULL,
  "token" text UNIQUE NOT NULL,
  "hashed_otp" varchar NOT NULL,
  "purpose" varchar NOT NULL,
  "attempts" integer NOT NULL DEFAULT 0,
  "expires_at" timestamptz NOT NULL,
  "valid" boolean NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "sessions" (
  "session_id" uuid PRIMARY KEY,
  "email" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE INDEX "idx_users_email" ON "users" ("email");

CREATE INDEX "idx_visits_provider_id" ON "visits" ("provider_id");

CREATE INDEX "idx_visits_member_id" ON "visits" ("members_id");

CREATE UNIQUE INDEX email_purpose_valid_key ON email_verification (email, purpose) WHERE valid = true;

ALTER TABLE "members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "members" ADD FOREIGN KEY ("insurance") REFERENCES "insurance" ("insurer_id");

ALTER TABLE "providers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "visits" ADD FOREIGN KEY ("provider_id") REFERENCES "providers" ("provider_id");

ALTER TABLE "visits" ADD FOREIGN KEY ("members_id") REFERENCES "members" ("member_id");
