import { BrowserRouter, Routes, Route } from "react-router-dom";

import TerminalPage from "./pages/TerminalPage";
import LoginPage from "./pages/LoginPage";
import FileExplorer from "./pages/FileExplorer";

import Navbar from "./components/NavBar";
import Footer from "./components/Footer";

export default function App() {
  return (
    <BrowserRouter>
      <div className="flex flex-col min-h-screen bg-gray-100">
        <Navbar />
        <main className="container mx-auto px-4 flex-grow ">
          <Routes>
            <Route path="/" element={<TerminalPage />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="/file-explorer" element={<FileExplorer />} />
          </Routes>
        </main>
        <Footer />
      </div>
    </BrowserRouter>
  );
}
