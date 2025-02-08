import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import LandingPage from "./pages/Landing";
import { dynamicTitle } from "./utils/dynamicTitle";
import { Toaster } from "./components/ui/sonner";
import ChatPage from "./pages/Chat";

function App() {
  dynamicTitle();
  return (
    <>
      <div className="App">
        <Toaster />
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<LandingPage />} />
            <Route path="/chat" element={<ChatPage />} />
          </Routes>
        </BrowserRouter>
      </div>
    </>
  );
}

export default App;
