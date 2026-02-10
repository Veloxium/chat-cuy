import { useUserStore } from "@/store/userStore";
import { IoLogOutOutline, IoSettings } from "react-icons/io5";
import { useNavigate } from "react-router-dom";
import { useRef, useState } from "react";
import axiosInstance from "@/utils/axiosInstance";
import { LoaderCircle } from "lucide-react";
import { FaCamera } from "react-icons/fa6";

const API_URL = import.meta.env.VITE_API_URL;

function getAvatarUrl(avatar?: string) {
  if (!avatar || avatar === "") {
    return "https://api.dicebear.com/9.x/adventurer/svg?seed=Mason";
  }
  if (avatar.startsWith("http")) {
    return avatar;
  }
  return `${API_URL}${avatar}`;
}

function SettingPage() {
  const user = useUserStore((state) => state.user);
  const logout = useUserStore((state) => state.logout);
  const updateAvatar = useUserStore((state) => state.updateAvatar);
  const navigate = useNavigate();
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [isUploading, setIsUploading] = useState(false);
  const [uploadError, setUploadError] = useState("");

  const handleLogout = () => {
    logout();
    sessionStorage.clear();
    navigate("/");
  };

  const handleAvatarClick = () => {
    fileInputRef.current?.click();
  };

  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    // Validate file type
    const allowedTypes = ["image/jpeg", "image/png", "image/gif", "image/webp"];
    if (!allowedTypes.includes(file.type)) {
      setUploadError("Only JPG, PNG, GIF, and WebP files are allowed");
      return;
    }

    // Validate file size (max 2MB)
    if (file.size > 2 * 1024 * 1024) {
      setUploadError("File size must be less than 2MB");
      return;
    }

    setIsUploading(true);
    setUploadError("");

    try {
      const formData = new FormData();
      formData.append("avatar", file);

      const res = await axiosInstance.put("/api/profile/avatar", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });

      const updatedUser = res.data.data;
      updateAvatar(updatedUser.avatar);
    } catch (err: any) {
      const message =
        err.response?.data?.message || "Failed to upload avatar";
      setUploadError(message);
    } finally {
      setIsUploading(false);
      // Reset file input
      if (fileInputRef.current) {
        fileInputRef.current.value = "";
      }
    }
  };

  return (
    <div>
      <div className="flex items-center gap-2">
        <p className="text-4xl font-forta">Settings</p>
        <IoSettings size={36} />
      </div>
      <div className="flex flex-col gap-4 mt-4">
        {/* Profile Section */}
        <div className="flex flex-col items-center gap-3 p-4 bg-white rounded-md">
          <div className="relative group">
            <div className="w-24 h-24 rounded-full overflow-hidden bg-white border-2 border-zprimary">
              <img
                src={getAvatarUrl(user?.avatar)}
                alt="avatar"
                className="w-full h-full object-cover"
              />
            </div>
            <button
              type="button"
              onClick={handleAvatarClick}
              disabled={isUploading}
              className="absolute bottom-0 right-0 w-8 h-8 bg-zprimary rounded-full flex items-center justify-center text-white hover:bg-hprimary ease-in duration-300 cursor-pointer"
            >
              {isUploading ? (
                <LoaderCircle className="animate-spin" size={16} />
              ) : (
                <FaCamera size={14} />
              )}
            </button>
            <input
              ref={fileInputRef}
              type="file"
              accept="image/jpeg,image/png,image/gif,image/webp"
              onChange={handleFileChange}
              className="hidden"
            />
          </div>
          <div className="text-center">
            <p className="text-lg font-semibold">{user?.username}</p>
            <p className="text-sm text-gray-500">{user?.email}</p>
            {user?.bio && (
              <p className="text-sm text-gray-400 mt-1">{user.bio}</p>
            )}
          </div>
          {uploadError && (
            <p className="text-sm text-red-500">{uploadError}</p>
          )}
        </div>

        {/* Logout */}
        <button
          type="button"
          onClick={handleLogout}
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
