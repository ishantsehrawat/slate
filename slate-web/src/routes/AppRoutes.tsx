// src/routes/AppRoutes.tsx
import { Routes, Route } from "react-router-dom";
import AppLayout from "@/components/layout/AppLayout";

import Home from "@/pages/Home";
import AuthAware from "@/pages/AuthAware";

export default function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<AuthAware />} />
      <Route
        path="/:id"
        element={
          <AppLayout>
            <Home />
          </AppLayout>
        }
      />
    </Routes>
  );
}
