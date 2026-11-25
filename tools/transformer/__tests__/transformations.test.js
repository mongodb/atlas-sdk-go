const fs = require("fs");
const { test, beforeEach, expect, jest: jestGlobal } = require("@jest/globals");
const {
  applyModelNameTransformations,
  transformAllOf,
  transformOneOf,
  applyAnyOfTransformations,
  applyDiscriminatorTransformations,
  applyOneOfTransformations,
  applyOperationIdOverrides,
  applyServersValidation,
} = require("../src/transformations");
const cases = require("./transformations-snapshots");
const {
  cloudProviderAccessFeatureUsageTransformation,
} = require("../src/atlasTransformations");

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

test("Transform nested oneOf preserves inline allOf properties", () => {
  transformOneOf(".components.schemas.StreamsDataCatalogConnection", api);
  transformOneOf(".components.schemas.StreamsConnection", api);
  expect(
    api.components.schemas.StreamsConnection.properties,
  ).toMatchInlineSnapshot(cases.NestedOneOfInlineAllOfProperties);
});

test("Transform nested anyOf preserves inline allOf properties", () => {
  applyAnyOfTransformations(api);
  expect(
    api.components.schemas.AnyOfConnection.properties,
  ).toMatchInlineSnapshot(cases.NestedAnyOfInlineAllOfProperties);
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

describe("applyServersValidation", () => {
  const singleRootServer = {
    servers: [{ url: "https://cloud.mongodb.com" }],
    paths: {
      "/api/atlas/v2/groups": { get: { operationId: "listGroups" } },
    },
  };

  test("passes with exactly one root server and no operation-level servers", () => {
    expect(applyServersValidation(singleRootServer)).toBe(singleRootServer);
  });

  test("throws when root servers are missing", () => {
    expect(() => applyServersValidation({ paths: {} })).toThrow(
      /exactly one root server/,
    );
  });

  test("throws when multiple root servers are defined", () => {
    const multipleServers = {
      servers: [
        { url: "https://cloud.mongodb.com" },
        { url: "https://api.mongodb.com" },
      ],
      paths: {},
    };
    expect(() => applyServersValidation(multipleServers)).toThrow(
      /exactly one root server/,
    );
  });

  test("throws when a path-level servers block is defined", () => {
    const pathServers = {
      servers: [{ url: "https://cloud.mongodb.com" }],
      paths: {
        "/api/atlas/v2/groups": {
          servers: [{ url: "https://api.mongodb.com" }],
          get: { operationId: "listGroups" },
        },
      },
    };
    expect(() => applyServersValidation(pathServers)).toThrow(
      /servers block at path level/,
    );
  });

  test("throws when an operation-level servers block is defined", () => {
    const operationServers = {
      servers: [{ url: "https://cloud.mongodb.com" }],
      paths: {
        "/api/atlas/v2/groups": {
          get: {
            operationId: "listGroups",
            servers: [{ url: "https://api.mongodb.com" }],
          },
        },
      },
    };
    expect(() => applyServersValidation(operationServers)).toThrow(
      /servers block at operation level/,
    );
  });
});

describe("cloudProviderAccessFeatureUsageTransformation", () => {
  function variant(featureIdRef) {
    return {
      allOf: [
        { $ref: "#/components/schemas/CloudProviderAccessFeatureUsage" },
        {
          properties: {
            featureId: { $ref: "#/components/schemas/" + featureIdRef },
          },
          type: "object",
        },
      ],
    };
  }

  const buildApi = () => ({
    components: {
      schemas: {
        CloudProviderAccessFeatureUsage: {
          discriminator: {
            mapping: {
              ATLAS_DATA_LAKE:
                "#/components/schemas/CloudProviderAccessDataLakeFeatureUsage",
              ENCRYPTION_AT_REST:
                "#/components/schemas/CloudProviderAccessEncryptionAtRestFeatureUsage",
              PUSH_BASED_LOG_EXPORT:
                "#/components/schemas/CloudProviderAccessPushBasedLogExportFeatureUsage",
            },
            propertyName: "featureType",
          },
          properties: { featureType: { type: "string" } },
          type: "object",
        },
        CloudProviderAccessDataLakeFeatureUsage: variant(
          "CloudProviderAccessFeatureUsageDataLakeFeatureId",
        ),
        CloudProviderAccessEncryptionAtRestFeatureUsage: variant(
          "ApiAtlasCloudProviderAccessFeatureUsageFeatureIdView",
        ),
        CloudProviderAccessPushBasedLogExportFeatureUsage: variant(
          "CloudProviderAccessFeatureUsagePushBasedLogExportFeatureId",
        ),
        CloudProviderAccessFeatureUsageDataLakeFeatureId: {
          properties: {
            groupId: { type: "string" },
            name: { type: "string" },
          },
          type: "object",
        },
        CloudProviderAccessFeatureUsagePushBasedLogExportFeatureId: {
          properties: {
            bucketName: { type: "string" },
            groupId: { type: "string" },
          },
          type: "object",
        },
        // The ENCRYPTION_AT_REST variant points at an empty schema.
        ApiAtlasCloudProviderAccessFeatureUsageFeatureIdView: {
          nullable: true,
          type: "object",
        },
      },
    },
  });

  test("collects every variant featureId into a synthetic oneOf model", () => {
    const api = buildApi();
    cloudProviderAccessFeatureUsageTransformation(api);

    expect(
      api.components.schemas.CloudProviderAccessFeatureUsage.properties
        .featureId,
    ).toEqual({
      $ref: "#/components/schemas/BaseCloudProviderAccessFeatureUsageFeatureId",
    });
    expect(
      api.components.schemas.BaseCloudProviderAccessFeatureUsageFeatureId,
    ).toEqual({
      oneOf: [
        {
          $ref: "#/components/schemas/CloudProviderAccessFeatureUsageDataLakeFeatureId",
        },
        {
          $ref: "#/components/schemas/CloudProviderAccessFeatureUsagePushBasedLogExportFeatureId",
        },
      ],
    });
  });

  test("strips featureId from the variant that is not collected", () => {
    const api = buildApi();
    cloudProviderAccessFeatureUsageTransformation(api);

    // An empty schema cannot be a oneOf entry, but the property still has to go.
    // Leaving it behind would overwrite the synthetic ref during flattening.
    expect(
      api.components.schemas.CloudProviderAccessEncryptionAtRestFeatureUsage
        .allOf[1].properties.featureId,
    ).toBeUndefined();
  });

  test("is a no-op when the parent schema is absent", () => {
    const api = { components: { schemas: {} } };
    expect(() =>
      cloudProviderAccessFeatureUsageTransformation(api),
    ).not.toThrow();
    expect(api.components.schemas).toEqual({});
  });

  test("keeps every variant field through the discriminator and oneOf passes", () => {
    const api = buildApi();
    cloudProviderAccessFeatureUsageTransformation(api);
    applyDiscriminatorTransformations(api);
    applyOneOfTransformations(api);

    // The parent has to survive flattening still pointing at the synthetic model.
    // If a child kept its own featureId it would overwrite this ref, which is the
    // bug the property stripping exists to prevent.
    expect(
      api.components.schemas.CloudProviderAccessFeatureUsage.properties
        .featureId,
    ).toEqual({
      $ref: "#/components/schemas/BaseCloudProviderAccessFeatureUsageFeatureId",
    });

    // The synthetic model is flattened in place, so it ends up holding the fields
    // of every variant rather than only the last one.
    expect(
      Object.keys(
        api.components.schemas.BaseCloudProviderAccessFeatureUsageFeatureId
          .properties,
      ).sort(),
    ).toEqual(["bucketName", "groupId", "name"]);
  });
});
