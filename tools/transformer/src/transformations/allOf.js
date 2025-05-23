const {
  getObjectFromReference,
  getNameFromYamlPath,
  getObjectFromYamlPath,
  getObjectNameFromReference,
  getAllObjects,
  getObjectProperties,
  isSchema,
} = require("../engine/readers");

const {
  removeParentFromAllOf,
  flattenAllOfObject,
} = require("../engine/transformers");

/**
 * Transforms provided API JSON file using allOf transformation
 *
 * @param {*} api OpenAPI JSON File
 * @param {*} allOfTransformations array of transformations to apply
 * @returns OpenAPI JSON File
 */
function applyAllOfTransformations(api) {
  const allOfTransformations = getAllObjects(api, isTransformable);

  console.info(
    "# AllOf transformations: ",
    allOfTransformations.map((e) => e.path),
  );

  for (let { path } of allOfTransformations) {
    if (isSchema(path)) {
      transformAllOf(path, api);
    }
  }
  return api;
}

// Moves all the properties of the parent into the children
function transformAllOf(objectPath, api) {
  const parentObject = getObjectFromYamlPath(objectPath, api);
  const parentName = getNameFromYamlPath(objectPath);

  if (!isTransformable(parentObject)) {
    throw new Error(`Invalid object for AllOf Transformation: ${parentName}`);
  }

  const expandedParent = getObjectProperties(parentObject);

  for (let childRef of parentObject.oneOf) {
    const childObject = getObjectFromReference(childRef, api);
    const childName = getObjectNameFromReference(childRef);
    if (!childObject) {
      throw new Error(
        `Missing object reference: ${childRef} for ${parentName}`,
      );
    }
    if (removeParentFromAllOf(childObject, parentName)) {
      console.debug(
        `AllOf: Moving ${parentName} (parent) properties into ${childName} (child) properties`,
      );
      if (!childObject.allOf) {
        childObject.allOf = {};
      }
      // Expand parent in child allOf
      childObject.allOf.push(expandedParent);
      // Flatten child allOf
      flattenAllOfObject(childObject);
    }
  }
  // Remove invalid fields
  delete parentObject.properties;
  delete parentObject.required;
}

function isTransformable(obj) {
  return obj.oneOf;
}

module.exports = {
  applyAllOfTransformations,
  transformAllOf,
};
