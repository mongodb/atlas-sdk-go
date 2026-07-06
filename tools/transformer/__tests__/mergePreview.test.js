const { test, expect } = require("@jest/globals");
const { mergePreview } = require("../src/engine/mergePreview");

function baseDoc() {
  return {
    paths: {
      "/clusters": {
        get: { operationId: "listClusters", tags: ["Clusters"] },
        post: { operationId: "createCluster", tags: ["Clusters"] },
      },
      "/projects": {
        get: { operationId: "listProjects", tags: ["Projects"] },
      },
    },
    components: {
      schemas: {
        Cluster: { type: "object", properties: { id: { type: "string" } } },
      },
    },
    tags: [{ name: "Clusters" }, { name: "Projects" }],
  };
}

test("same path+method: preview overrides the stable operation", () => {
  const preview = {
    paths: {
      "/clusters": {
        get: {
          operationId: "listClusters",
          tags: ["Clusters"],
          responses: {
            200: {
              content: { "application/vnd.atlas.preview+json": {} },
            },
          },
        },
      },
    },
  };

  const merged = mergePreview(baseDoc(), preview);

  expect(
    Object.keys(merged.paths["/clusters"].get.responses["200"].content),
  ).toContain("application/vnd.atlas.preview+json");
  // Sibling method on the same path is preserved.
  expect(merged.paths["/clusters"].post.operationId).toEqual("createCluster");
});

test("same path, method only in preview: both methods kept", () => {
  const preview = {
    paths: {
      "/clusters": {
        delete: { operationId: "deleteClusterPreview", tags: ["Clusters"] },
      },
    },
  };

  const merged = mergePreview(baseDoc(), preview);

  expect(merged.paths["/clusters"].get.operationId).toEqual("listClusters");
  expect(merged.paths["/clusters"].post.operationId).toEqual("createCluster");
  expect(merged.paths["/clusters"].delete.operationId).toEqual(
    "deleteClusterPreview",
  );
});

test("new preview path is added as-is", () => {
  const preview = {
    paths: {
      "/streams": {
        get: { operationId: "listStreamsPreview", tags: ["Streams"] },
      },
    },
  };

  const merged = mergePreview(baseDoc(), preview);

  expect(merged.paths["/streams"].get.operationId).toEqual(
    "listStreamsPreview",
  );
  // Existing paths untouched.
  expect(merged.paths["/projects"].get.operationId).toEqual("listProjects");
});

test("components are merged and preview entries override matching ones", () => {
  const preview = {
    components: {
      schemas: {
        // Overrides existing base schema.
        Cluster: {
          type: "object",
          properties: { previewField: { type: "string" } },
        },
        // New schema referenced by preview operations.
        StreamPreview: { type: "object" },
      },
      parameters: {
        previewParam: { name: "preview", in: "query" },
      },
    },
  };

  const merged = mergePreview(baseDoc(), preview);

  expect(merged.components.schemas.Cluster.properties).toHaveProperty(
    "previewField",
  );
  expect(merged.components.schemas).toHaveProperty("StreamPreview");
  expect(merged.components.parameters).toHaveProperty("previewParam");
});

test("preview tags are appended without duplicating existing ones", () => {
  const preview = {
    tags: [{ name: "Clusters" }, { name: "Streams" }],
  };

  const merged = mergePreview(baseDoc(), preview);

  const names = merged.tags.map((t) => t.name);
  expect(names).toEqual(["Clusters", "Projects", "Streams"]);
});

test("missing preview doc returns the base unchanged", () => {
  const base = baseDoc();
  expect(mergePreview(base, undefined)).toBe(base);
});
