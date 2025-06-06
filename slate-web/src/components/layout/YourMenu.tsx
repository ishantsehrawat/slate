import {
  Menubar,
  MenubarContent,
  MenubarItem,
  MenubarMenu,
  MenubarSeparator,
  MenubarTrigger,
} from "@/components/ui/menubar";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { Bolt, LogOut, SquarePlus } from "lucide-react";
import { useState } from "react";
import { SettingsDialog } from "../dialogs/SettingsDialog";

export function YourMenu() {
  const [isSettingsOpen, setIsSettingsOpen] = useState(false);

  const toggleSettings = () => {
    setIsSettingsOpen((prev) => !prev);
  };

  return (
    <>
      <Menubar>
        <MenubarMenu>
          <MenubarTrigger>
            <Avatar>
              <AvatarImage src="https://github.com/shadcn.png" />
              <AvatarFallback>CN</AvatarFallback>
            </Avatar>
          </MenubarTrigger>
          <MenubarContent>
            <MenubarItem disabled>ishantsehrawat75@gmail.com</MenubarItem>
            <MenubarSeparator />
            <MenubarItem onClick={toggleSettings}>
              <Bolt className="h-4 w-4" />
              Settings
            </MenubarItem>
            <MenubarItem>
              <SquarePlus className="h-4 w-4" />
              Keyboard Shortcuts
            </MenubarItem>
            <MenubarSeparator />
            <MenubarItem>
              <LogOut className="h-4 w-4" />
              Logout
            </MenubarItem>
          </MenubarContent>
        </MenubarMenu>
      </Menubar>
      <SettingsDialog
        isOpen={isSettingsOpen}
        onOpenChange={setIsSettingsOpen}
      />
    </>
  );
}
