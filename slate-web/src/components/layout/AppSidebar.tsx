import { useEffect, useState } from "react";
import { Calendar, Search, Square, SquarePen } from "lucide-react";
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarTrigger,
} from "@/components/ui/sidebar";
import { fetchJournalsGrouped } from "@/API/api";
import type { IJournalGroup } from "@/interfaces/Journal";
import { useNavigate } from "react-router-dom";

// Static menu items
const items = [
  { title: "New Slate", url: "/0", icon: SquarePen },
  { title: "Search Slates", url: "#", icon: Search },
  { title: "Calendar", url: "#", icon: Calendar },
];

export function AppSidebar() {
  const navigate = useNavigate();
  const [groupedJournals, setGroupedJournals] = useState<IJournalGroup[]>([]);

  useEffect(() => {
    fetchJournalsGrouped()
      .then((data: IJournalGroup[]) => {
        setGroupedJournals(data);
      })
      .catch((err) => console.error("Error fetching journals:", err));
  }, []);

  const openJournal = (id: string) => {
    navigate(`/${id}`, { replace: true });
  };

  return (
    <Sidebar>
      <SidebarHeader>
        <div className="flex justify-between items-center px-1">
          <Square size={24} />
          <SidebarTrigger />
        </div>
      </SidebarHeader>

      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <a href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </a>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>

        {/* Render journal groups from backend */}
        {groupedJournals?.map(({ label, journals }) => (
          <SidebarGroup key={label}>
            <SidebarGroupLabel>{label}</SidebarGroupLabel>
            {journals?.map((entry) => (
              <SidebarMenuButton
                key={entry.id}
                onClick={() => openJournal(entry?.hash)}
              >
                {entry.title}
              </SidebarMenuButton>
            ))}
          </SidebarGroup>
        ))}
      </SidebarContent>
    </Sidebar>
  );
}
