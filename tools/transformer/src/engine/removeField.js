/**
 * Removes field from openapi
 * @param {*} api
 * @param {*} fieldName field to remove
 * @returns
 */
module.exports = function removeField(api, fieldName) {
  const hasSchemas = api && api.components && api.components.schemas;
  if (!hasSchemas) {
    throw new Error("Missing schemas in openapi");
  }
  // Recursive function to traverse the OpenAPI object
  function removePropertyField(obj) {
    if (typeof obj !== "object" || obj === null) {
      return;
    }

    if (Array.isArray(obj)) {
      for (let i = 0; i < obj.length; i++) {
        removePropertyField(obj[i], fieldName);
      }
    } else {
      // Remove enum field if present
      if (obj.hasOwnProperty(fieldName)) {
        delete obj[fieldName];
      }

      // Traverse nested properties
      for (const prop in obj) {
        if (obj.hasOwnProperty(prop)) {
          removePropertyField(obj[prop], fieldName);
        }
      }
    }
  }

  // Start removing enum fields from the OpenAPI object
  removePropertyField(api, fieldName);
  return api;
};
