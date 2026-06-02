import fs from "node:fs/promises";
import path from "node:path";
import { ensureDir, loadArticle, parseArgs } from "./lib.mjs";

const args = parseArgs(process.argv.slice(2));
const root = path.resolve("../../docs/note-handson-code-articles");
const outDir = path.resolve("dist/note-ready");

await ensureDir(outDir);

const targets = args.file
  ? [path.resolve(args.file)]
  : (await fs.readdir(root))
      .filter((name) => name.endsWith(".md"))
      .sort()
      .map((name) => path.join(root, name));

for (const filePath of targets) {
  const article = await loadArticle(filePath);
  const txt = convertMarkdownToNoteText(article.markdown);
  const fragmentHtml = wrapFragmentHtml(article.title, article.html);
  const html = wrapHtml(article.title, fragmentHtml);
  const base = path.join(outDir, path.basename(filePath, ".md"));

  await fs.writeFile(`${base}.txt`, txt, "utf8");
  await fs.writeFile(`${base}.fragment.html`, fragmentHtml, "utf8");
  await fs.writeFile(`${base}.html`, html, "utf8");
  console.log(`built: ${path.basename(base)}.{txt,fragment.html,html}`);
}

function convertMarkdownToNoteText(markdown) {
  const lines = markdown.split("\n");
  const out = [];
  let inCode = false;

  for (const rawLine of lines) {
    const line = rawLine.replace(/\t/g, "  ");

    if (line.startsWith("```")) {
      if (!inCode) {
        inCode = true;
        out.push("----- code -----");
      } else {
        inCode = false;
        out.push("----- /code -----");
        out.push("");
      }
      continue;
    }

    if (inCode) {
      out.push(line);
      continue;
    }

    if (line.startsWith("# ")) {
      out.push(line.slice(2).trim());
      out.push("");
      continue;
    }

    if (line.startsWith("## ")) {
      out.push(`【${line.slice(3).trim()}】`);
      out.push("");
      continue;
    }

    if (line.startsWith("### ")) {
      out.push(`■ ${line.slice(4).trim()}`);
      continue;
    }

    if (/^\s*[-*]\s+/.test(line)) {
      out.push(line.replace(/^\s*[-*]\s+/, "・"));
      continue;
    }

    if (/^\s*\d+\.\s+/.test(line)) {
      out.push(line.trim());
      continue;
    }

    if (line.trim() === "") {
      if (out.length > 0 && out[out.length - 1] !== "") {
        out.push("");
      }
      continue;
    }

    out.push(stripInlineMarkdown(line));
  }

  return cleanupTrailingBlankLines(out).join("\n");
}

function stripInlineMarkdown(line) {
  return line
    .replace(/`([^`]+)`/g, "$1")
    .replace(/\*\*([^*]+)\*\*/g, "$1")
    .replace(/\*([^*]+)\*/g, "$1")
    .replace(/\[([^\]]+)\]\([^)]+\)/g, "$1");
}

function cleanupTrailingBlankLines(lines) {
  const cleaned = [...lines];
  while (cleaned.length > 0 && cleaned[cleaned.length - 1] === "") {
    cleaned.pop();
  }
  return cleaned;
}

function wrapHtml(title, bodyHtml) {
  return `<!doctype html>
<html lang="ja">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>${escapeHtml(title)}</title>
    <style>
      body {
        max-width: 760px;
        margin: 40px auto;
        padding: 0 20px 60px;
        font-family: -apple-system, BlinkMacSystemFont, "Hiragino Sans", "Yu Gothic", sans-serif;
        line-height: 1.9;
        color: #222;
      }
      h1, h2, h3 {
        line-height: 1.5;
      }
      pre {
        overflow-x: auto;
        padding: 16px;
        border: 1px solid #ddd;
        border-radius: 8px;
        background: #fafafa;
      }
      code {
        font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
      }
      blockquote {
        margin-left: 0;
        padding-left: 16px;
        border-left: 4px solid #ddd;
        color: #555;
      }
    </style>
  </head>
  <body>
${bodyHtml}
  </body>
</html>
`;
}

function wrapFragmentHtml(title, bodyHtml) {
  return `<article data-note-ready="true">
<h1>${escapeHtml(title)}</h1>
${bodyHtml
  .replace(/^<h1>.*?<\/h1>\n?/s, "")
  .replace(/<h2>/g, '<h2 data-note-heading="section">')
  .replace(/<pre><code class="language-([^"]+)">/g, '<pre data-note-code="true"><code data-lang="$1">')}
</article>`;
}

function escapeHtml(value) {
  return value
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;");
}
