// src/pages/AuthAwarePage.tsx
import Cookies from "js-cookie";
import Login from "./Login";
import Home from "./Home";
import AppLayout from "@/components/layout/AppLayout";

export default function AuthAware() {
  const token = Cookies.get("auth_token");

  if (token) {
    // User is authenticated, show Home inside layout
    return (
      <AppLayout>
        <Home />
      </AppLayout>
    );
  }

  // No token, show login screen
  return <Login />;
}
