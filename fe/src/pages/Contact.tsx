import dataChatItem from "@/data/chatitem";
import SideBarLayout from "@/layout/SideBarLayout";
import { FaUserPlus } from "react-icons/fa6";
import { Link } from "react-router-dom";
import { motion } from "framer-motion";
import { container, child } from "@/components/animation/oneonone";
import ZBackground from "@/components/custom/zbackground";

function ContactPage() {
  return (
    <div>
      <div className="relative">
        <p className="text-4xl font-forta">List Contact</p>
        <motion.div
          variants={container}
          initial="hidden"
          animate="show"
          exit="exit"
          className="flex flex-col gap-2 mt-4"
        >
          {dataChatItem.map((item, index) => (
            <motion.div
              variants={child}
              key={index}
              className="p-2 rounded-md flex items-center gap-2 hover:bg-white"
            >
              <div className="w-16 h-16 rounded-full border relative overflow-hidden bg-white">
                <img src={item.image} alt="avatar" />
              </div>
              <Link to={"#"} className="flex-1">
                <div className="flex-1 flex flex-col">
                  <div className="flex justify-between">
                    <p className="text-lg font-semibold">{item.name}</p>
                  </div>
                </div>
              </Link>
            </motion.div>
          ))}
        </motion.div>
        <div className="fixed bottom-8 right-20">
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
