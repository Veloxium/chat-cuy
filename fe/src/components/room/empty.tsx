import { useEffect, useState } from "react";
import ZBackground from "../custom/zbackground";
import { motion } from "framer-motion";

const textList = [
  "Sigma? +10.000 auraaa!",
  "Gedagedigeda ooo ~",
  "Apa yang dicari orang sigma ?",
  "Bintang Skibidi !!!",
  "Skibidi in the chat",
  "Anomali Sigma",
  "Ngoding bang",
  "Serabutan ngoding",
];

function Empty() {
  const [textIndex, setTextIndex] = useState<number>(0);
  const [displayText, setDisplayText] = useState<string>("");
  const [isDeleting, setIsDeleting] = useState<boolean>(false);

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
    <div className="relative w-full z-10 h-full flex items-center justify-center p-10 gap-4">
      <p className="font-forta text-white">CHATCUY</p>
      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 0.5 }}
        className="text-4xl font-thin text-white font-forta"
      >
        {displayText}
        <motion.span
          animate={{ opacity: [1, 0, 1] }}
          transition={{ duration: 0.8, repeat: Infinity }}
        >
          |
        </motion.span>
      </motion.div>
      <ZBackground />
    </div>
  );
}

export default Empty;
