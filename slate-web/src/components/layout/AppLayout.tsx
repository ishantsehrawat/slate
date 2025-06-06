import React, { useState } from "react";
import { SidebarProvider } from "../ui/sidebar";
import { AppSidebar } from "./AppSidebar";
import Header from "./Header";
import { useTheme } from "@/context/ThemeProvider";

export default function AppLayout({ children }: { children: React.ReactNode }) {
  const { resolvedTheme } = useTheme(); // theme can be "system", but resolvedTheme is "dark"/"light"
  // Access theme and toggleTheme
  const [isSidebarOpen, setIsSidebarOpen] = useState(true);

  return (
    <SidebarProvider
      defaultOpen={isSidebarOpen}
      open={isSidebarOpen}
      onOpenChange={(open) => {
        setIsSidebarOpen(open);
      }}
    >
      <div className={`flex w-[100dvw] overflow-hidden ${resolvedTheme}`}>
        <AppSidebar />
        <main className="flex-1 pt-14 transition-[margin-left] duration-200">
          <Header isSidebarOpen={isSidebarOpen} />
          {children}
        </main>
      </div>
    </SidebarProvider>
  );
}
