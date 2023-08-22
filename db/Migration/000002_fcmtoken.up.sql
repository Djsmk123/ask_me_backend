CREATE TABLE "Token" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INTEGER NOT NULL,
  "jwt_token" VARCHAR NOT NULL,
  "is_valid" INTEGER DEFAULT 1 NOT NULL,
  "created_at" timestamptz DEFAULT now() NOT NULL
);

ALTER TABLE "Token" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");


CREATE TABLE "FcmToken" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INTEGER NOT NULL,
  "fcm_token" VARCHAR NOT NULL,
  "is_valid" INTEGER DEFAULT 1 NOT NULL,
  "created_at" timestamptz DEFAULT now() NOT NULL
);

ALTER TABLE "FcmToken" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");
