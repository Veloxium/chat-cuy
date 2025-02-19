import SideBarLayout from "@/layout/SideBarLayout";
import AddContactPage from "@/pages/AddContact";
import ChatPage from "@/pages/Chat";
import ContactPage from "@/pages/Contact";
import RoomPage from "@/pages/Room";
import SettingPage from "@/pages/Setting";
import StoryPage from "@/pages/Story";
import { Navigate, Route, Routes } from "react-router-dom";

export const ProtectedRoutes = () => {
  return (
    <SideBarLayout>
      <Routes>
        <Route path="/chat" element={<ChatPage />} />
        <Route path="/chat/:username" element={<RoomPage />} />
        <Route path="/story" element={<StoryPage />} />
        <Route path="/contact" element={<ContactPage />} />
        <Route path="/addcontact" element={<AddContactPage />} />
        <Route path="/setting" element={<SettingPage />} />
        <Route path="*" element={<Navigate to="/chat" replace />} />
      </Routes>
    </SideBarLayout>
  );
};
