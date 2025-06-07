import {
  Menubar,
  MenubarContent,
  MenubarItem,
  MenubarMenu,
  MenubarSeparator,
  MenubarTrigger,
} from "@/components/ui/menubar";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { Bolt, LogIn, LogOut, SquarePlus } from "lucide-react";
import { useState } from "react";
import { SettingsDialog } from "../dialogs/SettingsDialog";
import KeyboardShortcutDialog from "../dialogs/KeyboardShortcutDialog";
import { googleLogin } from "@/API/api";

export function YourMenu() {
  const [isSettingsOpen, setIsSettingsOpen] = useState(false);
  const [isShortcutOpen, setShortcutOpen] = useState(false);

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
            <MenubarItem onClick={() => setIsSettingsOpen((prev) => !prev)}>
              <Bolt className="h-4 w-4" />
              Settings
            </MenubarItem>
            <MenubarItem onClick={() => setShortcutOpen((prev) => !prev)}>
              <SquarePlus className="h-4 w-4" />
              Keyboard Shortcuts
            </MenubarItem>
            <MenubarSeparator />
            <MenubarItem onClick={googleLogin}>
              <LogIn className="h-4 w-4" />
              Login
            </MenubarItem>
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

      <KeyboardShortcutDialog
        isOpen={isShortcutOpen}
        onOpenChange={setShortcutOpen}
      />
    </>
  );
}
