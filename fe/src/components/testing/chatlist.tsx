import useSocket, { SOCKET_SERVER_URL } from "@/hooks/useSocket";
import axios from "axios";
import React, { useEffect } from "react";


interface ChatListProps {
  userId: string;
  setRoomId: (roomId: string) => void;
}

interface Message {
  id: string;
  text: string;
  createdAt: string;
  updatedAt: string;
  senderId: string;
  chatRoomId: string;
}

interface ChatRoom {
  id: string;
  name: string;
  unreadCount: number;
  messages: Message[];
  createdAt: string;
  updatedAt: string;
}


const ChatList: React.FC<ChatListProps> = ({ userId, setRoomId }) => {
  const [chatRooms, setChatRooms] = React.useState<ChatRoom[]>([]);
  const { enterChatRoom } = useSocket(userId);

   const getChatRooms = async (userId: string) => {
     try {
       const { data } = await axios.get<ChatRoom[]>(
         `${SOCKET_SERVER_URL}/chatrooms/${userId}`
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
      {chatRooms.map((room,index) => (
        <div
          key={room.id}
          onClick={() => {
            enterChatRoom(room.id);
            setRoomId(room.id);
          }}
          className="chat-room"
        >
          <p>
            <strong>{room.name}</strong> -{" "}
            {room.messages ? room.messages[index].text : "No messages yet"}
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
