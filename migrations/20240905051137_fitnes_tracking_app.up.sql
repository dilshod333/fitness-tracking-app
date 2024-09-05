CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY,
    "username" VARCHAR,
    "email" VARCHAR,
    "password_hash" VARCHAR,
    "profile" JSONB
);
