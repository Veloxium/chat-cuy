CREATE TABLE group_invitations(
    id SERIAL PRIMARY KEY,
    chat_id UUID NOT NULL,
    sender_id UUID NOT NULL,
    receiver_id UUID NOT NULL,
    invite_token VARCHAR(100) NOT NULL,
    is_accepted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP + INTERVAL '24 HOURS',
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_group_invitations_chats FOREIGN KEY(chat_id) REFERENCES chats(id) ON DELETE CASCADE,
    CONSTRAINT fk_group_invitations_users_sender FOREIGN KEY(sender_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_group_invitations_users_receiver FOREIGN KEY(receiver_id) REFERENCES users(id) ON DELETE CASCADE
);
