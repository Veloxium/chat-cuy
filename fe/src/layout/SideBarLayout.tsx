import ZBackground from "@/components/custom/zbackground";
import IndexLayout from "@/layout/IndexLayout";
import React from "react";
import {
  IoChatbubbleEllipsesOutline,
  IoPeopleOutline,
  IoSettingsOutline,
} from "react-icons/io5";
import { LuCircleDotDashed } from "react-icons/lu";

function SideBarLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex">
      <div className="w-14 bg-zbase-200 flex flex-col h-screen justify-between items-center">
        <div className="h-14 w-14 relative flex items-center overflow-hidden z-10">
          <ZBackground />
          <p className="font-forta text-white text-center">Chat Cuy</p>
        </div>
        <div className="space-y-4 flex-1 mt-20">
          <div className="border border-zprimary bg-zprimary w-10 h-10 place-items-center place-content-center rounded-md">
            <IoChatbubbleEllipsesOutline color="white" size={26} />
          </div>
          <div className="border border-slate-400 w-10 h-10 place-items-center place-content-center rounded-md">
            <LuCircleDotDashed className="text-slate-500" size={26} />
          </div>
          <div className="border border-slate-400 w-10 h-10 place-items-center place-content-center rounded-md">
            <IoPeopleOutline className="text-slate-500" size={26} />
          </div>
          <div className="border border-slate-400 w-10 h-10 place-items-center place-content-center rounded-md">
            <IoSettingsOutline className="text-slate-500" size={26} />
          </div>
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
      <div className="h-screen flex-1 overflow-auto py-4 bg-zbase-100">
        <IndexLayout>{children}</IndexLayout>
      </div>
    </div>
  );
}

export default SideBarLayout;
