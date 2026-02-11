const fs = require("fs");
const { test, beforeEach, expect, jest: jestGlobal } = require("@jest/globals");
const {
  applyModelNameTransformations,
  transformAllOf,
  transformOneOf,
  applyOperationIdOverrides,
} = require("../src/transformations");
const cases = require("./transformations-snapshots");

let api = {};
beforeEach(() => {
  api = JSON.parse(fs.readFileSync(__dirname + "/input.json", "utf8"));
});

test("Transform oneOf enum", () => {
  transformOneOf(
    ".components.schemas.ApiAtlasRegionConfigView.properties.regionName",
    api,
  );
  expect(
    api.components.schemas.ApiAtlasRegionConfigView.properties.regionName,
  ).toMatchInlineSnapshot(cases.EnumOneOf);
});

test("Transform oneOf model", () => {
  transformOneOf(".components.schemas.ApiAtlasHardwareSpecView", api);
  expect(
    api.components.schemas.ApiAtlasHardwareSpecView.properties,
  ).toMatchInlineSnapshot(cases.PropertiesOneOf);
});

test("Transform oneOf model with discriminator emits x-xgen-discriminator extension", () => {
  transformOneOf(".components.schemas.ApiAtlasRegionConfigView", api);
  const result = api.components.schemas.ApiAtlasRegionConfigView;

  // Extension is present and discriminator/oneOf are removed
  expect(result["x-xgen-discriminator"]).toBeDefined();
  expect(result.discriminator).toBeUndefined();
  expect(result.oneOf).toBeUndefined();

  // Snapshot the full extension
  expect(result["x-xgen-discriminator"]).toMatchInlineSnapshot(
    cases.DiscriminatorExtension,
  );
});

test("Transform AllOf model", () => {
  transformAllOf(".components.schemas.ApiAtlasRegionConfigView", api);
  expect(api.components.schemas.ApiAtlasRegionConfigView).toMatchInlineSnapshot(
    cases.ParentAllOf,
  );
  expect(
    api.components.schemas.ApiAtlasAWSRegionConfigView,
  ).toMatchInlineSnapshot(cases.ChildAllOf);
});

test("Transform already transformed model ", () => {
  // First transform
  transformAllOf(".components.schemas.ApiAtlasRegionConfigView", api);
  // Second transform with no effect
  transformAllOf(".components.schemas.ApiAtlasRegionConfigView", api);

  expect(api.components.schemas.ApiAtlasRegionConfigView).toMatchInlineSnapshot(
    cases.ParentAllOf,
  );
  expect(
    api.components.schemas.ApiAtlasAWSRegionConfigView,
  ).toMatchInlineSnapshot(cases.ChildAllOf);
});

test("applyModelNameTransformations", () => {
  api = applyModelNameTransformations(api, "ApiAtlas", "View");
  api = applyModelNameTransformations(api, "Api", "View");
  for (const modelKey in api.components.schemas) {
    expect(modelKey.startsWith("ApiAtlas")).toBeFalsy();
    expect(modelKey.startsWith("Api")).toBeFalsy();
    expect(modelKey.endsWith("View")).toBeFalsy();
  }
});

test("applyOperationIdOverrides", () => {
  api = applyOperationIdOverrides(api);
  expect(
    api.paths["/api/atlas/v1.5/groups/{groupId}/clusters"].post.operationId,
  ).toEqual("createCluster");
  expect(api.paths["/api/atlas/v2/example/info"].get.operationId).toEqual(
    "getVersionedExample",
  );
  expect(
    api.paths["/api/atlas/v2/example/info"].get["x-xgen-operation-id-override"],
  ).toBeFalsy();
});
