import { Share } from "lucide-react";
import { Button } from "../ui/button";
import { SidebarTrigger } from "../ui/sidebar";

import { YourMenu } from "./YourMenu";

export default function Header({ isSidebarOpen }: { isSidebarOpen: boolean }) {
  return (
    <header
      className={`h-14 px-6 flex items-center justify-between border-b shadow-sm fixed top-0 transition-[left,width] duration-200 ${
        isSidebarOpen ? "left-[16rem] w-[calc(100%-16rem)]" : "left-0 w-full"
      } z-10 app-background`}
    >
      <div className="flex items-center gap-2">
        {!isSidebarOpen && <SidebarTrigger />}
        <h1 className="text-xl">Slate</h1>
      </div>

      <div className="flex gap-4 items-center">
        <Button variant="outline" size="sm">
          <Share /> Share
        </Button>
        <YourMenu />
      </div>
    </header>
  );
}
