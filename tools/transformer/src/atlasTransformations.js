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
    "NotificationViewForNdsGroup",
    "ApiAtlasServerlessTenantEndpointView",
    "ApiAtlasEndpointServiceView",
  ]);

  addOneOfTransformNested(openapi,[
    "ApiAtlasFTSAnalyzersViewManual.properties.tokenizer",
    "ApiAtlasFTSAnalyzersViewManual.properties.charFilters.items"
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
    if (openapi.components.schemas[name]) {
      openapi.components.schemas[name]["x-xgen-go-transform"] = "merge-oneOf";
    } else {
      console.warn("Missing object to add x-xgen-go-transform" + object);
    }
  });
}

// Patch  "x-xgen-go-transform": "merge-oneOf" for nested objects
function addOneOfTransformNested(openapi, nestedObjectNames) {
  nestedObjectNames.forEach((nestedObject) => {
    const keys = nestedObject.split('.');
    let currentObj = openapi.components.schemas;
    for (let i = 0; i < keys.length; i++) {
      const key = keys[i];
      if (currentObj.hasOwnProperty(key)) {
        currentObj = currentObj[key];
        if (i === keys.length - 1) {
          currentObj["x-xgen-go-transform"] = "merge-oneOf";
        }
      }
      else {
        console.warn("Missing object to add x-xgen-go-transform" + nestedObject);
        break;
      }
    }
  });
}
