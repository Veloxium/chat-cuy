import { useQuery } from "@tanstack/react-query";
import axiosInstance from "./axiosInstance";
import { useUserStore } from "@/store/userStore";

export interface ChatRoomWithLastMessage {
  id: number;
  name: string;
  is_group: boolean;
  unreadCount: number;
  lastMessage: string;
  lastMessageTime: string;
  lastSenderId: number;
  otherUserId: number;
  otherUsername: string;
  otherAvatar: string;
  otherBio: string;
  created_at: string;
  updated_at: string;
}

export const useChatRooms = () => {
  const user = useUserStore((state) => state.user);
  return useQuery<ChatRoomWithLastMessage[]>({
    queryKey: ["getchatrooms", user?.id],
    queryFn: async () => {
      const response = await axiosInstance.get(`/api/chatrooms/${user?.id}`);
      return response.data;
    },
    enabled: !!user?.id,
  });
};
