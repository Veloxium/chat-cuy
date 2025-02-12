import ChatBar from "@/components/chatbar/chatbar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import IndexLayout from "@/layout/IndexLayout";
import { IoArrowBackSharp, IoSendSharp } from "react-icons/io5";
import { BsThreeDotsVertical } from "react-icons/bs";
import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { io, Socket } from "socket.io-client";
import ZBackground from "@/components/custom/zbackground";

let socket: Socket = {} as Socket;

function RoomPage() {
  const { username } = useParams<{ username: string }>();
  const [msg, setMsg] = useState<string>("");

  const socketInitializer = async () => {
    await fetch("/api/socket");
    socket = io({
      reconnection: true,
      reconnectionAttempts: 5,
      reconnectionDelay: 1000,
    });

    socket.on("connect", () => {
      console.log("connected");
    });

    socket.on("disconnect", () => {
      console.log("disconnected");
    });

    socket.on(
      "message",
      (msg: { msg: string; timestamp: string; id: string }) => {
        // setChat((prevChat) => [
        //   ...prevChat,
        //   {
        //     msg: msg.msg,
        //     timestamp: msg.timestamp,
        //     isSender: false,
        //   },
        // ]);
      }
    );
  };

  const scrollBottom = () => {
    const chatroom = document.getElementById("chatroom");
    chatroom?.scrollTo(0, chatroom.scrollHeight);
    return () => {
      if (socket) {
        socket.disconnect();
      }
    };
  };

  const onSubmitHandler = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (socket && socket.connected) {
      try {
        socket.emit("message", {
          msg: msg,
          timestamp: new Date().toISOString(),
        });
      } catch (error) {
        console.error(error);
      } finally {
        setMsg("");
        scrollBottom();
      }
    } else {
      console.error("Socket is not connected");
    }
  };

  useEffect(() => {
    socketInitializer();
    scrollBottom();
  }, []);

  return (
    <IndexLayout>
      <div className="flex flex-col justify-between h-screen overflow-hidden">
        <div className="p-2 rounded-md flex items-center gap-2 bg-zbase-100">
          <button
            onClick={() => window.history.back()}
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
        <form onSubmit={onSubmitHandler}>
          <div className="py-6 px-1 flex items-center gap-2">
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
        <div className="relative h-10">
            <ZBackground />
        </div>
      </div>
    </IndexLayout>
  );
}

export default RoomPage;
