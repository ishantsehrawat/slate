import {
  createContext,
  useState,
  useEffect,
  useContext,
  useCallback,
} from "react";
import type { ReactNode } from "react";

type ThemeMode = "light" | "dark" | "system";
type ResolvedTheme = "light" | "dark";

interface ThemeContextType {
  theme: ThemeMode;
  resolvedTheme: ResolvedTheme;
  setTheme: (theme: ThemeMode) => void;
}

const ThemeContext = createContext<ThemeContextType>({
  theme: "system",
  resolvedTheme: "light",
  setTheme: () => {},
});

export const useTheme = () => useContext(ThemeContext);

const ThemeProvider = ({ children }: { children: ReactNode }) => {
  const getPreferredTheme = (): ThemeMode => {
    const stored = localStorage.getItem("theme") as ThemeMode | null;
    return stored || "system";
  };

  const [theme, setThemeState] = useState<ThemeMode>(() => getPreferredTheme());
  const [resolvedTheme, setResolvedTheme] = useState<ResolvedTheme>("light");

  const applyResolvedTheme = useCallback((mode: ResolvedTheme) => {
    document.documentElement.className = mode;
    setResolvedTheme(mode);
  }, []);

  const setTheme = (newTheme: ThemeMode) => {
    localStorage.setItem("theme", newTheme);
    setThemeState(newTheme);

    if (newTheme === "system") {
      const prefersDark = window.matchMedia(
        "(prefers-color-scheme: dark)"
      ).matches;
      applyResolvedTheme(prefersDark ? "dark" : "light");
    } else {
      applyResolvedTheme(newTheme);
    }
  };

  useEffect(() => {
    setTheme(getPreferredTheme()); // Apply on mount

    const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)");
    const handleChange = (e: MediaQueryListEvent) => {
      if (theme === "system") {
        applyResolvedTheme(e.matches ? "dark" : "light");
      }
    };

    mediaQuery.addEventListener("change", handleChange);
    return () => mediaQuery.removeEventListener("change", handleChange);
  }, [theme, applyResolvedTheme]);

  return (
    <ThemeContext.Provider value={{ theme, resolvedTheme, setTheme }}>
      {children}
    </ThemeContext.Provider>
  );
};

export default ThemeProvider;
