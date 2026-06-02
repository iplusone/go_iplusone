import readline from "node:readline/promises";
import { stdin as input, stdout as output } from "node:process";
import { chromium } from "playwright";
import { ensureDir, resolveConfig } from "./lib.mjs";

const config = resolveConfig();

await ensureDir("storage");

const browser = await chromium.launch({ headless: false });
const context = await browser.newContext();
const page = await context.newPage();

console.log(`Opening ${config.baseUrl}`);
console.log("Log in to note in the opened browser.");
console.log("After login is complete, return here and press Enter.");

await page.goto(config.baseUrl, { waitUntil: "domcontentloaded" });

const rl = readline.createInterface({ input, output });
await rl.question("");
rl.close();

await context.storageState({ path: config.storagePath });
console.log(`Saved login state to ${config.storagePath}`);

await browser.close();
