CREATE TABLE "users" (
  "user_id" uuid PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "is_onboarded" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "members" (
  "member_id" uuid PRIMARY KEY,
  "user_id" uuid UNIQUE NOT NULL,
  "gender" varchar NOT NULL,
  "date_of_birth" date NOT NULL,
  "insurance" varchar NOT NULL,
  "address_line_one" varchar NOT NULL,
  "address_line_two" varchar NOT NULL,
  "accepted_terms" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "providers" (
  "provider_id" uuid PRIMARY KEY,
  "user_id" uuid UNIQUE NOT NULL,
  "credentials" varchar NOT NULL,
  "specialization" varchar NOT NULL,
  "is_available" boolean NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "provider_availability" (
  "availability_id" uuid PRIMARY KEY,
  "provider_id" uuid UNIQUE NOT NULL,
  "day_of_week" int NOT NULL,
  "start_time" time NOT NULL,
  "end_time" time NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "locations" (
  "location_id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "address" text NOT NULL,
  "latitude" decimal(9,6) NOT NULL,
  "longitude" decimal(9,6) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "provider_locations" (
  "provider_id" uuid NOT NULL,
  "location_id" uuid NOT NULL
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

CREATE UNIQUE INDEX unique_provider_availability
ON provider_availability (provider_id, day_of_week, start_time, end_time);

CREATE UNIQUE INDEX unique_provider_location
ON provider_locations (provider_id, location_id);

COMMENT ON COLUMN "provider_availability"."day_of_week" IS '0 = Sunday, 6 = Saturday';

ALTER TABLE "members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "providers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "provider_availability" ADD FOREIGN KEY ("provider_id") REFERENCES "providers" ("provider_id");

ALTER TABLE "provider_locations" ADD FOREIGN KEY ("provider_id") REFERENCES "providers" ("provider_id");

ALTER TABLE "provider_locations" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("location_id");

ALTER TABLE "visits" ADD FOREIGN KEY ("provider_id") REFERENCES "providers" ("provider_id");

ALTER TABLE "visits" ADD FOREIGN KEY ("members_id") REFERENCES "members" ("member_id");
