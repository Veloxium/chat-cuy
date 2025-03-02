CREATE TABLE "contacts" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "username" varchar NOT NULL,
    "avatar" varchar NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL,
    CONSTRAINT "fk_contacts_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);
