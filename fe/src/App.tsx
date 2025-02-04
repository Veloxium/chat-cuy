import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import LandingPage from "./pages/Landing";
import { dynamicTitle } from "./utils/dynamicTitle";
import { Toaster } from "./components/ui/sonner";
import HomePage from "./pages/Home";

function App() {
  dynamicTitle();
  return (
    <>
      <div className="App">
        <Toaster />
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<LandingPage />} />
            <Route path="/home" element={<HomePage />} />
          </Routes>
        </BrowserRouter>
      </div>
    </>
  );
}

export default App;
