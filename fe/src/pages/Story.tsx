import { child, container } from "@/components/animation/oneonone";
import ZBackground from "@/components/custom/zbackground";
import { dataStoryItem, StoryItem } from "@/data/storyitem";
import { useRoomStore } from "@/store/roomStore";
import { useStoryStore } from "@/store/storyStore";
import { motion } from "framer-motion";
import { FaPlus } from "react-icons/fa6";

function StoryPage() {
  const clearRoom = useRoomStore((state) => state.clearRoom);
  const addStory = useStoryStore((state) => state.addStory);

  const handleStory = (item : StoryItem) => {
    addStory(item);
    clearRoom();
  };
  return (
    <div>
      <p className="text-4xl font-forta">Stories</p>
      <motion.div
        variants={container()}
        initial="hidden"
        animate="show"
        exit="exit"
        className="grid grid-cols-2 gap-2 mt-4"
      >
        <motion.button
          variants={child}
          className="relative overflow-hidden z-10 w-full aspect-square place-content-center place-items-center text-center rounded-md border-2 border-dashed border-zprimary"
        >
          <div className="absolute -z-10 w-full h-full top-0">
            <ZBackground />
          </div>
          <FaPlus size={26} className="text-white" />
        </motion.button>
        {dataStoryItem.map((item, i) => (
          <motion.button
            onClick={() => handleStory(item)}
            variants={child}
            key={i}
            className="relative overflow-hidden w-full aspect-square flex justify-center items-center text-center rounded-md bg-black"
          >
            <img
              src={item.url}
              alt="story"
              loading="lazy"
            />
            <div className="bg-black w-full h-full absolute bg-opacity-25" />
            <p className="text-white absolute top-2 left-2 font-semibold">
              Username
            </p>
          </motion.button>
        ))}
      </motion.div>
    </div>
  );
}

export default StoryPage;
