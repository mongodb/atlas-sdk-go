/**
 * Transforms the model name into array of that model name for situations when schema needs to be an array of that model.
 * Transformation acts inline - model object becomes array of that model object and items are organized inline.
 *
 * Temporary workaround for https://jira.mongodb.org/browse/CLOUDP-173481
 *
 * @param {*} api OpenAPI JSON File
 * @param
 * @returns OpenAPI JSON File
 */
function applyArrayTransformations(api, modelNames) {
  const hasSchemas = api && api.components && api.components.schemas;
  if (!hasSchemas) {
    throw new Error("Missing schemas in openapi");
  }
  for (const modelName of modelNames) {
    if (!api.components.schemas[modelName]) {
      // For cases when openapi is already fixed
      console.info(`Missing schema ${modelName} in openapi`);
    } else if (api.components.schemas[modelName].type !== "array") {
      model = api.components.schemas[modelName];
      api.components.schemas[modelName] = {
        type: "array",
        items: {
          ...model,
        },
      };
    }
  }
  return api;
}

module.exports = {
  applyArrayTransformations,
};
