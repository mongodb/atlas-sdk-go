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

const {
  resolveOpenAPIReference
} = require("./engine/transformers");

const removeUnusedSchemas = require("./engine/removeUnused");

const ignoredModelNames = require("./name.ignore.json").ignoreModels;

/**
 * Function specifies list of transformations to run
 */
module.exports = function runTransformations(openapi) {
  openapi = searchAPIIssuesTransformation(openapi);

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

// Workarounds for issues in the search API
function searchAPIIssuesTransformation(openapi) {
  // API is overly complex in this case and provides no value as typed interface
  if (openapi.components.schemas.SearchIndexResponse) {
    const responseParent = openapi.components.schemas.SearchIndexResponse;
    if (responseParent.properties && responseParent.properties.latestDefinition) {
      if (responseParent.discriminator && responseParent.discriminator.mapping) {
        responseParent.properties.latestDefinition = { oneOf: [] };
        for (const mappingKey in responseParent.discriminator.mapping) {
          const ref = responseParent.discriminator.mapping[mappingKey];
          if (!ref) {
            continue; // Skip if there's no reference
          }
          const reference = resolveOpenAPIReference(openapi, ref);
          responseParent.properties.latestDefinition.oneOf.push({
            $ref: reference.allOf && reference.allOf[1] && reference.allOf[1].properties && reference.allOf[1].properties.latestDefinition && reference.allOf[1].properties.latestDefinition.$ref
          });
          delete reference.allOf[1]?.properties?.latestDefinition;
        }
      }
    }
  }
  console.debug("SearchIndexResponse", openapi.components?.schemas?.SearchIndexResponse);
  return openapi;
}

