import { child, container } from "@/components/animation/oneonone";
import ZBackground from "@/components/custom/zbackground";
import { motion } from "framer-motion";
import { FaPlus } from "react-icons/fa6";

function StoryPage() {
  return (
    <div>
      <p className="text-4xl font-forta">Stories</p>
      <motion.div
        variants={container}
        initial="hidden"
        animate="show"
        exit="exit"
        className="grid grid-cols-2 gap-2 mt-4"
      >
        <motion.div
          variants={child}
          className="relative overflow-hidden z-10 w-full aspect-square place-content-center place-items-center text-center rounded-md border-2 border-dashed border-zprimary"
        >
          <div className="absolute -z-10 w-full h-full top-0">
            <ZBackground />
          </div>
          <FaPlus size={26} className="text-white" />
        </motion.div>
        {Array.from({ length: 9 }).map((_, i) => (
          <motion.div
            variants={child}
            key={i}
            className="w-full aspect-square place-content-center place-items-center text-center rounded-md bg-yellow-500"
          >
            Story {i + 1}
          </motion.div>
        ))}
      </motion.div>
    </div>
  );
}

export default StoryPage;
