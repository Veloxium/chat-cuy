import ChatBar from "@/components/chatbar/chatbar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { RoomType } from "@/store/roomStore";
import { useUserStore } from "@/store/userStore";
import { useState } from "react";
import { BsThreeDotsVertical } from "react-icons/bs";
import { IoCloseCircle, IoSendSharp } from "react-icons/io5";
import { useParams } from "react-router-dom";
import { motion } from "framer-motion";
import { container } from "@/components/animation/oneonone";

function Room({ item, clearRoom }: { item: RoomType; clearRoom: () => void }) {
  const { username } = useParams<{ username: string }>();
  // const socket = useSocket();
  const [msg, setMsg] = useState("");
  const [messages, setMessages] = useState<string[]>([]);
  const [message, setMessage] = useState("");
  const [roomID, setRoomID] = useState("");

  const user1 = useUserStore((state) => state.user?.username);

  // useEffect(() => {
  //   if (!socket || !user1 || !username) {
  //     console.log(socket);
  //     console.log(user1);
  //     console.log(username);
  //     console.log("Socket or user1 or username is not defined");
  //     return;
  //   };

  //   const generatedRoomID = `${user1}_${username}`;
  //   setRoomID(generatedRoomID);
  //   socket.emit("join-room", { user1, user2: username });

  //   socket.on("receive-message", (msg: string) => {
  //     setMessages((prev) => [...prev, msg]);
  //   });

  //   return () => {
  //     socket.off("receive-message");
  //   };
  // }, [socket, user1, username]);

  const scrollBottom = () => {
    const chatroom = document.getElementById("chatroom");
    chatroom?.scrollTo(0, chatroom.scrollHeight);
  };

  const onSubmitHandler = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!msg) {
      console.log("Message is empty");
    }
    // sendMessage();

    scrollBottom();
  };

  // const sendMessage = () => {
  //   if (!socket || !roomID || !msg) return;
  //   socket.emit("send-message", { roomID, message });
  //   setMessage("");
  // };

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
            <IoCloseCircle size={28} className="text-red-500" />
          </button>
          <div className="w-16 h-16 rounded-full border relative overflow-hidden bg-white">
            <img src={item.avatar} alt="avatar" className="w-16 h-16" />
          </div>
          <div className="flex-1">
            <div className="flex-1 flex justify-between">
              <div className="flex flex-1 flex-col">
                <p className="text-lg font-semibold capitalize">
                  {item.username}
                </p>
                {/* Bio User */}
                <p className="text-sm line-clamp-1">Serabutan Ngoding</p>
              </div>
              <div className="flex items-center p-2">
                <BsThreeDotsVertical size={26} />
              </div>
            </div>
          </div>
        </div>
        <motion.div
          id="chatroom"
          variants={container}
          initial="hidden"
          animate="show"
          exit="exit"
          className="no-scrollbar w-full flex-1 flex flex-col rounded-md gap-2 overflow-y-scroll mt-2"
        >
          <div className="h-6" />
          <ChatBar
            id={1}
            msg={{
              msg: "Hello pukimak kelass, kelas tidak king anjayyy",
              timestamp: "2021-01-01T11:56:00Z",
              isSender: true,
            }}
          />
          <ChatBar
            id={2}
            msg={{
              msg: "Anjaayy bang",
              timestamp: "2021-01-01T11:57:00Z",
              isSender: false,
            }}
          />
          <ChatBar
            id={1}
            msg={{
              msg: "Itu dia, yang di kelas, yang di kelas",
              timestamp: "2021-01-01T12:00:00Z",
              isSender: true,
            }}
          />
        </motion.div>
        <div className="fixed md:relative w-[calc(100%-5rem)] md:w-full bottom-4 flex items-center gap-2">
          <Input
            placeholder="Type a message"
            className="bg-white h-10"
            type="text"
            onChange={(e) => setMsg(e.target.value)}
            value={msg}
          />
          <Button type="submit" className="w-20 h-10 bg-zprimary">
            <IoSendSharp />
          </Button>
        </div>
      </div>
    </div>
  );
}

export default Room;
