import React from "react";
import ReactDOM from "react-dom/client";
import AppLayout from "./components/layout/AppLayout";
import Home from "./pages/Home";
import "./styles/index.css"; // Tailwind should be imported here
import ThemeProvider from "./context/ThemeProvider"; // Import ThemeProvider

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ThemeProvider>
      {" "}
      {/* Wrap the application with ThemeProvider */}
      <AppLayout>
        <Home />
      </AppLayout>
    </ThemeProvider>
  </React.StrictMode>
);
