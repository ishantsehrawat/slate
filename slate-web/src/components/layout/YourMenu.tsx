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
import KeyboardShortcutDialog from "../dialogs/KeyboardShortcutDialog";
import { createJournal, logout } from "@/API/api";
import type { IJournal } from "@/interfaces/Journal";
import { useNavigate } from "react-router-dom";
import { toast } from "sonner";

export function YourMenu() {
  const navigate = useNavigate();
  const [isSettingsOpen, setIsSettingsOpen] = useState(false);
  const [isShortcutOpen, setShortcutOpen] = useState(false);

  const writeJournal = () => {
    const journal: IJournal = {
      title: "siduoghiskdjbpowuildvks",
      content: "piuvoihxjc ]asdib x0fibpjosdoubh ioudhbp9duifbhdjp 9oudihj",
    };

    createJournal(journal).then(
      (data) => {
        console.log(data);
      },
      (err) => {
        console.error(err);
      }
    );
  };

  const handleLogout = async () => {
    const { message } = await logout();
    console.log(message);
    toast.success(message);

    navigate("/", { replace: true });
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
            <MenubarItem onClick={() => setIsSettingsOpen((prev) => !prev)}>
              <Bolt className="h-4 w-4" />
              Settings
            </MenubarItem>
            <MenubarItem onClick={() => setShortcutOpen((prev) => !prev)}>
              <SquarePlus className="h-4 w-4" />
              Keyboard Shortcuts
            </MenubarItem>
            <MenubarItem onClick={writeJournal}>
              <SquarePlus className="h-4 w-4" />
              Create journal
            </MenubarItem>
            <MenubarSeparator />
            <MenubarItem onClick={handleLogout}>
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
