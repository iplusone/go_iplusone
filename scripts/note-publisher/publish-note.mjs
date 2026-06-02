import path from "node:path";
import { chromium } from "playwright";
import {
  ensureDir,
  loadArticle,
  parseArgs,
  printHelp,
  resolveConfig
} from "./lib.mjs";

const args = parseArgs(process.argv.slice(2));
if (args.help || !args.file) {
  printHelp();
  process.exit(args.help ? 0 : 1);
}

const config = resolveConfig();
const article = await loadArticle(path.resolve(args.file));

await ensureDir(config.artifactDir);

const browser = await chromium.launch({ headless: args.headless });
const context = await browser.newContext({
  storageState: config.storagePath
});
const page = await context.newPage();

try {
  console.log(`Opening editor: ${config.newPostUrl}`);
  await page.goto(config.newPostUrl, { waitUntil: "domcontentloaded" });

  await fillTitle(page, article.title, config.titleSelector);
  await fillBody(page, article, config.editorSelector);

  await page.screenshot({
    path: path.join(config.artifactDir, `${article.slug}-edited.png`),
    fullPage: true
  });

  if (!args.publish) {
    await trySaveDraft(page, config.draftButtonText);
    console.log("Draft save flow attempted.");
  } else {
    console.log("Publish flag is set, but this first version still stops at the editor for safety.");
  }

  await page.screenshot({
    path: path.join(config.artifactDir, `${article.slug}-final.png`),
    fullPage: true
  });
} finally {
  await browser.close();
}

async function fillTitle(page, title, selector) {
  const locator = page.locator(selector).first();
  await locator.waitFor({ state: "visible", timeout: 15000 });
  await locator.click();
  await locator.fill("");
  await locator.type(title);
  console.log(`Filled title: ${title}`);
}

async function fillBody(page, article, selector) {
  const locators = page.locator(selector);
  const count = await locators.count();
  if (count === 0) {
    throw new Error(`Editor not found with selector: ${selector}`);
  }

  let editor = locators.nth(Math.max(0, count - 1));
  await editor.waitFor({ state: "visible", timeout: 15000 });
  await editor.click();

  await page.evaluate(
    async ({ html, text }) => {
      const clipboard = navigator.clipboard;
      if (!clipboard || !window.ClipboardItem) {
        throw new Error("Clipboard API is not available in this browser context.");
      }

      const item = new window.ClipboardItem({
        "text/plain": new Blob([text], { type: "text/plain" }),
        "text/html": new Blob([html], { type: "text/html" })
      });
      await clipboard.write([item]);
    },
    {
      html: article.html,
      text: article.markdown
    }
  );

  const modifier = process.platform === "darwin" ? "Meta" : "Control";
  await page.keyboard.press(`${modifier}+A`).catch(() => {});
  await page.keyboard.press("Backspace").catch(() => {});
  await page.keyboard.press(`${modifier}+V`);
  console.log("Pasted article body into editor.");
}

async function trySaveDraft(page, buttonText) {
  const buttons = page.getByRole("button", { name: new RegExp(buttonText) });
  const count = await buttons.count();
  if (count === 0) {
    console.warn(`Draft button not found by text: ${buttonText}`);
    return;
  }

  await buttons.first().click();
  await page.waitForTimeout(2000);
}
