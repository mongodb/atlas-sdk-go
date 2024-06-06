const {
  getObjectFromReference,
  getObjectNameFromReference,
} = require("../engine/readers");

function filterObjectProperties(object, filter = (_k, _v) => true) {
  const filteredObj = Object.keys(object)
    .filter((key) => filter(key, object[key]))
    .reduce((aggregationObj, key) => {
      aggregationObj[key] = object[key];
      return aggregationObj;
    }, {});

  return filteredObj;
}

// Detects duplicates in array of object properties
// Function detects duplicates in type and $ref values.
function detectDuplicates(objArray) {
  const allKeys = {};
  const duplicates = [];

  for (const obj of objArray) {
    if (obj) {
      for (const key of Object.keys(obj)) {
        const currentEntry = obj[key];
        const currentValue = currentEntry.type || currentEntry.$ref;

        if (allKeys[key]) {
          const previousEntry = allKeys[key];
          const previousValue = previousEntry.type || previousEntry.$ref;

          // Check for typeRefMismatch
          const isTypeRefMismatch = previousValue !== currentValue;

          if (isTypeRefMismatch) {
            duplicates.push({
              key,
              firstValue: previousValue,
              secondValue: currentValue,
              typeRefMismatch: isTypeRefMismatch,
            });
          }
        } else {
          allKeys[key] = currentEntry;
        }
      }
    }
  }

  return duplicates;
}

function expandReference(obj, apiObject) {
  if (!obj || !obj["$ref"]) {
    return obj;
  }

  const dereferencedObj = getObjectFromReference(obj, apiObject);
  delete obj["$ref"];

  Object.keys(dereferencedObj).forEach((key) => {
    obj[key] = dereferencedObj[key];
  });

  return obj;
}

function flattenAllOfObject(obj, apiObject) {
  if (!obj.allOf) {
    return obj;
  }

  if (!obj.properties) {
    obj.properties = {};
  }

  if (!obj.required) {
    obj.required = new Set();
  } else {
    obj.required = new Set(obj.required);
  }

  for (let parent of obj.allOf) {
    parent = expandReference(parent, apiObject);
    if (obj.properties) {
      obj.properties = mergeObjects(obj.properties, parent.properties);
    }

    if (parent.required) {
      parent.required.forEach((item) => obj.required.add(item));
    }
  }
  obj.required = [...obj.required];
  delete obj.allOf;

  return obj;
}

function mergeObjects(...objs) {
  let mergedObj = {};

  for (let obj of objs) {
    if (obj) {
      mergedObj = { ...mergedObj, ...obj };
    }
  }

  return mergedObj;
}

function removeParentFromAllOf(child, parentName) {
  if (!child.allOf) {
    return false;
  }

  const initialLength = child.allOf.length;
  child.allOf = child.allOf.filter((childAllOfItem) => {
    const objName = getObjectNameFromReference(childAllOfItem);
    return objName !== parentName;
  });

  if (initialLength === child.allOf.length) {
    return false;
  }

  return true;
}

function resolveOpenAPIReference(openapi, ref) {
  if (!ref.startsWith("#/")) {
    throw new Error("Invalid reference format: " + ref);
  }
  const parts = ref.split("/");
  // Skip the first empty part
  parts.shift();

  let current = openapi;
  for (const part of parts) {
    if (!current.hasOwnProperty(part)) {
      throw new Error("Reference not found: + ref");
    }
    current = current[part];
  }
  return current;
}


module.exports = {
  removeParentFromAllOf,
  detectDuplicates,
  mergeObjects,
  filterObjectProperties,
  flattenAllOfObject,
  resolveOpenAPIReference
};
