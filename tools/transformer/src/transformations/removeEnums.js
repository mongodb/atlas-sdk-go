/**
 * Remove all enums from the schema.
 * Enums would not allow us to ensure API contract and will introduce breaking changes.
 * Enums should be defined as metadata in the schema.
 * @param {*} api OpenAPI JSON File
 * @param modelNames
 * @returns OpenAPI JSON File
 */
function applyRemoveEnumsTransformations(api) {
  const hasSchemas = api && api.components && api.components.schemas;
  if (!hasSchemas) {
    throw new Error("Missing schemas in openapi");
  }
  // Recursive function to traverse the OpenAPI object
  function removeEnums(obj) {
    if (typeof obj !== "object" || obj === null) {
      return;
    }

    if (Array.isArray(obj)) {
      for (let i = 0; i < obj.length; i++) {
        removeEnums(obj[i]);
      }
    } else {
      // Remove enum field if present
      if (obj.hasOwnProperty("enum")) {
        delete obj.enum;
      }

      // Traverse nested properties
      for (const prop in obj) {
        if (obj.hasOwnProperty(prop)) {
          removeEnums(obj[prop]);
        }
      }
    }
  }

  // Start removing enum fields from the OpenAPI object
  removeEnums(api);
  return api
}

module.exports = {
  applyRemoveEnumsTransformations,
};
