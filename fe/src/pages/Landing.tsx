import ZBackground from "@/components/custom/zbackground";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import IndexLayout from "@/layout/IndexLayout";
import { FaGoogle } from "react-icons/fa";
import { CiCircleQuestion } from "react-icons/ci";
import { toast } from "sonner";
import { Link } from "react-router-dom";
import { motion } from "motion/react";

function LandingPage() {
  const container = {
    hidden: { opacity: 0 },
    show: {
      opacity: 1,
      transition: {
        delayChildren: 0.5,
        staggerChildren: 0.2,
      },
    },
  };

  const item = {
    hidden: { opacity: 0 },
    show: { opacity: 1 },
  };

  return (
    <IndexLayout>
      <ZBackground />
      <div className="min-h-screen flex items-center justify-center">
        <motion.div
          initial={{ scale: 0.5, opacity: 0 }}
          animate={{ scale: 1, opacity: 1 }}
          transition={{ duration: 0.5 }}
          className="bg-white text-zprimary border border-zprimary flex flex-col items-center justify-center py-8 px-4 w-full max-w-lg rounded-lg"
        >
          <div className="text-center mb-6">
            <p className="text-2xl font-forta">Welcome to</p>
            <p className="text-5xl font-forta flex">Chat Cuy <span className="hidden md:block">😲</span></p>
            <p className="font-bold mt-2 text-sm md:text-lg">
              Chat dengan siapa pun yang kamu mau
            </p>
          </div>
          <motion.form
            variants={container}
            initial="hidden"
            animate="show"
            className="text-black flex flex-col gap-4 items-left w-full max-w-sm text-left"
          >
            <motion.p variants={item} className="-mb-2">
              Email
            </motion.p>
            <motion.div variants={item}>
              <Input placeholder="Masukan Email" />
            </motion.div>

            <motion.p variants={item} className="-mb-2">
              Password
            </motion.p>
            <motion.div variants={item}>
              <Input placeholder="Masukan Password" />
            </motion.div>
            <motion.div variants={item}>
              <Link to="/home">
                <button className="flex w-full max-w-sm gap-2 items-center justify-center cursor-pointer">
                  <p className="text-lg border-2 border-zprimary bg-zprimary hover:bg-hprimary text-white font-medium rounded-md px-4 py-2 w-full ease-in duration-300">
                    Login Cuy
                  </p>
                </button>
              </Link>
            </motion.div>
            <motion.div variants={item}>
              <Separator className="w-full max-w-sm" />
            </motion.div>
            <motion.div
              variants={item}
              className="flex w-full gap-4 items-center justify-center"
            >
              <button
                className="flex w-full max-w-sm gap-2 group items-center justify-center cursor-pointer"
                // onClick={() =>
                //   toast("Event has been created", {
                //     description: "Sunday, December 03, 2023 at 9:00 AM",
                //     action: {
                //       label: "Undo",
                //       onClick: () => console.log("Undo"),
                //     },
                //   })
                // }
              >
                <p className="bg-white border-2 border-zprimary group-hover:bg-hprimary text-zprimary group-hover:text-white font-medium rounded-md px-4 py-2 w-full ease-in duration-300">
                  With Google
                </p>
                <div className="bg-white border-2 border-zprimary group-hover:bg-hprimary rounded-md px-2 py-2 ease-in duration-300">
                  <FaGoogle
                    size={24}
                    className="text-zprimary group-hover:text-white ease-in duration-300"
                  />
                </div>
              </button>
            </motion.div>
          </motion.form>
          <a
            target="_blank"
            href={"https://github.com/Veloxium/chat-cuy"}
            className="mt-4 flex items-center text-xs gap-1 hover:underline justify-end w-full max-w-sm"
          >
            <p>Our Repository</p>
            <CiCircleQuestion size={16} />
          </a>
        </motion.div>
      </div>
    </IndexLayout>
  );
}

export default LandingPage;
