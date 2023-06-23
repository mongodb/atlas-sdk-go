/**
 * Remove all AdditonalProperties object section
 * That generates an map[string]map[string]string objects for SDK
 * but for our API much safer is to rely on the map[string]interface{}
 * This is due to the fact that dynamic field can be array which would
 * not be possible to represent in map[string]map[string]string structure.
 * @param {*} api OpenAPI JSON File
 * @param modelNames
 * @returns OpenAPI JSON File
 */
function applyRemoveObjectAdditonalProperties(api) {
  const hasSchemas = api && api.components && api.components.schemas;
  if (!hasSchemas) {
    throw new Error("Missing schemas in openapi");
  }
  // Recursive function to traverse the OpenAPI object
  function removeObjectAdditonalProperties(obj) {
    if (typeof obj !== "object" || obj === null) {
      return;
    }

    if (Array.isArray(obj)) {
      for (let i = 0; i < obj.length; i++) {
        removeObjectAdditonalProperties(obj[i]);
      }
    } else {
      // Remove enum field if present
      if (obj.hasOwnProperty("additionalProperties")) {
        if (obj.additionalProperties.type === "object") {
          delete obj.additionalProperties.type;
        }
      }

      // Traverse nested properties
      for (const prop in obj) {
        if (obj.hasOwnProperty(prop)) {
          removeObjectAdditonalProperties(obj[prop]);
        }
      }
    }
  }

  // Start removing enum fields from the OpenAPI object
  removeObjectAdditonalProperties(api);
  return api;
}

module.exports = {
  applyRemoveObjectAdditonalProperties,
};
