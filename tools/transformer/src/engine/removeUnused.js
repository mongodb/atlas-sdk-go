/**
 * Helper function used to traverse the OpenAPI document and remove unused schemas
 * @param {*} openapiData
 * @returns
 */
function removeUnusedSchemas(openapiData) {
  const allRefs = [];
  const usedRefs = [];

  // Find all $ref occurrences in the OpenAPI document
  findRefs(openapiData, allRefs);

  // Extract unique references used in the OpenAPI document
  allRefs.forEach((ref) => {
    const refParts = ref.split("/");
    if (refParts[1] === "components" && refParts[2] === "schemas") {
      usedRefs.push(refParts[3]);
    }
  });
  let hasSchemaChanges = false;
  // Remove unused schemas from 'components.schemas'
  const schemas = openapiData.components && openapiData.components.schemas;
  if (schemas) {
    Object.keys(schemas).forEach((schemaName) => {
      if (!usedRefs.includes(schemaName)) {
        delete schemas[schemaName];
        hasSchemaChanges = true;
      }
    });
  }
  // Return the updated OpenAPI document
  return hasSchemaChanges;
}

/**
 * Recursive function for finding all $ref occurrences in the OpenAPI document
 * @param {*} obj
 * @param {*} refs
 */
function findRefs(obj, refs) {
  if (typeof obj === "object" && obj !== null) {
    if (Array.isArray(obj)) {
      obj.forEach((item) => findRefs(item, refs));
    } else {
      Object.keys(obj).forEach((key) => {
        if (key === "$ref") {
          refs.push(obj[key]);
        } else {
          findRefs(obj[key], refs);
        }
      });
    }
  }
}

module.exports = removeUnusedSchemas;
