import ChatList from "@/components/testing/chatlist";
import ChatRoom from "@/components/testing/chatroom";
import axios from "axios";
import { useEffect, useState } from "react";

interface ChatRoom {
  roomId: string;
  lastMessageSender: string;
  lastMessage: string;
  unreadCount: number;
}

function TestingPage() {
  const [roomId, setRoomId] = useState<string | null>(null);
  const userId = "cm77c5do80000p1ncjzv3tg30";

  return (
    <div>
      <ChatList userId={userId} setRoomId={setRoomId} />
      {roomId && <ChatRoom userId={userId} roomId={roomId} />}
    </div>
  );
}

export default TestingPage;
