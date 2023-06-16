const {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations,
  applyDiscriminatorTransformations,
  transformOneOfProperties,
  applyRemoveEnumsTransformations,
  applyRemoveObjectAdditonalProperties,
  applyAddExperimentalTag,
} = require("./transformations");

const removeUnusedSchemas = require("./engine/removeUnused");
const { getObjectFromYamlPath } = require("./engine/readers");

const ignoredModelNames = require("./name.ignore.json").ignoreModels;
const stableOperationIds = require("./operations.stable.json").stableIds;

/**
 * Function specifies list of transformations to run
 */
module.exports = function runTransformations(openapi) {
  // Patching till upstream change will be merged
  // Will be applied upstream as well
  addOneOfTransform(openapi, [
    // To be removed after CLOUDP-170462 is available upstream
    ".components.schemas.NotificationViewForNdsGroup",
    ".components.schemas.ApiAtlasServerlessTenantEndpointView",
    ".components.schemas.ApiAtlasEndpointServiceView",

    // Renamed objects
    ".components.schemas.AlertNotificationViewGroup",
    ".components.schemas.ServerlessTenantEndpoint",
    ".components.schemas.EndpointService",

    ".components.schemas.ApiAtlasFTSAnalyzersViewManual.properties.charFilters.items",
    ".components.schemas.ApiAtlasFTSAnalyzersViewManual.properties.tokenizer",
  ]);

  openapi = applyAddExperimentalTag(openapi, stableOperationIds);
  openapi = applyDiscriminatorTransformations(openapi);
  openapi = applyOneOfTransformations(openapi);
  openapi = applyAllOfTransformations(openapi);

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

  if (openapi.components.schemas.ApiError) {
    openapi.components.schemas.ApiError.properties.parameters.items = {};
  }

  if (openapi.components.schemas.ApiAtlasFTSAnalyzers) {
    filtersObj = openapi.components.schemas.ApiAtlasFTSAnalyzers;
    if (filtersObj.properties.tokenFilters) {
      filtersObj.properties.tokenFilters.items = {};
    }
    if (filtersObj.properties.charFilters) {
      filtersObj.properties.charFilters.items = {};
    }
  }

  applyRemoveEnumsTransformations(openapi);
  applyRemoveObjectAdditonalProperties(openapi);

  // Required for RegionConfig
  workaroundNestedTransformations(openapi);
  // Required for StreamsTenant
  workaroundReadOnly(openapi);

  let hasSchemaChanges = true;
  // Remove referencing objects that become unused
  while (hasSchemaChanges) {
    console.info("Checking for unused schemas");
    hasSchemaChanges = removeUnusedSchemas(openapi);
  }

  return openapi;
};

// Schema contains both read and write only.
function workaroundReadOnly(openapi) {
  const tenantObj = openapi.components.schemas.StreamsTenant;
  delete tenantObj?.properties?.connections?.readOnly;
  delete tenantObj?.properties?.connections?.writeOnly;
}

function workaroundNestedTransformations(openapi) {
  let parentObject;
  try {
    parentObject = getObjectFromYamlPath(
      ".components.schemas.RegionConfig",
      openapi
    );
  } catch (e) {}
  try {
    parentObject =
      parentObject ||
      getObjectFromYamlPath(".components.schemas.CloudRegionConfig", openapi);
  } catch (e) {}
  if (parentObject) {
    transformOneOfProperties(parentObject, openapi);
  } else {
    throw new Error("RegionConfig cannot be renamed" + e);
  }
}

// Patch  "x-xgen-go-transform": "merge-oneOf"
function addOneOfTransform(openapi, objectNames) {
  objectNames.forEach((name) => {
    try {
      schemaObj = getObjectFromYamlPath(name, openapi);
      if (schemaObj) {
        schemaObj["x-xgen-go-transform"] = "merge-oneOf";
      } else {
        console.warn("Missing object to add x-xgen-go-transform", name);
      }
    } catch (e) {
      console.warn("Missing object to add x-xgen-go-transform", name);
    }
  });
}
