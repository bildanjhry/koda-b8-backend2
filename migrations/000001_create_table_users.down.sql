CREATE TABLE "users" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "email" VARCHAR(30),
    "password" VARCHAR(30),
    "created_at" TIMESTAMP DEFAULT NOW(),
    "updated_at" TIMESTAMP DEFAULT NOW()
);


SELECT * FROM "users";