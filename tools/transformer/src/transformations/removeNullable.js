const removeField = require("../engine/removeField");

/**
 * Remove nullable from the schema.
 * @param {*} api OpenAPI JSON File
 * @param modelNames
 * @returns OpenAPI JSON File
 */
function applyRemoveNullableTransformations(api) {
  return removeField(api, "nullable");
}

module.exports = {
  applyRemoveNullableTransformations,
};
