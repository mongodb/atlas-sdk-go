const { getAPI, saveAPI } = require("./engine/apifile");
const log = require("simple-node-logger").createSimpleLogger();
const { runAllTransformations } = require("./atlasTransformations");

// Override default logger
global.console = log;
log.setLevel(process.env.XGEN_LOGGING_LEVEL || "warn");

let { doc, apiFileLocation } = getAPI(process.argv.slice(2));

doc = runAllTransformations(doc);
saveAPI(doc, apiFileLocation);
