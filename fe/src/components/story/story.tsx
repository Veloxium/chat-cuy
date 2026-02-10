import { Input } from "@/components/ui/input";
import { useStoryStore } from "@/store/storyStore";
import { useEffect, useRef, useState } from "react";
import { IoPauseSharp, IoPlay, IoSendSharp } from "react-icons/io5";
import { Button } from "../ui/button";
import { IoIosArrowBack } from "react-icons/io";

const API_URL = import.meta.env.VITE_API_URL;

function getImageUrl(url?: string) {
  if (!url || url === "") return "";
  if (url.startsWith("http")) return url;
  return `${API_URL}${url}`;
}

function getAvatarUrl(avatar?: string) {
  if (!avatar || avatar === "") {
    return "https://api.dicebear.com/9.x/adventurer/svg?seed=Mason";
  }
  if (avatar.startsWith("http")) return avatar;
  return `${API_URL}${avatar}`;
}

function Story() {
  const time = 15;
  const story = useStoryStore((state) => state.story);
  const [storyTime, setStoryTime] = useState<number>(0);
  const clearStory = useStoryStore((state) => state.clearStory);
  const [isPause, setIsPause] = useState<boolean>(false);
  const intervalRef = useRef<NodeJS.Timeout | null>(null);

  useEffect(() => {
    if (!isPause) {
      intervalRef.current = setInterval(() => {
        setStoryTime((prevTime) => {
          if (prevTime + 1 === time) {
            clearStory();
            clearInterval(intervalRef.current!);
          }
          return prevTime + 1;
        });
      }, 1000);
    } else if (intervalRef.current) {
      clearInterval(intervalRef.current);
    }

    return () => {
      if (intervalRef.current) {
        clearInterval(intervalRef.current);
      }
    };
  }, [isPause]);

  useEffect(() => {
    if (story) {
      setStoryTime(0);
    }
  }, [story]);

  return (
    <div className="relative h-full w-full place-content-center place-items-center bg-black">
      <div className="absolute w-full top-0">
        <div className="flex items-center gap-2 py-4 bg-black bg-opacity-15 justify-between">
          <div className="flex items-center gap-2">
            <button
              onClick={() => {
                clearStory();
              }}
              className="flex items-center p-2"
            >
              <IoIosArrowBack size={20} className="text-white" />
            </button>
            <div className="w-10 h-10 rounded-full overflow-hidden bg-white">
              <img
                src={getAvatarUrl(story?.avatar)}
                alt="avatar"
                className="object-cover w-full h-full"
              />
            </div>
            <p className="font-bold text-white">{story?.username}</p>
          </div>
          <div
            onClick={() => setIsPause(!isPause)}
            className="w-8 h-8 mr-2 bg-white rounded-full place-content-center place-items-center"
          >
            {isPause ? <IoPlay size={20} /> : <IoPauseSharp size={20} />}
          </div>
        </div>
        <div
          style={{
            width: `${(storyTime / time) * 100}%`,
            transition: "width 1s linear",
          }}
          className={`h-1 bg-white`}
        />
      </div>
      <div className="fixed bg-black bg-opacity-15 w-[calc(100%-3rem)] md:absolute bottom-0 px-4 pb-4 pt-2">
        <p className="text-center text-white mb-2">{story?.caption}</p>
        <div className="flex items-center gap-2">
          <div className="w-full bg-white bg-opacity-25 rounded-lg group">
            <Input
              placeholder="Type a message"
              className="h-10 pl-4 group-focus-within:bg-white rounded-lg"
              type="text"
            />
          </div>
          <Button className="w-20 h-10 bg-zprimary rounded-md place-content-center place-items-center">
            <IoSendSharp size={26} />
          </Button>
        </div>
      </div>
      <div className="h-full w-full flex justify-center items-center">
        <img src={getImageUrl(story?.image_url)} className="object-contain" />
      </div>
    </div>
  );
}

export default Story;
