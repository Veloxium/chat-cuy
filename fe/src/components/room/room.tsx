import { container } from "@/components/animation/oneonone";
import ChatBar from "@/components/chatbar/chatbar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import useSocket, { Message } from "@/hooks/useSocket";
import { RoomType } from "@/store/roomStore";
import { useUserStore } from "@/store/userStore";
import { motion } from "framer-motion";
import { useEffect, useRef, useState } from "react";
import { BsThreeDotsVertical } from "react-icons/bs";
import { IoIosArrowBack } from "react-icons/io";
import { IoSendSharp } from "react-icons/io5";
import { MdBlock } from "react-icons/md";
import { LoaderCircle } from "lucide-react";

const API_URL = import.meta.env.VITE_API_URL;

function getAvatarUrl(avatar?: string) {
  if (!avatar || avatar === "") {
    return "https://api.dicebear.com/9.x/adventurer/svg?seed=Mason";
  }
  if (avatar.startsWith("http")) return avatar;
  return `${API_URL}${avatar}`;
}

function Room({ item, clearRoom }: { item: RoomType; clearRoom: () => void }) {
  const [msg, setMsg] = useState("");
  const [isLoadingMessages, setIsLoadingMessages] = useState(false);
  const chatEndRef = useRef<HTMLDivElement>(null);

  const user = useUserStore((state) => state.user);
  const userId = user?.id ?? 0;

  const { messages, enterChatRoom, sendMessage, getMessages } = useSocket(userId);

  const roomMessages: Message[] = messages[item.roomId] || [];

  // Enter the chat room and load messages when room changes
  useEffect(() => {
    if (item.roomId) {
      enterChatRoom(item.roomId);
      setIsLoadingMessages(true);
      getMessages(item.roomId).finally(() => setIsLoadingMessages(false));
    }
  }, [item.roomId, enterChatRoom, getMessages]);

  // Auto-scroll to bottom when messages change
  useEffect(() => {
    chatEndRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [roomMessages.length]);

  const onSubmitHandler = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!msg.trim()) return;
    sendMessage(item.roomId, msg);
    setMsg("");
  };

  return (
    <div className="relative h-full w-full p-4">
      <div className="relative flex flex-col justify-between h-full w-full">
        <div className="p-2 rounded-md flex items-center gap-2 bg-zbase-100">
          <button
            onClick={() => {
              clearRoom();
            }}
            className="flex items-center p-2"
          >
            <IoIosArrowBack size={28} />
          </button>
          <div className="w-16 h-16 rounded-full border relative overflow-hidden bg-white">
            <img src={getAvatarUrl(item.avatar)} alt="avatar" className="w-16 h-16 object-cover" />
          </div>
          <div className="flex-1">
            <div className="flex-1 flex justify-between">
              <div className="flex flex-1 flex-col">
                <p className="text-lg font-semibold capitalize">
                  {item.username}
                </p>
                <p className="text-sm line-clamp-1">{item.bio}</p>
              </div>
              <div className="flex items-center p-2">
                <Popover>
                  <PopoverTrigger>
                    <BsThreeDotsVertical size={26} />
                  </PopoverTrigger>
                  <PopoverContent className="w-full p-0">
                    <button className="flex items-center gap-1 px-5 py-3 rounded-md text-red-500 hover:text-white hover:bg-red-500">
                      <p>Blokir</p>
                      <MdBlock />
                    </button>
                  </PopoverContent>
                </Popover>
              </div>
            </div>
          </div>
        </div>
        <motion.div
          id="chatroom"
          variants={container(true)}
          initial="hidden"
          animate="show"
          exit="exit"
          className="no-scrollbar w-full flex-1 flex flex-col rounded-md gap-4 overflow-y-scroll mt-2"
        >
          <div className="h-2" />
          {isLoadingMessages && (
            <div className="flex justify-center py-4">
              <LoaderCircle className="animate-spin text-zprimary" size={28} />
            </div>
          )}
          {roomMessages.map((message) => (
            <ChatBar
              key={message.id}
              id={message.id}
              msg={{
                msg: message.text,
                timestamp: message.createdAt,
                isSender: message.senderId === userId,
              }}
            />
          ))}
          <div ref={chatEndRef} />
        </motion.div>
        <form
          onSubmit={onSubmitHandler}
          className="fixed md:relative w-[calc(100%-5rem)] md:w-full bottom-4 flex items-center gap-2"
        >
          <Input
            placeholder="Type a message"
            className="bg-white h-12 pl-4"
            type="text"
            onChange={(e) => setMsg(e.target.value)}
            value={msg}
          />
          <Button type="submit" className="w-20 h-10 bg-zprimary">
            <IoSendSharp />
          </Button>
        </form>
      </div>
    </div>
  );
}

export default Room;
