const stableOperationIds = require("../operations.stable.json").stableIds;

/**
 * Transforms the path methods to mark them as "experimental" where paths are not flagged as stable.
 *
 * @param {*} api OpenAPI JSON File
 * @returns OpenAPI JSON File
 */
function applyAddExperimentalTag(api) {
  Object.values(api.paths).forEach((path) => {
    Object.values(path).forEach((requestMethod) => {
      if (
        requestMethod.operationId &&
        !stableOperationIds.includes(requestMethod.operationId)
      ) {
        requestMethod.description = `${requestMethod.description}`;
        requestMethod["x-xgen-experimental"] = true;
      }
    });
  });
  return api;
}

module.exports = {
  applyAddExperimentalTag,
};
