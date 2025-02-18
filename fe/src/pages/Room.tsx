import ChatBar from "@/components/chatbar/chatbar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import useSocket from "@/hooks/use-socket";
import IndexLayout from "@/layout/IndexLayout";
import { useUserStore } from "@/store/userStore";
import { useEffect, useState } from "react";
import { BsThreeDotsVertical } from "react-icons/bs";
import { IoArrowBackSharp, IoSendSharp } from "react-icons/io5";
import { useParams } from "react-router-dom";

function RoomPage() {
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
    <IndexLayout>
      <div className="flex flex-col justify-between h-screen overflow-hidden">
        <div className="p-2 rounded-md flex items-center gap-2 bg-zbase-100">
          <button
            onClick={() => window.location.replace("/chat")}
            className="flex items-center p-2"
          >
            <IoArrowBackSharp size={26} />
          </button>
          <div className="w-16 h-16 rounded-full border relative overflow-hidden bg-white">
            <img
              src={"https://api.dicebear.com/9.x/adventurer/svg?seed=Jude"}
              alt="avatar"
            />
          </div>
          <div className="flex-1">
            <div className="flex-1 flex justify-between">
              <div className="flex flex-1 flex-col">
                <p className="text-lg font-semibold capitalize">{username}</p>
                <p className="text-sm">Serabutan Ngoding</p>
              </div>
              <div className="flex items-center p-2">
                <BsThreeDotsVertical size={26} />
              </div>
            </div>
          </div>
        </div>
        <div
          id="chatroom"
          className="flex-1 flex flex-col h-full rounded-md gap-2 overflow-y-scroll mt-2"
        >
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
        </div>
        <div className="fixed w-[calc(100vw-32px)] bottom-0 bg-zbase-100">
          <form onSubmit={onSubmitHandler}>
            <div className="w-full py-4 px-1 flex items-center gap-2">
              <Input
                placeholder="Type a message"
                className="h-10"
                type="text"
                onChange={(e) => setMsg(e.target.value)}
                value={msg}
              />
              <Button type="submit" className="h-10 bg-zprimary">
                <IoSendSharp />
              </Button>
            </div>
          </form>
        </div>
      </div>
    </IndexLayout>
  );
}

export default RoomPage;
