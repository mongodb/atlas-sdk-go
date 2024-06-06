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
    responseParent = openapi.components.schemas.SearchIndexResponse;
    if (responseParent.properties.latestDefinition) {
      if(responseParent.discriminator.mapping){
        responseParent.properties.latestDefinition = {oneOf: []};
        for(mappingKey in responseParent.discriminator.mapping){
          var ref = responseParent.discriminator.mapping[mappingKey];
          var reference = resolveOpenAPIReference(openapi,ref)
          responseParent.properties.latestDefinition.oneOf.push({
            $ref: reference.allOf[1].properties.latestDefinition.$ref
          });
          delete reference.allOf[1].properties.latestDefinition
        }
      }
    }
  }
  console.error("SearchIndexResponse", openapi.components.schemas.SearchIndexResponse)
  console.error("VectorSearchIndexResponse", openapi.components.schemas.VectorSearchIndexResponse)
  console.error("TextSearchIndexResponse", openapi.components.schemas.TextSearchIndexResponse)
  return openapi;
}

function resolveOpenAPIReference(openapi, ref) {
  if (!ref.startsWith("#/")) {
    throw new Error("Invalid reference format: " + ref);
  }
  const parts = ref.split("/");
  // Skip the first empty part
  parts.shift();

  let current = openapi;
  for (const part of parts) {
    if (!current.hasOwnProperty(part)) {
      throw new Error("Reference not found: " + ref);
    }
    current = current[part];
  }
  return current;
}
