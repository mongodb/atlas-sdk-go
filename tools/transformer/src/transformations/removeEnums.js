const removeField = require("../engine/removeField");

/**
 * Remove all enums from the schema.
 * Enums would not allow us to ensure API contract and will introduce breaking changes.
 * Enums should be defined as metadata in the schema.
 * @param {*} api OpenAPI JSON File
 * @param modelNames
 * @returns OpenAPI JSON File
 */
function applyRemoveEnumsTransformations(api) {
  // First handle schemas that are primarily enum definitions and are referenced
  handleEnumOnlySchemas(api);
  
  // Then remove enum fields from remaining schemas
  return removeField(api, "enum");
}

/**
 * Handle schemas that are primarily enum definitions and are referenced via $ref.
 * For such schemas, inline the base type where they're referenced and remove the schema.
 * @param {*} api OpenAPI JSON File
 */
function handleEnumOnlySchemas(api) {
  if (!api.components || !api.components.schemas) {
    return;
  }

  // Find all $ref references in the document
  const allRefs = [];
  findRefs(api, allRefs);
  
  // Extract schema references
  const referencedSchemas = new Set();
  allRefs.forEach((ref) => {
    const refParts = ref.split("/");
    if (refParts[1] === "components" && refParts[2] === "schemas") {
      referencedSchemas.add(refParts[3]);
    }
  });

  // Identify schemas that are primarily enum definitions
  const enumOnlySchemas = [];
  Object.keys(api.components.schemas).forEach((schemaName) => {
    const schema = api.components.schemas[schemaName];
    if (isEnumOnlySchema(schema) && referencedSchemas.has(schemaName)) {
      enumOnlySchemas.push(schemaName);
    }
  });

  // For each enum-only schema, replace references with inline type definition
  enumOnlySchemas.forEach((schemaName) => {
    const schema = api.components.schemas[schemaName];
    const inlineDefinition = createInlineDefinition(schema);
    
    // Replace all references to this schema with the inline definition
    replaceSchemaReferences(api, schemaName, inlineDefinition);
    
    // Remove the schema
    delete api.components.schemas[schemaName];
    
    console.info(`Removed enum-only schema '${schemaName}' and inlined its type definition`);
  });
}

/**
 * Check if a schema is primarily an enum definition
 * @param {*} schema 
 * @returns boolean
 */
function isEnumOnlySchema(schema) {
  // A schema is considered enum-only if:
  // 1. It has an enum field
  // 2. It has a type field (usually string)
  // 3. It doesn't have properties or complex structures
  return schema.enum && 
         schema.type && 
         !schema.properties && 
         !schema.allOf && 
         !schema.oneOf && 
         !schema.anyOf &&
         !schema.items;
}

/**
 * Create an inline definition from an enum schema
 * @param {*} schema 
 * @returns object
 */
function createInlineDefinition(schema) {
  const inlineDefinition = {
    type: schema.type
  };
  
  // Preserve important metadata but remove enum
  if (schema.description) {
    inlineDefinition.description = schema.description;
  }
  if (schema.example) {
    inlineDefinition.example = schema.example;
  }
  if (schema.title) {
    inlineDefinition.title = schema.title;
  }
  
  return inlineDefinition;
}

/**
 * Replace all references to a schema with an inline definition
 * @param {*} api 
 * @param {*} schemaName 
 * @param {*} inlineDefinition 
 */
function replaceSchemaReferences(api, schemaName, inlineDefinition) {
  const targetRef = `#/components/schemas/${schemaName}`;
  
  function replaceRefs(obj) {
    if (typeof obj !== "object" || obj === null) {
      return;
    }

    if (Array.isArray(obj)) {
      for (let i = 0; i < obj.length; i++) {
        replaceRefs(obj[i]);
      }
    } else {
      for (const key in obj) {
        if (key === "$ref" && obj[key] === targetRef) {
          // Replace the $ref with the inline definition
          delete obj[key];
          Object.assign(obj, inlineDefinition);
        } else {
          replaceRefs(obj[key]);
        }
      }
    }
  }
  
  replaceRefs(api);
}

/**
 * Recursive function for finding all $ref occurrences in the OpenAPI document
 * @param {*} obj
 * @param {*} refs
 */
function findRefs(obj, refs) {
  if (typeof obj === "object" && obj !== null) {
    if (Array.isArray(obj)) {
      obj.forEach((item) => findRefs(item, refs));
    } else {
      Object.keys(obj).forEach((key) => {
        if (key === "$ref") {
          refs.push(obj[key]);
        } else {
          findRefs(obj[key], refs);
        }
      });
    }
  }
}

module.exports = {
  applyRemoveEnumsTransformations,
};
