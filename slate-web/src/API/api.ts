import type { IJournalGroup } from "@/interfaces/Journal";

const BASE_URL = "http://api.slate.com:80";

export function googleLogin(): void {
  window.location.href = `${BASE_URL}/auth/google/login`;
}

// ✅ GET grouped journals (requires auth cookie)
export async function fetchJournalsGrouped(): Promise<IJournalGroup[]> {
  const res = await fetch(`${BASE_URL}/api/journals`, {
    credentials: "include", // ensures cookies (auth_token) are sent
  });

  if (!res.ok) throw new Error("Failed to fetch journals");
  return res.json();
}

// ✅ GET single journal (requires auth cookie)
export const getJournal = async (id: string) => {
  const res = await fetch(`${BASE_URL}/api/journals/${id}`, {
    credentials: "include",
  });

  if (!res.ok) throw new Error("Failed to fetch journal");
  return res.json();
};

// ✅ POST create journal (requires auth cookie)
export const createJournal = async (data: {
  title: string;
  content: string;
}) => {
  const res = await fetch(`${BASE_URL}/api/journals`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    credentials: "include",
    body: JSON.stringify(data),
  });

  if (!res.ok) throw new Error("Failed to create journal");
  return res.json();
};

// ✅ PUT update journal (requires auth cookie)
export const updateJournal = async (
  id: string,
  data: { title?: string; content?: string }
) => {
  const res = await fetch(`${BASE_URL}/api/journals/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    credentials: "include",
    body: JSON.stringify(data),
  });

  if (!res.ok) throw new Error("Failed to update journal");
  return res.json();
};
