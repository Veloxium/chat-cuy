import { useRoomStore } from "@/store/roomStore";
import { getHourMinute } from "@/utils/timeFormat";
import { Link } from "react-router-dom";

type ChatItemsProps = {
  items: {
    id: number;
    name: string;
    message: {
      text: string;
      timestamp: string;
      sender_id: number;
    };
    image: string;
  };
};

function ChatItem({ items }: ChatItemsProps) {
  const addRoom = useRoomStore((state) => state.addRoom);
  const addRoomHandler = () => {
    addRoom({
      id: items.id,
      avatar: items.image,
      username: items.name,
      bio: "Behhhhh",
    });
  }
  return (
    <div className="p-2 rounded-md flex items-center gap-2 hover:bg-white">
      <div className="place-content-center w-16 h-16 rounded-full border relative overflow-hidden bg-white">  
        <img src={items.image} alt="avatar" />
      </div>
      <button onClick={addRoomHandler} className="flex-1 cursor-pointer">
        <div className="flex-1 flex flex-col">
          <div className="flex justify-between">
            <p className="text-lg font-semibold">{items.name}</p>
            <p className="text-sm">{getHourMinute(items.message.timestamp)}</p>
          </div>
          <div className="flex justify-start">
            <div className="flex justify-start">
              <div className="w-[140px]">
                <p className="line-clamp-1 text-left break-all">
                  {items.message.sender_id === 1
                    ? "Anda: " + items.message.text
                    : items.message.text}
                </p>
              </div>
            </div>
          </div>
        </div>
      </button>
    </div>
  );
}

export default ChatItem;
