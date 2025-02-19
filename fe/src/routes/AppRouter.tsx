import { useUserStore } from "@/store/userStore";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { ProtectedRoutes } from "./ProtectedRoutes";
import { PublicRoutes } from "./PublicRoutes";


export const AppRouter = () => {
  const user = useUserStore((state) => state.user);
  return (
    <BrowserRouter>
      <Routes>
        {user?.username ? (
          <Route path="/*" element={<ProtectedRoutes />} />
        ) : (
          <Route path="/*" element={<PublicRoutes />} />
        )}

        <Route path="*" element={<Navigate to="/login" replace />} />
      </Routes>
    </BrowserRouter>
  );
};
