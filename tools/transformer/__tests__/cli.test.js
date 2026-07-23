const fs = require("fs");
const os = require("os");
const path = require("path");
const { execFileSync } = require("child_process");
const { parse } = require("yaml");
const { test, expect } = require("@jest/globals");

const CLI_PATH = path.join(__dirname, "../src/cli.js");

function writeYaml(dir, name, doc) {
  const { stringify } = require("yaml");
  const filePath = path.join(dir, name);
  fs.writeFileSync(filePath, stringify(doc));
  return filePath;
}

test("merge-preview CLI layers multiple --preview flags in order", () => {
  const dir = fs.mkdtempSync(path.join(os.tmpdir(), "merge-preview-cli-"));

  const base = {
    paths: {
      "/clusters": {
        get: { operationId: "listClusters", tags: ["Clusters"] },
      },
    },
    components: { schemas: {} },
    tags: [{ name: "Clusters" }],
  };
  const publicPreview = {
    paths: {
      "/streams": {
        get: { operationId: "listStreamsPreview", tags: ["Streams"] },
      },
    },
  };
  const privatePreview = {
    paths: {
      "/rateLimits": {
        get: { operationId: "listRateLimits", tags: ["Rate Limiting"] },
      },
    },
  };

  const basePath = writeYaml(dir, "base.yaml", base);
  const publicPreviewPath = writeYaml(
    dir,
    "public-preview.yaml",
    publicPreview,
  );
  const privatePreviewPath = writeYaml(
    dir,
    "private-preview.yaml",
    privatePreview,
  );
  const outputPath = path.join(dir, "out.yaml");

  execFileSync(
    process.execPath,
    [
      CLI_PATH,
      "merge-preview",
      "--base",
      basePath,
      "--preview",
      publicPreviewPath,
      "--preview",
      privatePreviewPath,
      "--output",
      outputPath,
    ],
    { encoding: "utf8" },
  );

  const merged = parse(fs.readFileSync(outputPath, "utf8"));

  // Base and both deltas' additions are all present in the output.
  expect(merged.paths["/clusters"].get.operationId).toEqual("listClusters");
  expect(merged.paths["/streams"].get.operationId).toEqual(
    "listStreamsPreview",
  );
  expect(merged.paths["/rateLimits"].get.operationId).toEqual("listRateLimits");

  fs.rmSync(dir, { recursive: true, force: true });
});
