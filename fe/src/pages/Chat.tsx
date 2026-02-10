import ChatItem from "@/components/chatitem/chatitem";
import { Input } from "@/components/ui/input";
import { useChatRooms } from "@/utils/queryChat";
import { useState } from "react";
import { CiSearch } from "react-icons/ci";
import { useDebounce } from "use-debounce";
import { motion } from "framer-motion";
import { container, child } from "@/components/animation/oneonone";
import { LoaderCircle } from "lucide-react";

function ChatPage() {
  const [searchTerm, setSearchTerm] = useState("");
  const [search] = useDebounce(searchTerm, 500);
  const { data: chatRooms, isLoading, isError } = useChatRooms();

  const chats = (chatRooms || []).filter(
    (item) =>
      item.otherUsername.toLowerCase().includes(search.toLowerCase()) ||
      (item.lastMessage || "").toLowerCase().includes(search.toLowerCase())
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
      {isLoading && (
        <div className="flex justify-center mt-8">
          <LoaderCircle className="animate-spin text-zprimary" size={32} />
        </div>
      )}
      {isError && (
        <p className="text-red-500 text-center mt-4">Gagal memuat chat</p>
      )}
      <motion.div
        variants={container()}
        initial="hidden"
        animate="show"
        exit="exit"
        className="flex flex-col space-y-2 mt-4"
      >
        {chats.map((item) => (
          <motion.div variants={child} key={item.id}>
            <ChatItem items={item} />
          </motion.div>
        ))}
      </motion.div>
    </div>
  );
}

export default ChatPage;
