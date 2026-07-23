#!/usr/bin/env node
const yargs = require("yargs/yargs");
const { hideBin } = require("yargs/helpers");
const { getAPI, saveAPI } = require("./engine/apifile");
const { mergePreview } = require("./engine/mergePreview");
const simpleLogger = require("simple-node-logger");
const {
  runFlatteningTransformations,
  runAllTransformations,
} = require("./atlasTransformations");

function readInput({ input }) {
  if (!input) {
    throw new Error("Missing --input");
  }
  return getAPI([input]).doc;
}

function writeOutput({ doc, output }) {
  if (!output) {
    throw new Error("Missing --output");
  }
  saveAPI(doc, output);
}

function configureLogger(level) {
  const log = simpleLogger.createSimpleLogger();
  log.setLevel(level || "warn");
  return log;
}

yargs(hideBin(process.argv))
  .command(
    "flatten",
    "Apply flattening transformations",
    (y) =>
      y
        .option("input", { alias: "i", type: "string", demandOption: true })
        .option("output", { alias: "o", type: "string", demandOption: true })
        .option("log-level", { type: "string" }),
    async (args) => {
      const log = configureLogger(args["log-level"]);
      const originalConsole = global.console;
      global.console = log;
      try {
        const doc = readInput(args);
        const out = runFlatteningTransformations(doc);
        writeOutput({ doc: out, output: args.output });
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
        .option("input", { alias: "i", type: "string", demandOption: true })
        .option("output", { alias: "o", type: "string", demandOption: true })
        .option("log-level", { type: "string" }),
    async (args) => {
      const log = configureLogger(args["log-level"]);
      const originalConsole = global.console;
      global.console = log;
      try {
        const doc = readInput(args);
        const out = runAllTransformations(doc);
        writeOutput({ doc: out, output: args.output });
      } finally {
        global.console = originalConsole;
      }
    },
  )
  .command(
    "merge-preview",
    "Merge one or more preview delta specs onto a full base spec",
    (y) =>
      y
        .option("base", { alias: "b", type: "string", demandOption: true })
        .option("preview", {
          alias: "p",
          type: "array",
          demandOption: true,
          describe:
            "Preview delta file. Repeat --preview to layer multiple deltas " +
            "(e.g. the public preview spec plus one or more private preview " +
            "specs) on top of the base, applied in the order given.",
        })
        .option("output", { alias: "o", type: "string", demandOption: true })
        .option("log-level", { type: "string" }),
    async (args) => {
      const log = configureLogger(args["log-level"]);
      const originalConsole = global.console;
      global.console = log;
      try {
        let out = getAPI([args.base]).doc;
        for (const previewFile of args.preview) {
          const preview = getAPI([previewFile]).doc;
          out = mergePreview(out, preview);
        }
        writeOutput({ doc: out, output: args.output });
      } finally {
        global.console = originalConsole;
      }
    },
  )
  .demandCommand(1)
  .strict()
  .help()
  .parse();
