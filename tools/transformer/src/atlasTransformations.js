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
  applyOperationIdOverrides,
  applyServersValidation,
} = require("./transformations");
const { resolveOpenAPIReference } = require("./engine/transformers");

const removeUnusedSchemas = require("./engine/removeUnused");

const ignoredModelNames = require("./name.ignore.json").ignoreModels;

/**
 * Function specifies list of transformations to run
 */
function runFlatteningTransformations(openapi) {
  // Handles complex polymorphic type from search API
  openapi = searchAPIIssuesTransformation(openapi);

  // Handles the per-variant featureId of CloudProviderAccessFeatureUsage
  openapi = cloudProviderAccessFeatureUsageTransformation(openapi);

  openapi = applyDiscriminatorTransformations(openapi);
  openapi = applyOneOfTransformations(openapi);
  openapi = applyAnyOfTransformations(openapi);
  openapi = applyAllOfTransformations(openapi);

  // Makes specified fields optional in their respective schemas (relevant for alert configuration schema)
  openapi = applyFieldTransformations(openapi);

  // Fixes invalid schema for search API
  if (openapi.components.schemas.ApiAtlasFTSAnalyzersViewManual) {
    filtersObj = openapi.components.schemas.ApiAtlasFTSAnalyzersViewManual;
    if (filtersObj.properties.tokenFilters) {
      filtersObj.properties.tokenFilters.items = {};
    }
    if (filtersObj.properties.charFilters) {
      filtersObj.properties.charFilters.items = {};
    }
  }

  // Removing unused schemas after flattening polymorphic types
  let hasSchemaChanges = true;
  while (hasSchemaChanges) {
    console.info("Checking for unused schemas");
    hasSchemaChanges = removeUnusedSchemas(openapi);
  }

  return openapi;
}

function runAllTransformations(openapi) {
  // SDK templates assume a single root server and no operation-level servers
  applyServersValidation(openapi);

  openapi = runFlatteningTransformations(openapi);

  // SDK specific transformations
  openapi = applyRemoveEnumsTransformations(openapi);
  openapi = applyRemoveNullableTransformations(openapi);
  openapi = applyRemoveObjectAdditionalProperties(openapi);
  openapi = reorderResponseBodies(openapi);

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

  if (openapi.components.schemas.ApiError) {
    openapi.components.schemas.ApiError.properties.parameters.items = {};
  }

  const result = removeRefsFromParameters(openapi, [
    "#/components/parameters/envelope",
    "#/components/parameters/pretty",
  ]);

  openapi = applyOperationIdOverrides(openapi);

  return openapi;
}

module.exports = {
  runFlatteningTransformations: runFlatteningTransformations,
  runAllTransformations: runAllTransformations,
  cloudProviderAccessFeatureUsageTransformation:
    cloudProviderAccessFeatureUsageTransformation,
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
    if (!model.modelObject) {
      console.error("SearchTransformation: Model object not found:", model);
      continue; // Skip to the next model if not found.
    }
    extractDiscriminatorChildProperty(
      openapi,
      model.modelObject,
      model.property,
      model.newModelName,
    );
  }
  return openapi;
}

// The children of a discriminated schema can each declare the same property with a
// different type. Flattening the children onto the parent keeps only the last mapping
// entry's type and silently drops the others. Collect the child refs into a synthetic
// oneOf model instead, so the merged model keeps the fields of every variant.
function extractDiscriminatorChildProperty(
  openapi,
  parent,
  property,
  newModelName,
) {
  if (!parent || !parent.discriminator || !parent.discriminator.mapping) {
    return;
  }

  const oneOfRefs = [];
  for (const childRef of Object.values(parent.discriminator.mapping)) {
    if (!childRef) {
      continue; // Skip if there's no reference
    }
    const child = resolveOpenAPIReference(openapi, childRef);
    const holder = (child.allOf || []).find(
      (item) => item.properties && item.properties[property],
    );
    if (!holder) {
      continue;
    }

    const propertyRef = holder.properties[property].$ref;
    // Strip the property from the child either way. Leaving it behind would overwrite
    // the synthetic ref when the children are flattened onto the parent.
    delete holder.properties[property];
    // transformOneOfProperties rejects a oneOf entry with neither properties nor allOf,
    // so an empty schema cannot be collected.
    if (propertyRef && hasPropertiesOrAllOf(openapi, propertyRef)) {
      oneOfRefs.push({ $ref: propertyRef });
    }
  }

  openapi.components.schemas[newModelName] = { oneOf: oneOfRefs };
  parent.properties[property] = {
    $ref: "#/components/schemas/" + newModelName,
  };
}

function hasPropertiesOrAllOf(openapi, ref) {
  const schema = resolveOpenAPIReference(openapi, ref);
  return !!(schema && (schema.properties || schema.allOf));
}

// Each variant of CloudProviderAccessFeatureUsage declares featureId with its own type.
// Without this the parent keeps only PUSH_BASED_LOG_EXPORT's type and the fields of the
// other variants, such as privateEndpointId, are dropped.
function cloudProviderAccessFeatureUsageTransformation(openapi) {
  extractDiscriminatorChildProperty(
    openapi,
    openapi.components.schemas.CloudProviderAccessFeatureUsage,
    "featureId",
    "BaseCloudProviderAccessFeatureUsageFeatureId",
  );
  return openapi;
}
