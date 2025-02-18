import axios from "axios";
import { useEffect, useState } from "react";
import { io, Socket } from "socket.io-client";

export const SOCKET_SERVER_URL = "http://localhost:3001";


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

const useSocket = (userId: string) => {
    const [socket, setSocket] = useState<Socket | null>(null);
    const [messages, setMessages] = useState<Record<string, Message[]>>({});
    const [activeRoom, setActiveRoom] = useState<string>("");

    useEffect(() => {
        const newSocket = io(SOCKET_SERVER_URL);
        setSocket(newSocket);

        newSocket.emit("joinRoom", userId);

        // newSocket.on("updateChatList", (updatedRooms: ChatRoom[]) => {
        //     setChatRooms(updatedRooms);
        // });

        newSocket.on("receiveMessage", (newMessage: Message) => {
            console.log(newMessage);
            setMessages((prev) => ({
                ...prev,
                [activeRoom]: [...(prev[activeRoom] || []), newMessage],
            }));
        });

        return () => {
            newSocket.disconnect();
        };
    }, [userId]);

    const enterChatRoom = (roomId: string) => {
        if (socket) {
            socket.emit("enterChatRoom", { roomId, userId });
            setActiveRoom(roomId);
        }
    };

    const sendMessage = (roomId: string, content: string, userId: string,) => {
        if (socket) {
            socket.emit("sendMessage", { roomId, content, userId });
        }
    };

    const getMessages = async (roomId: string) => {
        try {
            const { data } = await axios.get<Message[]>(
                `${SOCKET_SERVER_URL}/messages/${roomId}`
            );
            setMessages((prev) => ({ ...prev, [roomId]: data }));
        } catch (error) {
            console.error("Failed to fetch messages", error);
        }
    };





    return { messages, activeRoom, enterChatRoom, sendMessage, getMessages };
};

export default useSocket;
