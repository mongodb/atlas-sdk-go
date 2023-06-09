const {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations,
  applyDiscriminatorTransformations,
  applyArrayTransformations,
  transformOneOfProperties,
  applyRemoveEnumsTransformations,
  applyRemoveObjectAdditonalProperties,
} = require("./transformations");

const removeUnusedSchemas = require("./engine/removeUnused");
const { getObjectFromYamlPath } = require("./engine/readers");

const ignoredModelNames = require("./name.ignore.json").ignoreModels;

/**
 * Function specifies list of transformations to run
 */
module.exports = function runTransformations(openapi) {
  // Patching till upstream change will be merged
  // Will be applied upstream as well
  addOneOfTransform(openapi, [
    ".components.schemas.NotificationViewForNdsGroup",
    ".components.schemas.ApiAtlasServerlessTenantEndpointView",
    ".components.schemas.ApiAtlasEndpointServiceView",
    ".components.schemas.ApiAtlasFTSAnalyzersViewManual.properties.charFilters.items",
    ".components.schemas.ApiAtlasFTSAnalyzersViewManual.properties.tokenizer",
  ]);

  openapi = applyDiscriminatorTransformations(openapi);
  openapi = applyOneOfTransformations(openapi);
  openapi = applyAllOfTransformations(openapi);

  openapi = applyModelNameTransformations(
    openapi,
    "Api",
    "",
    ignoredModelNames
  );
  openapi = applyModelNameTransformations(
    openapi,
    "Atlas",
    "",
    ignoredModelNames
  );
  openapi = applyModelNameTransformations(
    openapi,
    "",
    "Manual",
    ignoredModelNames
  );
  openapi = applyModelNameTransformations(
    openapi,
    "",
    "View",
    ignoredModelNames
  );

  // Temp workaround for
  // https://jira.mongodb.org/browse/CLOUDP-166120
  openapi.components.responses.noBody = {
    content: {
      "application/vnd.atlas.2023-01-01+json": {
        example: "",
      },
    },
    description: "This endpoint does not return a response body",
  };

  // Temp workaround for CLOUDP-168427
  if (openapi.components.schemas.Error) {
    openapi.components.schemas.Error.properties.parameters.items = {};
  }

  // Temp workaround for CLOUDP-170462
  openapi = applyArrayTransformations(openapi, [
    "EventTypeForNdsGroup",
    "EventTypeForOrg",
  ]);

  applyRemoveEnumsTransformations(openapi);
  applyRemoveObjectAdditonalProperties(openapi);

  // Required for RegionConfig
  workaroundNestedTransformations(openapi);

  let hasSchemaChanges = true;
  // Remove referencing objects that become unused
  while (hasSchemaChanges) {
    console.info("Checking for unused schemas");
    hasSchemaChanges = removeUnusedSchemas(openapi);
  }

  return openapi;
};

function workaroundNestedTransformations(openapi) {
  try {
    const parentObject = getObjectFromYamlPath(
      ".components.schemas.RegionConfig",
      openapi
    );
    if (parentObject) {
      transformOneOfProperties(parentObject, openapi);
    }
  } catch (e) {
    throw new Error("ApiAtlasRegionConfigView cannot be renamed" + e);
  }
}

// Patch  "x-xgen-go-transform": "merge-oneOf"
function addOneOfTransform(openapi, objectNames) {
  objectNames.forEach((name) => {
    schemObj = getObjectFromYamlPath(name, openapi);
    if (schemObj) {
      schemObj["x-xgen-go-transform"] = "merge-oneOf";
    } else {
      console.warn("Missing object to add x-xgen-go-transform", name);
    }
  });
}
