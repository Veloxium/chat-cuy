import { useRoomStore } from "@/store/roomStore";
import { useUserStore } from "@/store/userStore";
import { ChatRoomWithLastMessage } from "@/utils/queryChat";
import { getHourMinute } from "@/utils/timeFormat";

const API_URL = import.meta.env.VITE_API_URL;

function getAvatarUrl(avatar?: string) {
  if (!avatar || avatar === "") {
    return "https://api.dicebear.com/9.x/adventurer/svg?seed=Mason";
  }
  if (avatar.startsWith("http")) return avatar;
  return `${API_URL}${avatar}`;
}

type ChatItemsProps = {
  items: ChatRoomWithLastMessage;
};

function ChatItem({ items }: ChatItemsProps) {
  const addRoom = useRoomStore((state) => state.addRoom);
  const userId = useUserStore((state) => state.user?.id);

  const addRoomHandler = () => {
    addRoom({
      roomId: items.id,
      otherUserId: items.otherUserId,
      avatar: items.otherAvatar,
      username: items.otherUsername,
      bio: items.otherBio,
    });
  };

  return (
    <div className="p-2 rounded-md flex items-center gap-2 hover:bg-white">
      <div className="place-content-center w-16 h-16 rounded-full border relative overflow-hidden bg-white">
        <img src={getAvatarUrl(items.otherAvatar)} alt="avatar" className="w-full h-full object-cover" />
      </div>
      <button onClick={addRoomHandler} className="flex-1 cursor-pointer">
        <div className="flex-1 flex flex-col">
          <div className="flex justify-between">
            <p className="text-lg font-semibold">{items.otherUsername}</p>
            {items.lastMessageTime && (
              <p className="text-sm">{getHourMinute(items.lastMessageTime)}</p>
            )}
          </div>
          <div className="flex justify-start">
            <div className="flex justify-start">
              <div className="w-[140px]">
                <p className="line-clamp-1 text-left break-all">
                  {items.lastMessage
                    ? items.lastSenderId === userId
                      ? "Anda: " + items.lastMessage
                      : items.lastMessage
                    : ""}
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
