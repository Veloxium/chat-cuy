import ChatItem from "@/components/chatitem/chatitem";
import { Input } from "@/components/ui/input";
import dataChatItem from "@/data/chatitem";
import SideBarLayout from "@/layout/SideBarLayout";
import { useUserStore } from "@/store/userStore";
import { useState } from "react";
import { CiSearch } from "react-icons/ci";
import { useDebounce } from "use-debounce";
import { motion } from "framer-motion";
import { container, child } from "@/components/animation/oneonone";

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
    <div>
      <p className="text-4xl font-forta">Chats</p>
      <div className="relative mt-4">
        <Input
          type="search"
          placeholder="Search"
          className="pl-10 h-10"
          onChange={handleSearchChange}
        />
        <CiSearch size={24} className="absolute top-[8px] left-2 " />
      </div>
      <motion.div
        variants={container()}
        initial="hidden"
        animate="show"
        exit="exit"
        className="flex flex-col space-y-2 mt-4"
      >
        {chats.map((item: any) => (
          <motion.div variants={child} key={item.id}>
            <ChatItem items={item} />
          </motion.div>
        ))}
      </motion.div>
    </div>
  );
}

export default ChatPage;
