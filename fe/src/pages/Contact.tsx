import axiosInstance from "@/utils/axiosInstance";
import { useRoomStore } from "@/store/roomStore";
import { FaUserPlus } from "react-icons/fa6";
import { Link } from "react-router-dom";
import { motion } from "framer-motion";
import { container, child } from "@/components/animation/oneonone";
import ZBackground from "@/components/custom/zbackground";
import { useQuery } from "@tanstack/react-query";
import { LoaderCircle } from "lucide-react";
import { useUserStore } from "@/store/userStore";

const API_URL = import.meta.env.VITE_API_URL;

function getAvatarUrl(avatar?: string) {
  if (!avatar || avatar === "") {
    return "https://api.dicebear.com/9.x/adventurer/svg?seed=Mason";
  }
  if (avatar.startsWith("http")) return avatar;
  return `${API_URL}${avatar}`;
}

interface ContactWithUser {
  id: number;
  user_id: number;
  contact_id: number;
  username: string;
  email: string;
  bio: string;
  avatar: string;
  created_at: string;
}

function ContactPage() {
  const userId = useUserStore((state) => state.user?.id);
  const addRoom = useRoomStore((state) => state.addRoom);

  const { data: contacts, isLoading, isError } = useQuery<ContactWithUser[]>({
    queryKey: ["getcontacts"],
    queryFn: async () => {
      const response = await axiosInstance.get("/api/contacts");
      return response.data;
    },
  });

  const handleContactClick = async (contact: ContactWithUser) => {
    try {
      // Create or get existing private chat room
      const res = await axiosInstance.post("/api/chatrooms", {
        user_id_1: userId,
        user_id_2: contact.contact_id,
      });
      const room = res.data.data;
      addRoom({
        roomId: room.id,
        otherUserId: contact.contact_id,
        avatar: contact.avatar,
        username: contact.username,
        bio: contact.bio,
      });
    } catch (err) {
      console.error("Failed to open chat room:", err);
    }
  };

  return (
    <div>
      <div className="relative md:max-w-96">
        <div className="flex">
          <p className="text-4xl font-forta">List Contact</p>
          <Link className="hidden md:block" to="/addcontact">
            <button className="relative z-10 overflow-hidden border-2 border-dashed border-zprimary py-2 px-4 flex justify-center items-center gap-1 rounded-lg text-white">
              <ZBackground />
              <FaUserPlus size={20} />
              <p className="font-semibold">Add</p>
            </button>
          </Link>
        </div>
        {isLoading && (
          <div className="flex justify-center mt-8">
            <LoaderCircle className="animate-spin text-zprimary" size={32} />
          </div>
        )}
        {isError && (
          <p className="text-red-500 text-center mt-4">Gagal memuat kontak</p>
        )}
        <motion.div
          variants={container()}
          initial="hidden"
          animate="show"
          exit="exit"
          className="flex flex-col gap-2 mt-4"
        >
          {(contacts || []).map((item) => (
            <motion.div
              variants={child}
              key={item.id}
              className="p-2 rounded-md flex items-center gap-2 hover:bg-white"
            >
              <div className="w-16 h-16 rounded-full border relative overflow-hidden bg-white">
                <img src={getAvatarUrl(item.avatar)} alt="avatar" className="w-full h-full object-cover" />
              </div>
              <button
                onClick={() => handleContactClick(item)}
                className="flex-1 cursor-pointer text-left"
              >
                <div className="flex-1 flex flex-col">
                  <p className="text-lg font-semibold">{item.username}</p>
                  {item.bio && (
                    <p className="text-sm text-gray-500 line-clamp-1">{item.bio}</p>
                  )}
                </div>
              </button>
            </motion.div>
          ))}
        </motion.div>
        <div className="md:hidden fixed bottom-8 right-20">
          <Link to="/addcontact">
            <button className="relative z-10 overflow-hidden border-2 border-dashed border-zprimary py-2 px-4 flex justify-center items-center gap-1 rounded-lg text-white">
              <ZBackground />
              <FaUserPlus size={20} />
              <p className="font-semibold">Add</p>
            </button>
          </Link>
        </div>
      </div>
    </div>
  );
}

export default ContactPage;
