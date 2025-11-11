#!/usr/bin/env node
const { getAPI, saveAPI } = require("./engine/apifile");
const log = require("simple-node-logger").createSimpleLogger();
const {
  runFlatteningTransformations,
  runAllTransformations,
} = require("./atlasTransformations");

// Override default logger
global.console = log;
log.setLevel(process.env.XGEN_LOGGING_LEVEL || "warn");

function printUsageAndExit() {
  console.error(
    "Usage: atlas-openapi-transformer <flatten|transform> <inputFile> <outputFile>",
  );
  process.exit(1);
}

const args = process.argv.slice(2);
if (args.length !== 3) {
  printUsageAndExit();
}

const [command, inputFile, outputFile] = args;

let { doc } = getAPI([inputFile]);

switch (command) {
  case "flatten":
    doc = runFlatteningTransformations(doc);
    break;
  case "transform":
    doc = runAllTransformations(doc);
    break;
  default:
    console.error(`Unknown command: ${command}`);
    printUsageAndExit();
}

saveAPI(doc, outputFile);

process.exit(0);


