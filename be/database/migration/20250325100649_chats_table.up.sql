CREATE TABLE chats(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    chat_name VARCHAR(100),
    is_group BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
