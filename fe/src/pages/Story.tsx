import { child, container } from "@/components/animation/oneonone";
import ZBackground from "@/components/custom/zbackground";
import { StoryItem } from "@/data/storyitem";
import { useRoomStore } from "@/store/roomStore";
import { useStoryStore } from "@/store/storyStore";
import { motion } from "framer-motion";
import { FaPlus } from "react-icons/fa6";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import axiosInstance from "@/utils/axiosInstance";
import { LoaderCircle } from "lucide-react";
import { useRef } from "react";

const API_URL = import.meta.env.VITE_API_URL;

function getImageUrl(url: string) {
  if (!url || url === "") return "";
  if (url.startsWith("http")) return url;
  return `${API_URL}${url}`;
}

function StoryPage() {
  const clearRoom = useRoomStore((state) => state.clearRoom);
  const addStory = useStoryStore((state) => state.addStory);
  const fileInputRef = useRef<HTMLInputElement>(null);
  const queryClient = useQueryClient();

  const { data: stories, isLoading } = useQuery<StoryItem[]>({
    queryKey: ["getstories"],
    queryFn: async () => {
      const response = await axiosInstance.get("/api/stories");
      return response.data;
    },
  });

  const uploadMutation = useMutation({
    mutationFn: async (file: File) => {
      const formData = new FormData();
      formData.append("image", file);
      const res = await axiosInstance.post("/api/stories", formData, {
        headers: { "Content-Type": "multipart/form-data" },
      });
      return res.data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["getstories"] });
    },
  });

  const handleStory = (item: StoryItem) => {
    addStory(item);
    clearRoom();
  };

  const handleAddStory = () => {
    fileInputRef.current?.click();
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    const allowedTypes = ["image/jpeg", "image/png", "image/gif", "image/webp"];
    if (!allowedTypes.includes(file.type)) return;
    if (file.size > 5 * 1024 * 1024) return;

    uploadMutation.mutate(file);

    if (fileInputRef.current) {
      fileInputRef.current.value = "";
    }
  };

  return (
    <div>
      <p className="text-4xl font-forta">Stories</p>
      {isLoading && (
        <div className="flex justify-center mt-8">
          <LoaderCircle className="animate-spin text-zprimary" size={32} />
        </div>
      )}
      <motion.div
        variants={container()}
        initial="hidden"
        animate="show"
        exit="exit"
        className="grid grid-cols-2 gap-2 mt-4"
      >
        <motion.button
          variants={child}
          onClick={handleAddStory}
          disabled={uploadMutation.isPending}
          className="relative overflow-hidden z-10 w-full aspect-square place-content-center place-items-center text-center rounded-md border-2 border-dashed border-zprimary cursor-pointer"
        >
          <div className="absolute -z-10 w-full h-full top-0">
            <ZBackground />
          </div>
          {uploadMutation.isPending ? (
            <LoaderCircle className="animate-spin text-white" size={26} />
          ) : (
            <FaPlus size={26} className="text-white" />
          )}
        </motion.button>
        <input
          ref={fileInputRef}
          type="file"
          accept="image/jpeg,image/png,image/gif,image/webp"
          onChange={handleFileChange}
          className="hidden"
        />
        {(stories || []).map((item, i) => (
          <motion.button
            onClick={() => handleStory(item)}
            variants={child}
            key={i}
            className="relative overflow-hidden w-full aspect-square flex justify-center items-center text-center rounded-md bg-black"
          >
            <img
              src={getImageUrl(item.image_url)}
              alt="story"
              loading="lazy"
              className="object-cover w-full h-full"
            />
            <div className="bg-black w-full h-full absolute bg-opacity-25" />
            <p className="text-white absolute top-2 left-2 font-semibold">
              {item.username}
            </p>
          </motion.button>
        ))}
      </motion.div>
    </div>
  );
}

export default StoryPage;
