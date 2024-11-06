/**
 * Reorder responses by using application/vnd.*json with the most recent date as the first valid response.
 * @param {*} api OpenAPI JSON File
 * @returns OpenAPI JSON File
 */
function reorderResponseBodies(api) {
  Object.values(api.paths).forEach((methods) => {
    Object.values(methods).forEach((operation) => {
      traverse(operation.responses);
    });
  });
  return api;

  function traverse(responses) {
    if (typeof responses !== "object") {
      return;
    }

    for (const prop in responses) {
      traverse(responses[prop]);
    }

    if (responses.content) {
      responses.content = sortByMediaType(responses.content);
    }
  }

  function sortByMediaType(content) {
    const ret = {};
    Object.keys(content)
      .sort((a, b) => compareMediaType(a, b))
      .forEach((key) => (ret[key] = content[key]));
    return ret;
  }

  function compareMediaType(a, b) {
    const matcher = /application\/vnd\.atlas\.(\d{4}-\d{2}-\d{2})\+json/;
    const atlasA = a.match(matcher);
    const atlasB = b.match(matcher);

    if (atlasA && atlasB) {
      return new Date(atlasB[1]) - new Date(atlasA[1]); // Newer date first
    } else if (atlasA) {
      return -1; // vnd.atlas json first
    } else if (atlasB) {
      return 1;
    } else {
      return a - b; // Default ordering if not vnd.atlas json
    }
  }
}

module.exports = {
  reorderResponseBodies,
};
