import useSocket from "@/hooks/use-socket";
import { useEffect, useState } from "react";
import { Message } from "react-hook-form";

interface ChatRoomProps {
  userId: string;
  roomId: string;
}

const ChatRoom: React.FC<ChatRoomProps> = ({ userId, roomId }) => {
  interface RoomMessages {
    [key: string]: { senderId: string; text: string }[];
  }

  const [roomMessages, setRoomMessages] = useState<RoomMessages>({});
  const { messages, sendMessage, getMessages } = useSocket(userId);
  const [newMessage, setNewMessage] = useState("");

  const handleSendMessage = () => {
    if (newMessage.trim()) {
      sendMessage(roomId, newMessage, userId);
      setNewMessage("");
    }
  };

  useEffect(() => {
    getMessages(roomId);
  }, [roomId]);

  useEffect(() => {
    return setRoomMessages(messages);
  }, [messages]);

  return (
    <div className="chat-room">
      <h2>Room: {roomId}</h2>
      <div className="messages">
        {roomMessages[roomId]?.map((msg, index) => (
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
