#!/usr/bin/env node
const yargs = require("yargs/yargs");
const { hideBin } = require("yargs/helpers");
const { parse, stringify } = require("yaml");
const fs = require("fs");
const simpleLogger = require("simple-node-logger");
const { runSDKTransformations, runFlatteningTransformations } = require("./atlasTransformations");

function readInput({ input, stdin }) {
  if (stdin) {
    return new Promise((resolve, reject) => {
      let data = "";
      process.stdin.setEncoding("utf8");
      process.stdin.on("data", (chunk) => (data += chunk));
      process.stdin.on("end", () => {
        try {
          resolve(parse(data));
        } catch (e) {
          reject(e);
        }
      });
      process.stdin.on("error", reject);
    });
  }
  if (!input) {
    throw new Error("Missing --input or use --stdin");
  }
  return Promise.resolve(parse(fs.readFileSync(input, "utf8")));
}

function writeOutput({ doc, output, stdout, input }) {
  const yaml = stringify(doc, { lineWidth: 9999 });
  if (stdout || (!output && !input)) {
    process.stdout.write(yaml);
    return;
  }
  const outPath = output || input;
  fs.writeFileSync(outPath, yaml);
}

function configureLogger(level) {
  const log = simpleLogger.createSimpleLogger();
  log.setLevel(level || process.env.XGEN_LOGGING_LEVEL || "warn");
  return log;
}

yargs(hideBin(process.argv))
  .command(
    "flatten",
    "Apply flattening transformations",
    (y) =>
      y
        .option("input", { alias: "i", type: "string" })
        .option("output", { alias: "o", type: "string" })
        .option("stdin", { type: "boolean", default: false })
        .option("stdout", { type: "boolean", default: false })
        .option("log-level", { type: "string" }),
    async (args) => {
      const log = configureLogger(args["log-level"]);
      const originalConsole = global.console;
      global.console = log; // direct internal console usage to our logger during CLI execution
      try {
        const doc = await readInput(args);
        const out = runFlatteningTransformations(doc);
        writeOutput({ doc: out, ...args });
      } finally {
        global.console = originalConsole;
      }
    },
  )
  .command(
    "transform",
    "Apply full transformations",
    (y) =>
      y
        .option("input", { alias: "i", type: "string" })
        .option("output", { alias: "o", type: "string" })
        .option("stdin", { type: "boolean", default: false })
        .option("stdout", { type: "boolean", default: false })
        .option("log-level", { type: "string" }),
    async (args) => {
      const log = configureLogger(args["log-level"]);
      const originalConsole = global.console;
      global.console = log;
      try {
        const doc = await readInput(args);
        const out = runSDKTransformations(doc);
        writeOutput({ doc: out, ...args });
      } finally {
        global.console = originalConsole;
      }
    },
  )
  .demandCommand(1)
  .strict()
  .help()
  .parse();


