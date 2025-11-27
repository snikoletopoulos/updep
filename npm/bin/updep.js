#!/usr/bin/env node

const { spawn } = require("child_process");
const path = require("path");
const os = require("os");

const binaryPath = path.join(
  __dirname,
  `updep${os.platform() === "win32" ? ".exe" : ""}`
);

const child = spawn(binaryPath, process.argv.slice(2), {
  stdio: "inherit",
  windowsHide: true,
});

child.on("exit", (code, signal) => {
  if (signal) {
    process.kill(process.pid, signal);
  } else {
    process.exit(code);
  }
});

child.on("error", err => {
  console.error(`Failed to start updep: ${err.message}`);
  process.exit(1);
});
