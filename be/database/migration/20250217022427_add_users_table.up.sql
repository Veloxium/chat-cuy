CREATE TABLE "users" (
   "id" bigserial PRIMARY KEY,
   "username" varchar NOT NULL,
   "email" varchar NOT NULL,
   "password" varchar NOT NULL,
   "avatar" varchar NOT NULL,
   "bio" varchar NOT NULL,
   "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   "deleted_at" TIMESTAMP NULL
)
