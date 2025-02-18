import ChatItem from "@/components/chatitem/chatitem";
import { Input } from "@/components/ui/input";
import dataChatItem from "@/data/chatitem";
import SideBarLayout from "@/layout/SideBarLayout";
import { useUserStore } from "@/store/userStore";
import axiosInstance from "@/utils/axiosInstance";
import { useChat } from "@/utils/queryChat";
import { useQuery } from "@tanstack/react-query";
import { useState } from "react";
import { CiSearch } from "react-icons/ci";
import { useDebounce } from "use-debounce";

function ChatPage() {
  const [searchTerm, setSearchTerm] = useState("");
  const [search] = useDebounce(searchTerm, 500);
  const { user } = useUserStore((state) => state);
  // const { data, status } = useQuery({
  //   queryKey: ["getchatrooms"],
  //   queryFn: async () => {
  //     const response = await axiosInstance.get("/chatrooms/" + user?.id);
  //     return response.data;
  //   },
  // });

  const chats = dataChatItem.filter(
    (item) =>
      item.name.toLowerCase().includes(search.toLowerCase()) ||
      item.message.text.toLowerCase().includes(search.toLowerCase())
  );

  const handleSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(e.target.value);
  };


  return (
    <SideBarLayout>
      <p className="text-4xl font-forta">Chats</p>
      <div className="relative mt-4">
        <Input
          type="search"
          placeholder="Search"
          className="pl-10"
          onChange={handleSearchChange}
        />
        <CiSearch size={24} className="absolute top-[6px] left-2 " />
      </div>
      <div className="flex flex-col space-y-2 mt-4">
        {chats.map((item: any) => (
          <ChatItem key={item.id} items={item} />
        ))}
      </div>
    </SideBarLayout>
  );
}

export default ChatPage;
