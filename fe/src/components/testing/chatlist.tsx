import useSocket from "@/hooks/useSocket";
import axiosInstance from "@/utils/axiosInstance";
import React, { useEffect } from "react";

interface ChatListProps {
  userId: number;
  setRoomId: (roomId: number) => void;
}

interface ChatRoom {
  id: number;
  name: string;
  unreadCount: number;
  lastMessage: string;
  otherUsername: string;
}

const ChatList: React.FC<ChatListProps> = ({ userId, setRoomId }) => {
  const [chatRooms, setChatRooms] = React.useState<ChatRoom[]>([]);
  const { enterChatRoom } = useSocket(userId);

  const getChatRooms = async (uid: number) => {
    try {
      const { data } = await axiosInstance.get<ChatRoom[]>(
        `/api/chatrooms/${uid}`
      );
      setChatRooms(data);
    } catch (error) {
      console.error("Failed to fetch chat rooms", error);
    }
  };

  useEffect(() => {
    getChatRooms(userId);
  }, [userId]);

  return (
    <div className="">
      <h2>Chat Rooms</h2>
      {chatRooms.map((room) => (
        <div
          key={room.id}
          onClick={() => {
            enterChatRoom(room.id);
            setRoomId(room.id);
          }}
          className="chat-room"
        >
          <p>
            <strong>{room.otherUsername || room.name}</strong> -{" "}
            {room.lastMessage || "No messages yet"}
          </p>
          {room.unreadCount > 0 && (
            <span className="unread">{room.unreadCount}</span>
          )}
        </div>
      ))}
    </div>
  );
};

export default ChatList;
