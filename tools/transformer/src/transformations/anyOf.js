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
 * @returns OpenAPI JSON File
 */
function applyAnyOfTransformations(api) {
  const anyOfTransformations = getAllObjects(api, (obj) => {
    return canApplyAnyOfTransformation(obj, api);
  });

  const transformationPaths = anyOfTransformations.map((e) => e.path);

  console.info(
    "# AnyOf transformations: " +
      JSON.stringify(transformationPaths, undefined, 2),
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
    getObjectFromReference(childRef, api),
  );
  const isEnum = childObjects.reduce(
    (isEnum, childObject) => isEnum && childObject.enum,
    true,
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
    getObjectFromReference(childRef, api),
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

// Moves all the properties of the children into the parent
function transformAnyOfProperties(parentObject, api) {
  const childObjects = parentObject.anyOf.map((childRef) =>
    getObjectFromReference(childRef, api),
  );

  for (let childObject of childObjects) {
    if (childObject.allOf) {
      // Children may define properties both directly and in inline allOf entries.
      childObject.properties = { ...childObject.properties };
      for (const allOfItem of childObject.allOf) {
        // We only add properties.
        if (allOfItem.properties) {
          childObject.properties = {
            ...childObject.properties,
            ...allOfItem.properties,
          };
        }
      }
    } else if (!childObject.properties) {
      // No direct properties and no allOf, invalid schema.
      console.error(
        "OpenAPI object is missing properties or allOf field. This is usually an error in the OpenAPI spec.",
      );
      console.error(
        "Please ensure that elements of anyOf schema are objects (classes) instead of single types (bool, string etc.).",
      );
      throw new Error(`${JSON.stringify(childObject, "", 2)}`);
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
        `## ${duplicatesSource} - Detected properties that would be overriden: ${duplicates}\n`,
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

function canApplyAnyOfTransformation(obj, _) {
  return !!obj.anyOf;
}

module.exports = {
  applyAnyOfTransformations,
};
