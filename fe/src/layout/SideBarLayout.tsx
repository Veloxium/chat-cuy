import Empty from "@/components/room/empty";
import Room from "@/components/room/room";
import Story from "@/components/story/story";
import IndexLayout from "@/layout/IndexLayout";
import { useRoomStore } from "@/store/roomStore";
import { useStoryStore } from "@/store/storyStore";
import { ReactNode, useEffect, useState } from "react";
import {
  IoChatbubbleEllipsesOutline,
  IoPeopleOutline,
  IoSettingsOutline,
} from "react-icons/io5";
import { LuCircleDotDashed } from "react-icons/lu";
import { Link, useLocation } from "react-router-dom";

function SideBarLayout({ children }: { children: ReactNode }) {
  const location = useLocation();
  const urlpath = location.pathname;
  const story = useStoryStore((state) => state.story);
  const roomID = useRoomStore((state) => state.room?.id);
  const room = useRoomStore((state) => state.room);
  const clearRoom = useRoomStore((state) => state.clearRoom);
  const clearStory = useStoryStore((state) => state.clearStory);

  const handlePageChange = () => {
    clearRoom();
    clearStory();
  };

  
  const options = [
    {
      name: "Chat",
      icon: (
        <IoChatbubbleEllipsesOutline className="text-slate-200" size={26} />
      ),
      path: "/chat",
    },
    {
      name: "Story",
      icon: <LuCircleDotDashed className="text-slate-200" size={26} />,
      path: "/story",
    },
    {
      name: "Contact",
      icon: <IoPeopleOutline className="text-slate-200" size={26} />,
      path: "/contact",
    },
    {
      name: "Setting",
      icon: <IoSettingsOutline className="text-slate-200" size={26} />,
      path: "/setting",
    },
  ];

  return (
    <div className="flex flex-row-reverse md:flex-row">
      <div className="bg-hprimary z-10 w-14 flex flex-col h-screen justify-between items-center">
        <p className="font-forta text-white text-center pt-4">Chat Cuy</p>
        <div className="flex-1 mt-20 mb-6">
          {options.map((item, index) => (
            <Link to={item.path} key={index} onClick={handlePageChange}>
              <div
                className={`border border-slate-300 my-6 w-10 h-10 place-items-center place-content-center rounded-md ease-in-out duration-300 ${
                  urlpath === item.path ? "bg-zprimary" : ""
                }`}
              >
                {item.icon}
              </div>
            </Link>
          ))}
        </div>
        <div className="mb-20">
          <div className="w-10 h-10 rounded-md overflow-hidden bg-white">
            <img
              src="https://api.dicebear.com/9.x/adventurer/svg?seed=Mason"
              alt="avatar"
            />
          </div>
        </div>
      </div>
      <div className="md:w-96 h-screen overflow-hidden flex-1 py-4 bg-zbase-100 overflow-y-scroll no-scrollbar">
        <IndexLayout>{children}</IndexLayout>
      </div>
      <div
        className={`md:relative left-0 w-[calc(100%-3rem)] md:flex overflow-hidden md:w-[calc(100%-24rem)] h-screen bg-zbase-200 ${
          roomID || story ? "flex" : "hidden"
        }`}
      >
        {roomID ? (
          <Room item={room!} clearRoom={clearRoom} />
        ) : story ? (
          <Story />
        ) : (
          <Empty />
        )}
      </div>
    </div>
  );
}

export default SideBarLayout;
