CREATE TABLE users (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   username VARCHAR(50) NOT NULL,
   email VARCHAR(100) NOT NULL,
   password VARCHAR(255) NOT NULL,
   profile_picture VARCHAR(255) NOT NULL,
   about_message VARCHAR(255) NOT NULL,
   is_online BOOLEAN DEFAULT FALSE,
   last_seen TIMESTAMP,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT NULL,
   deleted_at TIMESTAMP NULL
);
