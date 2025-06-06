import { useEffect, useState } from "react";

type Platform = "mac" | "windows";

const detectPlatform = (): Platform => {
  if (typeof navigator === "undefined") return "windows";
  return /Mac|iPhone|iPod|iPad/.test(navigator.platform) ? "mac" : "windows";
};

export function useKeyCombo(keys: string[]) {
  const [platform, setPlatform] = useState<Platform>("windows");

  useEffect(() => {
    setPlatform(detectPlatform());
  }, []);

  return keys
    .map((key) => {
      if (key.toLowerCase() === "mod") {
        return platform === "mac" ? "⌘" : "Ctrl";
      }
      if (key.toLowerCase() === "shift") {
        return "⇧";
      }
      if (key.toLowerCase() === "alt") {
        return platform === "mac" ? "⌥" : "Alt";
      }
      if (key.toLowerCase() === "enter") {
        return "↩";
      }
      return key.toUpperCase();
    })
    .join(" + ");
}
