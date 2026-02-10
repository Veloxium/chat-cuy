DROP INDEX IF EXISTS idx_contacts_contact_id;
DROP INDEX IF EXISTS idx_contacts_user_id;
DROP INDEX IF EXISTS idx_chat_room_members_chat_room_id;
DROP INDEX IF EXISTS idx_chat_room_members_user_id;
DROP INDEX IF EXISTS idx_messages_sender_id;
DROP INDEX IF EXISTS idx_messages_chat_room_id;

DROP TABLE IF EXISTS contacts;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS chat_room_members;
DROP TABLE IF EXISTS chat_rooms;
DROP TABLE IF EXISTS users;
