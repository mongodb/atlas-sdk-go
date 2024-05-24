const {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations,
  applyDiscriminatorTransformations,
  applyRemoveEnumsTransformations,
  applyRemoveObjectAdditonalProperties,
  applyAnyOfTransformations,
  applyRemoveNullableTransformations,
  removeRefsFromParameters,
} = require("./transformations");

const removeUnusedSchemas = require("./engine/removeUnused");

const ignoredModelNames = require("./name.ignore.json").ignoreModels;

/**
 * Function specifies list of transformations to run
 */
module.exports = function runTransformations(openapi) {
  openapi = applyDiscriminatorTransformations(openapi);
  openapi = applyOneOfTransformations(openapi);
  openapi = applyAnyOfTransformations(openapi);
  openapi = applyAllOfTransformations(openapi);
  openapi = applyRemoveEnumsTransformations(openapi);
  openapi = applyRemoveNullableTransformations(openapi);
  openapi = applyRemoveObjectAdditonalProperties(openapi);

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

  if (openapi.components.schemas.ApiAtlasFTSAnalyzers) {
    filtersObj = openapi.components.schemas.ApiAtlasFTSAnalyzers;
    if (filtersObj.properties.tokenFilters) {
      filtersObj.properties.tokenFilters.items = {};
    }
    if (filtersObj.properties.charFilters) {
      filtersObj.properties.charFilters.items = {};
    }
  }

  let hasSchemaChanges = true;
  // Remove referencing objects that become unused
  while (hasSchemaChanges) {
    console.info("Checking for unused schemas");
    hasSchemaChanges = removeUnusedSchemas(openapi);
  }

  if (openapi.components.schemas.ApiError) {
    openapi.components.schemas.ApiError.properties.parameters.items = {};
  }

  const result = removeRefsFromParameters(openapi, [
    "#/components/parameters/envelope",
    "#/components/parameters/pretty",
  ]);

  return openapi;
};
