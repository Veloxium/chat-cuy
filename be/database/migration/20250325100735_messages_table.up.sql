CREATE TABLE messages(
    id SERIAL PRIMARY KEY,
    chat_id UUID NOT NULL,
    sender_id UUID NOT NULL,
    content TEXT NOT NULL,
    message_type VARCHAR(20) DEFAULT 'text',
    media_url VARCHAR(255),
    is_read BOOLEAN DEFAULT FALSE,
    is_reply BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_messages_chats FOREIGN KEY(chat_id) REFERENCES chats(id) ON DELETE CASCADE,
    CONSTRAINT fk_messages_users FOREIGN KEY(sender_id) REFERENCES users(id) ON DELETE CASCADE
);
