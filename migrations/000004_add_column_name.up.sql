ALTER TABLE "users" ADD COLUMN "name" VARCHAR(40);

SELECT "id", "name", "email", "password", "created_at", "updated_at", "picture" 
FROM "users" ORDER BY "id" LIMIT 5 OFFSET 0;