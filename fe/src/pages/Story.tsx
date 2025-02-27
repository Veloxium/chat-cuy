import { child, container } from "@/components/animation/oneonone";
import ZBackground from "@/components/custom/zbackground";
import { useRoomStore } from "@/store/roomStore";
import { useStoryStore } from "@/store/storyStore";
import { motion } from "framer-motion";
import { FaPlus } from "react-icons/fa6";

function StoryPage() {
  const clearRoom = useRoomStore((state) => state.clearRoom);
  const addStory = useStoryStore((state) => state.addStory);

  const handleStory = () => {
    addStory({
      id: 1,
      username: "Sigma",
      url: "https://i.pinimg.com/736x/5b/8e/b5/5b8eb524dc01b7201d50df40b4034f7e.jpg",
    });
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
        {Array.from({ length: 9 }).map((_, i) => (
          <motion.button
            onClick={handleStory}
            variants={child}
            key={i}
            className="w-full aspect-square place-content-center place-items-center text-center rounded-md bg-yellow-500"
          >
            Story {i + 1}
          </motion.button>
        ))}
      </motion.div>
    </div>
  );
}

export default StoryPage;
