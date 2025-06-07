// src/Tiptap.tsx
import MenuBar from "@/components/layout/MenuBar";
import CodeBlock from "@tiptap/extension-code-block";
import CodeBlockLowlight from "@tiptap/extension-code-block-lowlight";
import { all, createLowlight } from "lowlight";
import Dropcursor from "@tiptap/extension-dropcursor";
import Placeholder from "@tiptap/extension-placeholder";
import { useEditor, EditorContent, Editor } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Link from "@tiptap/extension-link";

import { useState } from "react";
import { Info } from "lucide-react";
import { useKeyCombo } from "@/hooks/useKeyCombo";

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
  Link.configure({
    openOnClick: false,
    autolink: true,
    defaultProtocol: "https",
    protocols: ["http", "https"],
    isAllowedUri: (url, ctx) => {
      try {
        // construct URL
        const parsedUrl = url.includes(":")
          ? new URL(url)
          : new URL(`${ctx.defaultProtocol}://${url}`);

        // use default validation
        if (!ctx.defaultValidate(parsedUrl.href)) {
          return false;
        }

        // disallowed protocols
        const disallowedProtocols = ["ftp", "file", "mailto"];
        const protocol = parsedUrl.protocol.replace(":", "");

        if (disallowedProtocols.includes(protocol)) {
          return false;
        }

        // only allow protocols specified in ctx.protocols
        const allowedProtocols = ctx.protocols.map((p) =>
          typeof p === "string" ? p : p.scheme
        );

        if (!allowedProtocols.includes(protocol)) {
          return false;
        }

        // disallowed domains
        const disallowedDomains = [
          "example-phishing.com",
          "malicious-site.net",
        ];
        const domain = parsedUrl.hostname;

        if (disallowedDomains.includes(domain)) {
          return false;
        }

        // all checks have passed
        return true;
      } catch {
        return false;
      }
    },
    shouldAutoLink: (url) => {
      try {
        // construct URL
        const parsedUrl = url.includes(":")
          ? new URL(url)
          : new URL(`https://${url}`);

        // only auto-link if the domain is not in the disallowed list
        const disallowedDomains = [
          "example-no-autolink.com",
          "another-no-autolink.com",
        ];
        const domain = parsedUrl.hostname;

        return !disallowedDomains.includes(domain);
      } catch {
        return false;
      }
    },
  }),
];
const Home = () => {
  const [editorState, setEditorState] = useState("");
  const [heading, setHeading] = useState("");

  const generateHeadingShortcut = useKeyCombo(["mod", "g"]);

  const setSlateHeading = (value: string) => {
    setHeading(value);
    document.title = value && "Slate";
  };

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
            onChange={(e) => setSlateHeading(e.target.value)}
            className="outline-none text-4xl font-bold w-full"
            aria-hidden="true"
          />
          <span className="flex items-center gap-1 text-xs">
            <Info className="h-4 w-4 mr-2" /> Press{" "}
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
