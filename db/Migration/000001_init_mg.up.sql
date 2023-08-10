CREATE TABLE "User" (
  "id" SERIAL PRIMARY KEY,
  "username" VARCHAR(50) NOT NULL,
  "email" VARCHAR(100) NOT NULL,
  "password_hash" VARCHAR(100),
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);

CREATE TABLE "Question" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INTEGER NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);

CREATE TABLE "Answer" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INTEGER NOT NULL,
  "question_id" INTEGER NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00'
);
ALTER TABLE "User" ADD CONSTRAINT  "unique_email_constraint"  UNIQUE ("email");
ALTER TABLE "User" ADD CONSTRAINT  "unique_username_constraint"  UNIQUE ("username");

ALTER TABLE "Question" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "Answer" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "Answer" ADD FOREIGN KEY ("question_id") REFERENCES "Question" ("id");
