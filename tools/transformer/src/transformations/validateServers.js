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

const GUIDANCE =
  "atlas-sdk-go decouples the generated default base URL from the spec servers " +
  "block (see tools/config/go-templates/configuration.mustache) and assumes the " +
  "spec defines exactly one root server and no operation-level servers. Before " +
  "changing the spec servers, review that template and internal/core DefaultCloudURL.";

/**
 * Validates the servers assumptions made by the Go SDK templates:
 * exactly one root server and no operation-level servers.
 * Throws an error with guidance if any assumption is broken.
 *
 * @param {*} api OpenAPI JSON File
 * @returns {*} unchanged OpenAPI JSON File
 */
function applyServersValidation(api) {
  if (!Array.isArray(api.servers) || api.servers.length !== 1) {
    throw new Error(
      `Expected exactly one root server in the OpenAPI spec, found ${
        Array.isArray(api.servers) ? api.servers.length : 0
      }. ${GUIDANCE}`,
    );
  }

  Object.keys(api.paths || {}).forEach((pathKey) => {
    const pathItem = api.paths[pathKey];
    if (!pathItem) {
      return;
    }
    if (pathItem.servers) {
      throw new Error(
        `Unexpected servers block at path level for "${pathKey}". ${GUIDANCE}`,
      );
    }
    validHttpMethods.forEach((method) => {
      const operation = pathItem[method];
      if (operation && operation.servers) {
        throw new Error(
          `Unexpected servers block at operation level for "${method.toUpperCase()} ${pathKey}". ${GUIDANCE}`,
        );
      }
    });
  });

  return api;
}

module.exports = { applyServersValidation };
