const {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations,
  applyDiscriminatorTransformations,
} = require("./transformations");

const removeUnusedSchemas = require("./engine/removeUnused");

const ignoredModelNames = require("./name.ignore.json").ignoreModels;

/**
 * Function specifies list of transformations to run
 */
module.exports = function runTransformations(openapi) {
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

  let hasSchemaChanges = true;
  // Remove referencing objects that become unused
  while (hasSchemaChanges) {
    console.info("Checking for unused schemas")
    openapi, hasSchemaChanges = removeUnusedSchemas(openapi);
  }
  
  return openapi;
};
