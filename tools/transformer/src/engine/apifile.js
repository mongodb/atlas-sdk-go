const { parse, stringify } = require("yaml");
const { readFileSync, writeFileSync } = require("fs");

/**
 * Read and parse API file as json
 *
 * @returns  {doc, apiFileLocation}
 */
function getAPI(apiFileLocation) {
  if (!apiFileLocation && !apiFileLocation[0]) {
    throw new Error("Missing positional argument for openapi file");
  }

  apiFileLocation = apiFileLocation[0];

  let doc;
  if (apiFileLocation) {
    doc = parse(readFileSync(apiFileLocation, "utf8"));
  } else {
    throw new Error("Missing location. Please set OPENAPI_FILE env variable");
  }
  return { doc, apiFileLocation };
}

/**
 * Save API to target location
 *
 * @param {*} doc openapi doc
 * @param {*} apiFileLocation location to save
 */
function saveAPI(doc, apiFileLocation) {
  writeFileSync(apiFileLocation, stringify(doc, { lineWidth: 9999 })); // default is 80 and would add newlines, we want to stay close to the original that uses long lines.
}

function readFromStdin() {
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

function writeToStdout(doc) {
  process.stdout.write(stringify(doc, { lineWidth: 9999 }));
}

module.exports = { getAPI, saveAPI, readFromStdin, writeToStdout };
