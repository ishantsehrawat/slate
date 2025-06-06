// src/Tiptap.tsx
import MenuBar from "@/components/layout/MenuBar";
import CodeBlock from "@tiptap/extension-code-block";
import CodeBlockLowlight from "@tiptap/extension-code-block-lowlight";
import { all, createLowlight } from "lowlight";
import Dropcursor from "@tiptap/extension-dropcursor";
import Placeholder from "@tiptap/extension-placeholder";
import { useEditor, EditorContent, Editor } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";

const lowlight = createLowlight(all);

const extensions = [
  StarterKit,
  Dropcursor,
  Placeholder.configure({
    placeholder: "Write something …",
  }),
  CodeBlock.configure({
    exitOnArrowDown: false,
    defaultLanguage: "javascript",
  }),
  CodeBlockLowlight.configure({
    lowlight,
    exitOnArrowDown: true,
  }),
];

import { useState } from "react";
import { Command, Info } from "lucide-react";
import { useKeyCombo } from "@/hooks/useKeyCombo";
const Home = () => {
  const [editorState, setEditorState] = useState("");
  const [heading, setHeading] = useState("");

  const generateHeadingShortcut = useKeyCombo(["mod", "g"]);

  const editor: Editor | null = useEditor({
    autofocus: true,
    extensions,
    content: editorState,
    onUpdate: ({ editor }) => {
      setEditorState(editor.getHTML());
    },
  });

  return (
    <div className="flex justify-center pb-32 pt-16">
      <div className="w-[794px] prose">
        <div className="flex flex-col gap-6">
          <input
            type="text"
            value={heading}
            placeholder="Start writing, we'll generate the heading"
            onChange={(e) => setHeading(e.target.value)}
            className="outline-none text-4xl w-full"
            aria-hidden="true"
          />
          <span className="flex items-center gap-1 text-xs">
            <Info className="h-5 w-5 mr-2" /> Press{" "}
            <span className="flex items-center gap-1 outline-1 rounded py-1 px-2">
              {generateHeadingShortcut}
            </span>{" "}
            to generate heading again
          </span>
        </div>
        <EditorContent editor={editor} />
        <MenuBar editor={editor} />
      </div>
    </div>
  );
};

export default Home;
