import { useUserStore } from "@/store/userStore";
import axiosInstance from "@/utils/axiosInstance";
import { useCallback, useEffect, useRef, useState } from "react";

export interface Message {
  id: number;
  text: string;
  senderId: number;
  chatRoomId: number;
  createdAt: string;
  updatedAt: string;
}

const WS_BASE_URL = import.meta.env.VITE_API_URL?.replace(/^http/, "ws") || "ws://localhost:3001";

const useSocket = (userId: number) => {
  const wsRef = useRef<WebSocket | null>(null);
  const [messages, setMessages] = useState<Record<number, Message[]>>({});
  const [activeRoom, setActiveRoom] = useState<number | null>(null);
  const activeRoomRef = useRef<number | null>(null);
  const reconnectTimeoutRef = useRef<ReturnType<typeof setTimeout>>();

  const getToken = () => useUserStore.getState().user?.token || "";

  const connect = useCallback(() => {
    const token = getToken();
    if (!token) return;

    const ws = new WebSocket(`${WS_BASE_URL}/ws?token=${token}`);

    ws.onopen = () => {
      // Send joinRoom to register this user on the hub
      ws.send(JSON.stringify({
        type: "joinRoom",
        payload: { userId },
      }));
    };

    ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        if (data.type === "receiveMessage") {
          const msg: Message = data.payload;
          setMessages((prev) => ({
            ...prev,
            [msg.chatRoomId]: [...(prev[msg.chatRoomId] || []), msg],
          }));
        }
      } catch (err) {
        console.error("Failed to parse WebSocket message:", err);
      }
    };

    ws.onclose = () => {
      // Auto-reconnect after 3 seconds
      reconnectTimeoutRef.current = setTimeout(() => {
        connect();
      }, 3000);
    };

    ws.onerror = (err) => {
      console.error("WebSocket error:", err);
      ws.close();
    };

    wsRef.current = ws;
  }, [userId]);

  useEffect(() => {
    connect();

    return () => {
      clearTimeout(reconnectTimeoutRef.current);
      wsRef.current?.close();
    };
  }, [connect]);

  const enterChatRoom = useCallback((roomId: number) => {
    setActiveRoom(roomId);
    activeRoomRef.current = roomId;
    const ws = wsRef.current;
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({
        type: "enterChatRoom",
        payload: { roomId, userId },
      }));
    }
  }, [userId]);

  const sendMessage = useCallback((roomId: number, content: string) => {
    const ws = wsRef.current;
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({
        type: "sendMessage",
        payload: { roomId, content, userId },
      }));
    }
  }, [userId]);

  const getMessages = useCallback(async (roomId: number) => {
    try {
      const { data } = await axiosInstance.get<Message[]>(`/api/messages/${roomId}`);
      setMessages((prev) => ({ ...prev, [roomId]: data }));
    } catch (error) {
      console.error("Failed to fetch messages:", error);
    }
  }, []);

  return { messages, activeRoom, enterChatRoom, sendMessage, getMessages };
};

export default useSocket;
