import ZBackground from "@/components/custom/zbackground";
import LoginForm from "@/components/form/loginform";
import RegisterForm from "@/components/form/registerform";
import IndexLayout from "@/layout/IndexLayout";
import { AnimatePresence, motion } from "motion/react";
import { useEffect, useState } from "react";
import { CiCircleQuestion } from "react-icons/ci";

const textList = [
  "Gabut? Chat Cuy aja !",
  "Gedagedigeda ooo ~",
  "Apa yang dicari orang sigma ?",
  "Bintang Skibidi !!!",
];

function LandingPage() {
  const [textIndex, setTextIndex] = useState<number>(0);
  const [displayText, setDisplayText] = useState<string>("");
  const [isDeleting, setIsDeleting] = useState<boolean>(false);
  const [isNewUser, setIsNewUser] = useState<boolean>(false);

  useEffect(() => {
    const currentText = textList[textIndex];
    let typingSpeed = isDeleting ? 50 : 100;

    const typingInterval = setInterval(() => {
      setDisplayText((prev) =>
        isDeleting ? prev.slice(0, -1) : currentText.slice(0, prev.length + 1)
      );

      if (!isDeleting && displayText === currentText) {
        setTimeout(() => setIsDeleting(true), 1000);
      }

      if (isDeleting && displayText === "") {
        setIsDeleting(false);
        setTextIndex((prev) => (prev + 1) % textList.length);
      }
    }, typingSpeed);

    return () => clearInterval(typingInterval);
  }, [displayText, isDeleting, textIndex]);

  return (
    <IndexLayout>
      <ZBackground />
      <div className="min-h-screen h-full py-10 flex items-center justify-center">
        <motion.div
          initial={{ scale: 0.5, opacity: 0 }}
          animate={{ scale: 1, opacity: 1 }}
          transition={{ duration: 0.5 }}
          className={`bg-white text-zprimary border border-zprimary flex flex-col items-center justify-center py-8 px-4 w-full rounded-lg ease-in-out duration-300 ${
            isNewUser ? "max-w-4xl" : "max-w-lg"
          }`}
        >
          <div className="flex flex-col items-center text-center mb-6">
            <p className="text-2xl font-forta">Selamat Datang di</p>
            <p className="text-5xl font-forta flex">
              Chat Cuy <span className="hidden md:block">😲</span>
            </p>
            <motion.div
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              transition={{ duration: 0.5 }}
              className="text-xs font-thin text-zprimary font-forta"
            >
              {displayText}
              <motion.span
                animate={{ opacity: [1, 0, 1] }}
                transition={{ duration: 0.8, repeat: Infinity }}
              >
                |
              </motion.span>
            </motion.div>
          </div>
          <AnimatePresence mode="wait">
            {isNewUser ? (
              <RegisterForm setIsNewUser={setIsNewUser} key="regis" />
            ) : (
              <LoginForm setIsNewUser={setIsNewUser} key="login" />
            )}
          </AnimatePresence>
          <a
            target="_blank"
            href={"https://github.com/Veloxium/chat-cuy"}
            className={`mt-4 flex items-center text-xs gap-1 hover:underline justify-end w-full ease-in-out duration-300 ${
              isNewUser ? "max-w-3xl" : "max-w-lg"
            }`}
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
