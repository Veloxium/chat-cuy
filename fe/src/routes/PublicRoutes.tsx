import LandingPage from "@/pages/Landing";
import TestingPage from "@/pages/Testing";
import { Navigate, Route, Routes } from "react-router-dom";

export const PublicRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/testing" element={<TestingPage />} />
      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  );
};
