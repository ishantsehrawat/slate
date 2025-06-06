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

// Menu items.
const items = [
  {
    title: "New Slate",
    url: "#",
    icon: SquarePen,
  },
  {
    title: "Search Slates",
    url: "#",
    icon: Search,
  },
  {
    title: "Calendar",
    url: "#",
    icon: Calendar,
  },
];

const slates = [
  {
    label: "Today",
    slates: [
      "Morning Reflections",
      "Work Notes",
      "Evening Gratitude",
      "Daily Goals",
    ],
  },
  {
    label: "Previous 7 Days",
    slates: [
      "Weekly Goals Review",
      "Project Updates",
      "Personal Highlights",
      "Team Feedback",
      "Weekly Retrospective",
    ],
  },
  {
    label: "Last Month",
    slates: [
      "Monthly Summary",
      "Lessons Learned",
      "Upcoming Plans",
      "Budget Review",
    ],
  },
  {
    label: "March",
    slates: [
      "Spring Planning",
      "March Achievements",
      "Ideas for Growth",
      "Event Preparations",
      "March Retrospective",
      "New Initiatives",
    ],
  },
  {
    label: "February",
    slates: ["Winter Wrap-Up", "February Highlights", "Goals for March"],
  },
];

export function AppSidebar() {
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
        {slates.map((slateGroup) => (
          <SidebarGroup>
            <SidebarGroupLabel>{slateGroup.label}</SidebarGroupLabel>
            {slateGroup.slates.map((slate) => (
              <SidebarMenuButton>{slate}</SidebarMenuButton>
            ))}
          </SidebarGroup>
        ))}
      </SidebarContent>
    </Sidebar>
  );
}
