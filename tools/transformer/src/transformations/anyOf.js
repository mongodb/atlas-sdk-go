const {
  getObjectFromReference,
  getNameFromYamlPath,
  getObjectFromYamlPath,
  getAllObjects,
} = require("../engine/readers");
const { detectDuplicates } = require("../engine/transformers");

/**
 * Transforms provided API JSON file using anyOf transformation
 *
 * @param {*} api OpenAPI JSON File
 * @param {*} anyOfTransformations array of transformations to apply
 * @returns OpenAPI JSON File
 */
function applyAnyOfTransformations(api) {
  const anyOfTransformations = getAllObjects(api, (obj) => {
    return canApplyAnyOfTransformation(obj, api);
  });

  transformationPaths = anyOfTransformations.map((e) => e.path);

  console.info(
    "# AnyOf transformations: " +
      JSON.stringify(transformationPaths, undefined, 2)
  );

  for (let { path } of anyOfTransformations) {
    transformAnyOf(path, api);
  }
  return api;
}

// Moves all the properties/enum values of the children into the parent
function transformAnyOf(objectPath, api) {
  const parentObject = getObjectFromYamlPath(objectPath, api);
  const parentName = getNameFromYamlPath(objectPath);

  if (!parentObject.anyOf) {
    throw new Error(`Invalid object for AnyOf Transformation: ${parentName}`);
  }

  // Expand references
  const childObjects = parentObject.anyOf.map((childRef) =>
    getObjectFromReference(childRef, api)
  );
  const isEnum = childObjects.reduce(
    (isEnum, childObject) => isEnum && childObject.enum,
    true
  );

  if (isEnum) {
    transformAnyOfEnum(parentObject, api);
  } else {
    transformAnyOfProperties(parentObject, api);
  }
}

// Moves all the enum values of the children into the parent
function transformAnyOfEnum(parentObject, api) {
  const childObjects = parentObject.anyOf.map((childRef) =>
    getObjectFromReference(childRef, api)
  );

  if (!parentObject.enum) {
    parentObject.enum = [];
  }
  parentObject.enum = new Set(parentObject.enum);

  for (let childObject of childObjects) {
    console.debug(`${childObject.title}: moving child enum values into parent`);
    childObject.enum.forEach((enumValue) => parentObject.enum.add(enumValue));
    // Requires all enums to be same type
    parentObject.type = childObject.type;
  }

  parentObject.enum = [...parentObject.enum];

  // Remove invalid fields
  delete parentObject.discriminator;
  delete parentObject.anyOf;
}

// Moves all the propertis of the children into the parent
function transformAnyOfProperties(parentObject, api) {
  const childObjects = parentObject.anyOf.map((childRef) =>
    getObjectFromReference(childRef, api)
  );

  for (let childObject of childObjects) {
    if (!childObject.properties) {
      // When we missing props, we need to look at the allOf parent child inheritance
      if (childObject.allOf) {
        childObject.properties = {};
        for (allOfItem of childObject.allOf) {
          // We only add properties.
          if (allOfItem.properties) {
            childObject.properties = {
              ...childObject.properties,
              ...allOfItem.properties,
            };
          }
        }
      } else {
        // Otherwise this situation require human intervention
        throw new Error(`${JSON.stringify(childObject, "", 2)}`);
      }
    }
    const childProperties = JSON.parse(JSON.stringify(childObject.properties));
    console.debug(`${childObject.title}: moving child properties into parent`);
    const duplicates = detectDuplicates([
      parentObject.properties,
      childProperties,
    ]);
    if (duplicates.length > 0) {
      const duplicatesSource = childObject.title || "";
      console.info(
        `## ${duplicatesSource} - Detected properties that would be overriden: ${duplicates}\n`
      );
    }
    parentObject.properties = {
      ...parentObject.properties,
      ...childProperties,
    };
  }

  // Remove invalid fields
  delete parentObject.discriminator;
  delete parentObject.anyOf;
}

function canApplyAnyOfTransformation(obj, api) {
  if (obj.anyOf) {
    return true;
  }
  return false;
}

module.exports = {
  applyAnyOfTransformations,
};
