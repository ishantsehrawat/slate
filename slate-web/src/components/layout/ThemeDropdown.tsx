import { useTheme } from "@/context/ThemeProvider";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";

export default function ThemeDropdown() {
  const { theme, setTheme } = useTheme();

  return (
    <div className="text-sm">
      <Select value={theme} onValueChange={(value) => setTheme(value as any)}>
        <SelectTrigger className="rounded border px-2 py-1 w-32">
          <SelectValue placeholder="Select theme" />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="system">System</SelectItem>
          <SelectItem value="light">Light</SelectItem>
          <SelectItem value="dark">Dark</SelectItem>
        </SelectContent>
      </Select>
    </div>
  );
}
