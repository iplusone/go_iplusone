import fs from "node:fs/promises";
import path from "node:path";
import matter from "gray-matter";
import MarkdownIt from "markdown-it";

const md = new MarkdownIt({
  html: false,
  linkify: true,
  typographer: false
});

export function parseArgs(argv) {
  const args = {
    headless: false,
    publish: false,
    file: ""
  };

  for (let i = 0; i < argv.length; i++) {
    const token = argv[i];
    if (token === "--headless") {
      args.headless = true;
      continue;
    }
    if (token === "--publish") {
      args.publish = true;
      continue;
    }
    if (token === "--file") {
      args.file = argv[i + 1] || "";
      i++;
      continue;
    }
    if (token === "--help" || token === "-h") {
      args.help = true;
    }
  }

  return args;
}

export async function ensureDir(dirPath) {
  await fs.mkdir(dirPath, { recursive: true });
}

export async function loadArticle(filePath) {
  const raw = await fs.readFile(filePath, "utf8");
  const parsed = matter(raw);
  const content = parsed.content.trim();
  const title = parsed.data.title || extractTitle(content) || path.basename(filePath, path.extname(filePath));
  const html = md.render(content);
  return {
    filePath,
    slug: toSlug(path.basename(filePath, path.extname(filePath))),
    title,
    markdown: content,
    html,
    data: parsed.data || {}
  };
}

function extractTitle(content) {
  const line = content.split("\n").find((entry) => entry.startsWith("# "));
  return line ? line.slice(2).trim() : "";
}

function toSlug(value) {
  return value
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, "-")
    .replace(/^-+|-+$/g, "");
}

export function resolveConfig() {
  return {
    baseUrl: process.env.NOTE_BASE_URL || "https://note.com",
    newPostUrl: process.env.NOTE_NEW_POST_URL || "https://note.com/notes/new",
    titleSelector:
      process.env.NOTE_TITLE_SELECTOR ||
      'textarea, input[placeholder*="タイトル"], input[placeholder*="title"]',
    editorSelector:
      process.env.NOTE_EDITOR_SELECTOR ||
      '[contenteditable="true"]',
    draftButtonText: process.env.NOTE_DRAFT_BUTTON_TEXT || "下書き",
    publishButtonText: process.env.NOTE_PUBLISH_BUTTON_TEXT || "公開",
    storagePath: path.resolve("storage/note-auth.json"),
    artifactDir: path.resolve("artifacts")
  };
}

export function printHelp() {
  console.log(`Usage:
  npm run draft -- --file <markdown-path> [--headless]
  npm run login
`);
}
