const stableOperationIds = require("../operations.stable.json").stableIds;

/**
 * Transforms the path description to mark paths as "experimental" where paths are not flagged as stable.
 * Transformation prepends the tag "(experimental)" at the start of the path description
 *
 *
 * @param {*} api OpenAPI JSON File
 * @returns OpenAPI JSON File
 */
function applyAddExperimentalToDescriptions(api) {
  Object.values(api.paths).forEach((path) => {
    Object.values(path).forEach((requestMethod) => {
      if (
        requestMethod.operationId &&
        !stableOperationIds.includes(requestMethod.operationId)
      ) {
        requestMethod.description = `(experimental) ${requestMethod.description}`;
      }
    });
  });
  return api;
}

module.exports = {
  applyAddExperimentalToDescriptions,
};
