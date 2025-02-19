import SideBarLayout from "@/layout/SideBarLayout";
import { IoLogOutOutline, IoSettings } from "react-icons/io5";

function SettingPage() {
  return (
    <div>
      <div className="flex items-center gap-2">
        <p className="text-4xl font-forta">Settings</p>
        <IoSettings size={36} />
      </div>
      <div className="flex flex-col gap-4 mt-4">
        <button
          type="button"
          className="w-full rounded-md px-4 py-2 ease-in duration-300 border-2 border-red-400 bg-red-600 hover:bg-red-500 text-white max-w-sm gap-2 items-center justify-center cursor-pointer"
        >
          <div className="flex justify-center items-center gap-2 font-medium">
            <p>Logout</p>
            <IoLogOutOutline size={26} />
          </div>
        </button>
      </div>
    </div>
  );
}

export default SettingPage;
