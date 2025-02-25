import { motion } from "framer-motion";
import { child } from "@/components/animation/oneonone";

interface ChatBarProps {
  id: number;
  msg: { msg: string; timestamp: string; isSender: boolean };
}

const formattedTime = (timestamp: string) =>
  new Date(timestamp).toLocaleTimeString([], {
    hour: "2-digit",
    minute: "2-digit",
    hour12: false,
  });

function ChatBar(items: ChatBarProps) {
  return (
    <motion.div
      variants={child}
      key={items.id}
      className={`flex ${items.msg.isSender ? "justify-end" : "justify-start"}`}
    >
      <div
        className={`relative overflow-hidden max-w-[80%] ${
          items.msg.isSender ? "pr-4" : "pl-4"
        }`}
      >
        <div
          className={`absolute -top-2 w-0 h-0 border-l-[10px] border-l-transparent border-r-[10px] border-r-transparent border-b-[20px] border-b-zprimary ${
            items.msg.isSender
              ? "right-1 rotate-[65deg]"
              : "left-1 -rotate-[65deg]"
          }`}
        />
        <div className="bg-zprimary text-white px-2 py-2 rounded-xl flex items-end gap-2">
          <p>{items.msg.msg}</p>
          <p className="text-[10px] text-slate-300">
            {formattedTime(items.msg.timestamp)}
          </p>
        </div>
      </div>
    </motion.div>
  );
}

export default ChatBar;