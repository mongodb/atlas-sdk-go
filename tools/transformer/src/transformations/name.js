const isModelIgnored = require("../engine/ignoreModel");

/**
 *  Transforms provided API JSON file by removing prefix and suffix from the
 *  API models
 *
 * @param {*} api OpenAPI JSON File
 * @param {*} prefix model prefix or empty string if only suffix is needed.
 * @param {*} suffix model suffix or empty string if only prefix is needed.
 * @param
 * @returns OpenAPI JSON File
 */
function applyModelNameTransformations(api, prefix, suffix, ignoreModels) {
  const hasSchemas = api && api.components && api.components.schemas;
  if (!hasSchemas) {
    throw new Error("Missing schemas in openapi");
  }

  for (const schemaKey of Object.keys(api.components.schemas)) {
    if (schemaKey.startsWith(prefix) && schemaKey.endsWith(suffix)) {
      if (isModelIgnored(schemaKey, ignoreModels)) {
        console.warn("Ignored rename for the object: " + schemaKey);
        continue;
      }

      const schemaTransformedName = schemaKey
        .replace(prefix, "")
        .replace(suffix, "");
      if (api.components.schemas[schemaTransformedName]) {
        throw new Error(
          `components.schemas already contain key ${schemaTransformedName}. Please rename ${schemaKey} object to avoid name overlap.`,
        );
      }
      api.components.schemas[schemaTransformedName] =
        api.components.schemas[schemaKey];
      delete api.components.schemas[schemaKey];
    }
  }

  let regexp = new RegExp(`"#/components/schemas/${prefix}(.*)${suffix}"`, "g");
  let apiString = JSON.stringify(api, undefined, 2);
  let apiToReplace = apiString.replace(regexp, '"#/components/schemas/$1"');
  api = JSON.parse(apiToReplace);
  return api;
}

module.exports = {
  applyModelNameTransformations,
};
