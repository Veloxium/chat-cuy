import useSocket from "@/hooks/useSocket";
import { Message } from "@/hooks/useSocket";
import { useEffect, useState } from "react";

interface ChatRoomProps {
  userId: number;
  roomId: number;
}

const ChatRoom: React.FC<ChatRoomProps> = ({ userId, roomId }) => {
  const { messages, sendMessage, getMessages } = useSocket(userId);
  const [newMessage, setNewMessage] = useState("");

  const roomMessages: Message[] = messages[roomId] || [];

  const handleSendMessage = () => {
    if (newMessage.trim()) {
      sendMessage(roomId, newMessage);
      setNewMessage("");
    }
  };

  useEffect(() => {
    getMessages(roomId);
  }, [roomId]);

  return (
    <div className="chat-room">
      <h2>Room: {roomId}</h2>
      <div className="messages">
        {roomMessages.map((msg, index) => (
          <p key={index}>
            <strong>{msg.senderId}:</strong> {msg.text}
          </p>
        ))}
      </div>
      <input
        type="text"
        value={newMessage}
        onChange={(e) => setNewMessage(e.target.value)}
      />
      <button onClick={handleSendMessage}>Send</button>
    </div>
  );
};

export default ChatRoom;
