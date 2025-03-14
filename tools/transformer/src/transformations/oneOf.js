const ignoreModel = require("../engine/ignoreModel");
const {
  getObjectFromReference,
  getNameFromYamlPath,
  getObjectFromYamlPath,
  getAllObjects,
} = require("../engine/readers");
const { ignoreModels } = require("../oneOf.ignore.json");
const { ignoredProperties } = require("../duplicate.ignore.json");
const { detectDuplicates } = require("../engine/transformers");

/**
 * Transforms provided API JSON file using oneOf transformation
 *
 * @param {*} api OpenAPI JSON File
 * @param {*} oneOfTransformations array of transformations to apply
 * @returns OpenAPI JSON File
 */
function applyOneOfTransformations(api) {
  const oneOfTransformations = getAllObjects(api, (obj) => {
    return canApplyOneOfTransformation(obj, api);
  }).filter((e) => !ignoreModels.includes(e.path));

  transformationPaths = oneOfTransformations.map((e) => e.path);
  console.info(
    "# OneOf transformations: " +
      JSON.stringify(transformationPaths, undefined, 2),
  );

  for (let { path } of oneOfTransformations) {
    transformOneOf(path, api);
  }
  return api;
}

// Moves all the properties/enum values of the children into the parent
function transformOneOf(objectPath, api) {
  const parentObject = getObjectFromYamlPath(objectPath, api);
  const parentName = getNameFromYamlPath(objectPath);

  if (!parentObject.oneOf) {
    throw new Error(`Invalid object for OneOf Transformation: ${parentName}`);
  }

  // Expand references
  const childObjects = parentObject.oneOf.map((childRef) =>
    getObjectFromReference(childRef, api),
  );
  const isEnum = childObjects.reduce(
    (isEnum, childObject) => isEnum && childObject.enum,
    true,
  );

  if (isEnum) {
    transformOneOfEnum(parentObject, api);
  } else {
    transformOneOfProperties(parentObject, api);
  }
}

// Moves all the enum values of the children into the parent
function transformOneOfEnum(parentObject, api) {
  const childObjects = parentObject.oneOf.map((childRef) =>
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
  delete parentObject.oneOf;
}

// Moves all the propertis of the children into the parent
function transformOneOfProperties(parentObject, api) {
  const childObjects = parentObject.oneOf.map((childRef) =>
    getObjectFromReference(childRef, api),
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
        console.error(
          "OpenAPI object is missing properties or allOf field. This is usually an error in the OpenAPI spec.",
        );
        console.error(
          "Please ensure that elements of oneOf schema are objects (classes) instead of single types (bool, string etc.).",
        );
        throw new Error(`${JSON.stringify(childObject, "", 2)}`);
      }
    }
    console.debug(`${childObject.title}: moving child properties into parent`);
    handleDuplicates(parentObject, childObject);
    const childProperties = JSON.parse(JSON.stringify(childObject.properties));

    parentObject.properties = {
      ...parentObject.properties,
      ...childProperties,
    };
  }

  // Remove invalid fields
  delete parentObject.discriminator;
  delete parentObject.oneOf;
}

function handleDuplicates(parentObject, childObject) {
  const duplicates = detectDuplicates([
    parentObject.properties,
    childObject.properties,
  ]);
  if (duplicates.length > 0) {
    const duplicatesSource = childObject.title || "";
    let mismatches = duplicates.filter((e) => e.typeRefMismatch);
    if (mismatches.length > 0) {
      mismatches = filterReferenceOrTypeMissmatch(mismatches);
      if (mismatches.length !== 0) {
        throw new Error(
          `${duplicatesSource} mismatch type detected: ${JSON.stringify(
            mismatches,
            undefined,
            2,
          )}`,
        );
      }
    }
    mergeDuplicates(duplicatesSource, duplicates, childObject, parentObject);
  }
}

// Uses ignore list for known missmatches and removes them
function filterReferenceOrTypeMissmatch(mismatches) {
  return mismatches.filter((mismatch) => {
    for (const ignoredProperty of ignoredProperties) {
      if (mismatch.key === ignoredProperty) {
        console.warn(
          "Type mismatch found when merging base types. Ignoring due to known mismatches",
          JSON.stringify(mismatch, undefined, 2),
        );
        return false;
      }
    }
    return true;
  });
}

// Merge duplicates as they are representing the same type.
// Add description representing alternative meaning
function mergeDuplicates(
  duplicatesSourceLabel,
  duplicates,
  childObject,
  parentObject,
) {
  console.info(
    `## ${duplicatesSourceLabel} - Detected properties that would be overriden: ${JSON.stringify(
      duplicates,
    )}\n`,
  );
  for (duplicate of duplicates) {
    childProperty = childObject.properties[duplicate.key];
    parentProperty = parentObject.properties[duplicate.key];

    if (
      parentProperty.description &&
      childProperty.description !== parentProperty.description
    )
      childProperty.description = parentProperty.description;
  }
}

function canApplyOneOfTransformation(obj, api) {
  if (obj.oneOf) {
    return true;
  }
  return false;
}

module.exports = {
  applyOneOfTransformations,
  transformOneOf,
  transformOneOfProperties,
};
