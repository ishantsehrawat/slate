import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { useTheme } from "@/context/ThemeProvider";
import { useState } from "react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import type { DialogProps } from "@/interfaces/DialogProps";

export function SettingsDialog({ isOpen, onOpenChange }: DialogProps) {
  const { theme, setTheme } = useTheme();

  const [autoSave, setAutoSave] = useState(true);
  const [autoGenerateHeading, setAutoGenerateHeading] = useState(true);
  const [editorWidth, setEditorWidth] = useState("fixed");
  const [fontSize, setFontSize] = useState("medium");

  const handleDeleteAll = () => {
    const confirmDelete = window.confirm(
      "Are you sure you want to delete all slates?"
    );
    if (confirmDelete) {
      console.log("All slates deleted.");
    }
  };

  return (
    <Dialog open={isOpen} onOpenChange={onOpenChange}>
      <DialogTrigger asChild>
        <Button variant="outline" className="hidden">
          Settings Dialog
        </Button>
      </DialogTrigger>

      <DialogContent className="sm:max-w-[700px]">
        <DialogHeader>
          <DialogTitle>Settings</DialogTitle>
          <DialogDescription>
            Customize how Slate behaves and appears.
          </DialogDescription>
        </DialogHeader>

        <div className="grid gap-8 py-4">
          {/* Appearance Section */}
          <section>
            <h3 className="text-lg font-medium mb-2">Appearance</h3>
            <div className="space-y-4">
              <div className="flex items-center justify-between">
                <Label>Theme</Label>
                <Select
                  value={theme}
                  onValueChange={(val) => setTheme(val as any)}
                >
                  <SelectTrigger className="w-[160px]">
                    <SelectValue placeholder="Select theme" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="system">System</SelectItem>
                    <SelectItem value="light">Light</SelectItem>
                    <SelectItem value="dark">Dark</SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <div className="flex items-center justify-between">
                <Label>Font Size</Label>
                <Select value={fontSize} onValueChange={setFontSize}>
                  <SelectTrigger className="w-[160px]">
                    <SelectValue placeholder="Font size" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="small">Small</SelectItem>
                    <SelectItem value="medium">Medium</SelectItem>
                    <SelectItem value="large">Large</SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <div className="flex items-center justify-between">
                <Label>Editor Width</Label>
                <Select value={editorWidth} onValueChange={setEditorWidth}>
                  <SelectTrigger className="w-[160px]">
                    <SelectValue placeholder="Editor width" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="fixed">Fixed (A4)</SelectItem>
                    <SelectItem value="fluid">Fluid (Full)</SelectItem>
                  </SelectContent>
                </Select>
              </div>
            </div>
          </section>

          {/* Behavior Section */}
          <section>
            <h3 className="text-lg font-medium mb-2">Behavior</h3>
            <div className="space-y-4">
              <div className="flex items-center justify-between">
                <Label htmlFor="auto-save">Auto Save</Label>
                <Switch
                  id="auto-save"
                  checked={autoSave}
                  onCheckedChange={setAutoSave}
                />
              </div>

              <div className="flex items-center justify-between">
                <Label htmlFor="auto-heading">Auto Generate Heading</Label>
                <Switch
                  id="auto-heading"
                  checked={autoGenerateHeading}
                  onCheckedChange={setAutoGenerateHeading}
                />
              </div>
            </div>
          </section>

          {/* Danger Zone */}
          <section>
            <h3 className="text-sm font-medium text-red-500 mb-2">
              Danger Zone
            </h3>
            <Button variant="destructive" onClick={handleDeleteAll}>
              Delete All Slates
            </Button>
          </section>
        </div>
      </DialogContent>
    </Dialog>
  );
}
