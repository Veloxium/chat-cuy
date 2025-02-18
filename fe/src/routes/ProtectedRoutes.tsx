import ChatPage from "@/pages/Chat";
import Profile from "@/pages/Profile";
import RoomPage from "@/pages/Room";
import StoryPage from "@/pages/Story";
import { Navigate, Route, Routes } from "react-router-dom";

export const ProtectedRoutes = () => {
   return (
      <Routes>
         <Route path="/chat" element={<ChatPage />} />
         <Route path="/chat/:username" element={<RoomPage />} />
         <Route path="/story" element={<StoryPage />} />
         <Route path="*" element={<Navigate to="/chat" replace />} />
         <Route path="/profile" element={<Profile />} />
      </Routes>
   );
};
