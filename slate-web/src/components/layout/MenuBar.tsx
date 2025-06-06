import { useState } from "react";
import { Editor } from "@tiptap/react";
import {
  Bold,
  Code,
  Italic,
  Underline,
  X,
  Ellipsis,
  List,
  ListOrdered,
  Quote,
  Strikethrough,
} from "lucide-react";
import { ToggleGroup, ToggleGroupItem } from "../ui/toggle-group";
import { Toggle } from "../ui/toggle";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";

const MenuBar = ({ editor }: { editor: Editor | null }) => {
  const [closeMenuBar, setCloseMenuBar] = useState(false);
  return (
    <div className="flex w-[794px] items-center justify-center fixed bottom-10 left-[60%] transform -translate-x-1/2 gap-4">
      {!closeMenuBar && (
        <>
          <ToggleGroup
            type="multiple"
            variant="outline"
            className="app-background"
          >
            <ToggleGroupItem
              value="bold"
              aria-label="Toggle bold"
              onClick={() => editor?.chain().focus().toggleBold().run()}
            >
              <Bold className="h-4 w-4" />
            </ToggleGroupItem>
            <ToggleGroupItem
              value="italic"
              aria-label="Toggle italic"
              onClick={() => editor?.chain().focus().toggleItalic().run()}
              disabled={!editor?.can().chain().focus().toggleItalic().run()}
            >
              <Italic className="h-4 w-4" />
            </ToggleGroupItem>
            {/* <ToggleGroupItem
              value="strikethrough"
              aria-label="Toggle strikethrough"
              onClick={() => editor?.chain().focus().toggleUnderline().run()}
              disabled={!editor?.can().chain().focus().toggleUnderline().run()}
            >
              <Underline className="h-4 w-4" />
            </ToggleGroupItem> */}
            <ToggleGroupItem
              value="code"
              aria-label="Toggle Code"
              onClick={() => editor?.chain().focus().toggleCode().run()}
              disabled={!editor?.can().chain().focus().toggleCode().run()}
            >
              <Code className="h-4 w-4" />
            </ToggleGroupItem>
            <ToggleGroupItem
              value="List"
              aria-label="Toggle List"
              onClick={() => editor?.chain().focus().toggleBulletList().run()}
              disabled={!editor?.can().chain().focus().toggleBulletList().run()}
            >
              <List className="h-4 w-4" />
            </ToggleGroupItem>
            <ToggleGroupItem
              value="ListOrdered"
              aria-label="Toggle ListOrdered"
              onClick={() => editor?.chain().focus().toggleOrderedList().run()}
              disabled={
                !editor?.can().chain().focus().toggleOrderedList().run()
              }
            >
              <ListOrdered className="h-4 w-4" />
            </ToggleGroupItem>
            <ToggleGroupItem
              value="Quote"
              aria-label="Toggle Quote"
              onClick={() => editor?.chain().focus().toggleBlockquote().run()}
              disabled={!editor?.can().chain().focus().toggleBlockquote().run()}
            >
              <Quote className="h-4 w-4" />
            </ToggleGroupItem>
            <ToggleGroupItem
              value="Strikethrough"
              aria-label="Toggle Strikethrough"
              onClick={() => editor?.chain().focus().toggleStrike().run()}
              disabled={!editor?.can().chain().focus().toggleStrike().run()}
            >
              <Strikethrough className="h-4 w-4" />
            </ToggleGroupItem>
          </ToggleGroup>
          <Select
            onValueChange={(value) => {
              if (value === "h2")
                editor?.chain().focus().setHeading({ level: 4 }).run();
              if (value === "h3")
                editor?.chain().focus().setHeading({ level: 5 }).run();
              if (value === "h4")
                editor?.chain().focus().setHeading({ level: 6 }).run();
              if (value === "p") editor?.chain().focus().setParagraph().run();
            }}
          >
            <SelectTrigger className="w-[180px] app-background">
              <SelectValue placeholder="Select a style" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectItem value="h2">Heading 1</SelectItem>
                <SelectItem value="h3">Heading 2</SelectItem>
                <SelectItem value="h4">Heading 3</SelectItem>
                <SelectItem value="p">Text</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
        </>
      )}
      <Toggle
        variant="outline"
        className="app-background"
        value="close"
        onClick={() => setCloseMenuBar(!closeMenuBar)}
      >
        {closeMenuBar ? (
          <Ellipsis className="h-4 w-4" />
        ) : (
          <X className="h-4 w-4" />
        )}
      </Toggle>
    </div>
  );
};

export default MenuBar;
