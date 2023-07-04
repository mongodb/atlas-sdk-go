const removeField = require("../engine/removeField");

/**
 * Remove all enums from the schema.
 * Enums would not allow us to ensure API contract and will introduce breaking changes.
 * Enums should be defined as metadata in the schema.
 * @param {*} api OpenAPI JSON File
 * @param modelNames
 * @returns OpenAPI JSON File
 */
function applyRemoveEnumsTransformations(api) {
  return removeField(api, "enum");
}

module.exports = {
  applyRemoveEnumsTransformations,
};
