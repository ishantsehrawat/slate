import type { DialogProps } from "@/interfaces/DialogProps";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { SHORTCUTS } from "@/constants/keyboardsShortcuts";
import { useKeyCombo } from "@/hooks/useKeyCombo";

export function KeyboardShortcutDialog({ isOpen, onOpenChange }: DialogProps) {
  return (
    <Dialog open={isOpen} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-3xl">
        <DialogHeader>
          <DialogTitle>Keyboard Shortcuts</DialogTitle>
        </DialogHeader>

        <Tabs defaultValue="editor" className="mt-4">
          <TabsList>
            <TabsTrigger value="editor">Editor Shortcuts</TabsTrigger>
            <TabsTrigger value="markdown">Markdown Syntax</TabsTrigger>
            <TabsTrigger value="app">App Shortcuts</TabsTrigger>
          </TabsList>

          <TabsContent value="editor" className="mt-4">
            <h3 className="text-lg font-semibold mb-2">Editor Shortcuts</h3>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Action</TableHead>
                  <TableHead>Shortcut</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {SHORTCUTS.editorShortcuts.map(({ action, shortcut }) => (
                  <TableRow key={action}>
                    <TableCell>{action}</TableCell>
                    <TableCell>
                      <kbd>{useKeyCombo(shortcut)}</kbd>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TabsContent>

          <TabsContent value="markdown" className="mt-4">
            <h3 className="text-lg font-semibold mb-2">Markdown Syntax</h3>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Action</TableHead>
                  <TableHead>Markdown</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {SHORTCUTS.markdownSyntax.map(({ action, syntax }) => (
                  <TableRow key={action}>
                    <TableCell>{action}</TableCell>
                    <TableCell>
                      <code>{syntax}</code>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TabsContent>

          <TabsContent value="app" className="mt-4">
            <h3 className="text-lg font-semibold mb-2">App Shortcuts</h3>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Action</TableHead>
                  <TableHead>Shortcut</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {SHORTCUTS.appShortcuts.map(({ action, shortcut }) => (
                  <TableRow key={action}>
                    <TableCell>{action}</TableCell>
                    <TableCell>
                      <kbd>{useKeyCombo(shortcut)}</kbd>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TabsContent>
        </Tabs>
      </DialogContent>
    </Dialog>
  );
}

export default KeyboardShortcutDialog;
