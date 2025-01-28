const {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations,
  applyDiscriminatorTransformations,
  applyRemoveEnumsTransformations,
  applyRemoveObjectAdditionalProperties,
  applyAnyOfTransformations,
  applyRemoveNullableTransformations,
  removeRefsFromParameters,
  reorderResponseBodies,
  applyFieldTransformations,
} = require("./transformations");
const { resolveOpenAPIReference } = require("./engine/transformers");

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
  openapi = applyRemoveObjectAdditionalProperties(openapi);
  openapi = reorderResponseBodies(openapi);
  openapi = applyFieldTransformations(openapi);

  openapi = applyModelNameTransformations(
    openapi,
    "",
    "Manual",
    ignoredModelNames,
  );
  openapi = applyModelNameTransformations(
    openapi,
    "",
    "View",
    ignoredModelNames,
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

// Temporary transformation until new search version is introduced.
function searchAPIIssuesTransformation(openapi) {
  const modelsToFix = [
    {
      modelObject: openapi.components.schemas.SearchIndexResponse,
      property: "latestDefinition",
      // Default class model are not correct
      newModelName: "BaseSearchIndexResponseLatestDefinition",
    },
    {
      modelObject: openapi.components.schemas.SearchIndexCreateRequest,
      property: "definition",
      // Default class model are not correct
      newModelName: "BaseSearchIndexCreateRequestDefinition",
    },
  ];

  for (const model of modelsToFix) {
    const responseParent = model.modelObject;
    if (!responseParent) {
      console.error("SearchTransformation: Model object not found:", model);
      continue; // Skip to the next model if not found.
    }

    if (responseParent.discriminator && responseParent.discriminator.mapping) {
      const transformedModel = (openapi.components.schemas[model.newModelName] =
      {
        oneOf: [],
      });
      responseParent.properties[model.property] = {
        $ref: "#/components/schemas/" + model.newModelName,
      };

      for (const mappingKey in responseParent.discriminator.mapping) {
        const ref = responseParent.discriminator.mapping[mappingKey];
        if (!ref) {
          continue; // Skip if there's no reference
        }
        const reference = resolveOpenAPIReference(openapi, ref);
        if (
          reference &&
          reference.allOf &&
          reference.allOf[1] &&
          reference.allOf[1].properties &&
          reference.allOf[1].properties[model.property]
        ) {
          transformedModel.oneOf.push({
            $ref: reference.allOf[1].properties[model.property].$ref,
          });
          delete reference.allOf[1].properties[model.property];
        }
      }
    }
    console.debug("SearchTransformation:", responseParent);
  }
  return openapi;
}
