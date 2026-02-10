import ChatList from "@/components/testing/chatlist";
import ChatRoom from "@/components/testing/chatroom";
import { useState } from "react";
import { useUserStore } from "@/store/userStore";

function TestingPage() {
  const [roomId, setRoomId] = useState<number | null>(null);
  const userId = useUserStore((state) => state.user?.id) ?? 0;

  return (
    <div>
      <ChatList userId={userId} setRoomId={setRoomId} />
      {roomId && <ChatRoom userId={userId} roomId={roomId} />}
    </div>
  );
}

export default TestingPage;
