export interface IJournalEntry {
  id: number;
  title: string;
  hash: string;
}

export interface IJournal {
  title: string;
  content: string;
}

export interface IJournalGroup {
  label: string;
  journals: IJournalEntry[];
}

export interface LogoutResponse {
  message: string;
}
