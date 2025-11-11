#!/usr/bin/env node
const yargs = require("yargs/yargs");
const { hideBin } = require("yargs/helpers");
const { getAPI, saveAPI } = require("./engine/apifile");
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
  .demandCommand(1)
  .strict()
  .help()
  .parse();
