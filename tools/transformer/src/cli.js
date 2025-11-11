#!/usr/bin/env node
const { getAPI, saveAPI } = require("./engine/apifile");
const {
  runFlatteningTransformations,
  runAllTransformations,
} = require("./atlasTransformations");

const log = require("simple-node-logger").createSimpleLogger();
// Override default logger
global.console = log;
log.setLevel("debug");

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

let doc;
try {
  ({ doc } = getAPI([inputFile]));
} catch (err) {
  console.error(`Failed to read or parse input file "${inputFile}": ${err.message}`);
  process.exit(1);
}

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

try {
  saveAPI(doc, outputFile);
} catch (err) {
  console.error(`Failed to write output file "${outputFile}": ${err.message}`);
  process.exit(1);
}

process.exit(0);
