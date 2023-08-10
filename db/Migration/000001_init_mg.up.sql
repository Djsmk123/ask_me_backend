CREATE TABLE "User" (
  "id" SERIAL PRIMARY KEY,
  "username" VARCHAR(50) NOT NULL,
  "email" VARCHAR(100) NOT NULL,
  "provider" VARCHAR(50) NOT NULL DEFAULT 'password-based',
  "password_hash" VARCHAR(100),
  "public_profile_image" VARCHAR(200) DEFAULT 'https://xsgames.co/randomusers/assets/avatars/pixel/1.jpg' NOT NULL,
  "private_profile_image" VARCHAR(200) DEFAULT 'https://xsgames.co/randomusers/assets/avatars/pixel/1.jpg' NOT NULL,
  "created_at" timestamptz DEFAULT now() NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "Question" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INTEGER NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" timestamptz DEFAULT now() NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE "Answer" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INTEGER NOT NULL,
  "question_id" INTEGER NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" timestamptz DEFAULT now() NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT now()
);

ALTER TABLE "User" ADD CONSTRAINT "unique_email_constraint" UNIQUE ("email");
ALTER TABLE "User" ADD CONSTRAINT "unique_username_constraint" UNIQUE ("username");

ALTER TABLE "Question" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");
ALTER TABLE "Answer" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");
ALTER TABLE "Answer" ADD FOREIGN KEY ("question_id") REFERENCES "Question" ("id");
