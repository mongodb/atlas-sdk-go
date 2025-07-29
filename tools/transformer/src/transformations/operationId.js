const OP_ID_OVERRIDE_EXTENSION = "x-xgen-operation-id-override";
const validHttpMethods = [
  "get",
  "put",
  "post",
  "delete",
  "options",
  "head",
  "patch",
  "trace",
];

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

    validHttpMethods.forEach((method) => {
      const operationItem = api.paths[pathKey][method];
      if (operationItem && operationItem[OP_ID_OVERRIDE_EXTENSION]) {
        // Update Operation ID based on override extension
        api.paths[pathKey][method].operationId =
          operationItem[OP_ID_OVERRIDE_EXTENSION];
        delete api.paths[pathKey][method][OP_ID_OVERRIDE_EXTENSION];
      }
    });
  });

  return api;
}

module.exports = {
  applyOperationIdOverrides,
};
