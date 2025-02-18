import { useQuery } from "@tanstack/react-query";
import axiosInstance from "./axiosInstance";
import { useUserStore } from "@/store/userStore";


export const useChat = async () => {
    const { user } = useUserStore((state) => state);
    return useQuery({
        queryKey: ["getchatrooms"],
        queryFn: async () => {
            const response = await axiosInstance.get("/api/chatrooms/" + user?.id);
            return response.data;
        },
    });
}