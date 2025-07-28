const OP_ID_OVERRIDE_EXTENSION = "x-xgen-operation-id-override";

/**
 * Replaces Operation IDs if operation ID overrides are present and removes the
 * override extension.
 *
 * @param {*} api OpenAPI JSON File
 * @returns {*} transformed OpenAPI JSON File
 */
function applyOperationIdOverrides(api) {
  const hasPaths = api && api.paths;
  if (!hasPaths) {
    throw new Error("Missing paths in openapi");
  }

  Object.keys(api.paths).forEach((pathKey) => {
    const pathItem = api.paths[pathKey];
    if (!pathItem) {
      throw new Error("Missing path item in openapi");
    }

    Object.keys(pathItem).forEach((operationKey) => {
      const operationItem = api.paths[pathKey][operationKey];
      if (!operationItem) {
        throw new Error("Missing operation item in openapi");
      }

      // Update Operation ID based on override extension
      if (operationItem[OP_ID_OVERRIDE_EXTENSION]) {
        api.paths[pathKey][operationKey].operationId =
          operationItem[OP_ID_OVERRIDE_EXTENSION];
        delete api.paths[pathKey][operationKey][OP_ID_OVERRIDE_EXTENSION];
      }
    });
  });

  return api;
}

module.exports = {
  applyOperationIdOverrides,
};
