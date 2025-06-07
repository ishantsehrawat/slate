export interface IJournalEntry {
  id: number;
  title: string;
  createdAt: string;
}

export interface IJournalGroup {
  label: string;
  journals: IJournalEntry[];
}
