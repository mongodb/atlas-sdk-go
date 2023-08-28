// **  Remove invalid parameters that do not apply for the end user
function removeRefsFromParameters(openapiJson, refsToRemove) {
  if (!openapiJson || typeof openapiJson !== "object") {
    throw new Error("Invalid input. Expected an OpenAPI JSON object.");
  }

  if (!Array.isArray(refsToRemove)) {
    throw new Error("Invalid input. Expected an array of $ref strings.");
  }

  if (!openapiJson.paths) {
    throw new Error('Input JSON does not contain "paths" property.');
  }

  const paths = openapiJson.paths;

  for (const path in paths) {
    const pathObject = paths[path];

    for (const method in pathObject) {
      const operation = pathObject[method];

      if (operation.parameters && Array.isArray(operation.parameters)) {
        operation.parameters = operation.parameters.filter((parameter) => {
          if (!refsToRemove.includes(parameter.$ref)) {
            return true;
          }
          return false;
        });
      }
    }
  }

  return openapiJson;
}

module.exports = { removeRefsFromParameters };
