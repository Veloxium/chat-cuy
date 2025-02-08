import { getHourMinute } from "@/utils/timeFormat";

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
  return (
    <div className="py-4 px-2 rounded-md flex items-center gap-2 hover:bg-white">
      <div className="w-16 h-16 rounded-full border relative overflow-hidden bg-white">
        <img src={items.image} alt="avatar" />
      </div>
      <div className="flex-1 flex flex-col">
        <div className="flex justify-between">
          <p className="text-lg font-semibold">{items.name}</p>
          <p>{getHourMinute(items.message.timestamp)}</p>
        </div>
        <div className="flex justify-start">
          <div className="flex justify-start">
            <div className="w-[120px]">
              <p className="line-clamp-1">
                {items.message.sender_id === 1
                  ? "Anda: " + items.message.text
                  : items.message.text}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default ChatItem;
