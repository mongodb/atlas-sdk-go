const { applyRemoveEnumsTransformations } = require("../src/transformations/removeEnums");

describe("removeEnums transformation", () => {
  test("should handle enum-only schemas referenced via $ref", () => {
    const api = {
      components: {
        schemas: {
          // This is an enum-only schema that gets referenced
          StatusEnum: {
            description: "Status of the resource",
            enum: ["ACTIVE", "INACTIVE", "PENDING"],
            example: "ACTIVE",
            title: "Status Types",
            type: "string"
          },
          // This schema references the enum-only schema
          Resource: {
            type: "object",
            properties: {
              id: {
                type: "string"
              },
              status: {
                $ref: "#/components/schemas/StatusEnum"
              }
            }
          }
        }
      },
      paths: {
        "/resource": {
          get: {
            responses: {
              "200": {
                content: {
                  "application/json": {
                    schema: {
                      $ref: "#/components/schemas/Resource"
                    }
                  }
                }
              }
            }
          }
        }
      }
    };

    const result = applyRemoveEnumsTransformations(api);

    // The StatusEnum schema should be removed
    expect(result.components.schemas.StatusEnum).toBeUndefined();

    // The reference should be replaced with inline type definition
    expect(result.components.schemas.Resource.properties.status).toEqual({
      description: "Status of the resource",
      example: "ACTIVE", 
      title: "Status Types",
      type: "string"
    });

    // Other properties should remain untouched
    expect(result.components.schemas.Resource.properties.id).toEqual({
      type: "string"
    });
  });

  test("should remove enums from regular properties", () => {
    const api = {
      components: {
        schemas: {
          Resource: {
            type: "object",
            properties: {
              status: {
                type: "string",
                enum: ["ACTIVE", "INACTIVE"],
                description: "Status of the resource"
              }
            }
          }
        }
      }
    };

    const result = applyRemoveEnumsTransformations(api);

    // The enum should be removed but other properties preserved
    expect(result.components.schemas.Resource.properties.status).toEqual({
      type: "string",
      description: "Status of the resource"
    });
  });

  test("should not remove schemas that have properties even if they have enums", () => {
    const api = {
      components: {
        schemas: {
          ComplexSchema: {
            type: "object",
            enum: ["VALUE1", "VALUE2"], // This has enum but also properties
            properties: {
              id: {
                type: "string"
              }
            }
          },
          Resource: {
            type: "object", 
            properties: {
              complex: {
                $ref: "#/components/schemas/ComplexSchema"
              }
            }
          }
        }
      }
    };

    const result = applyRemoveEnumsTransformations(api);

    // ComplexSchema should still exist because it has properties
    expect(result.components.schemas.ComplexSchema).toBeDefined();
    expect(result.components.schemas.ComplexSchema.properties).toBeDefined();
    
    // But enum should be removed
    expect(result.components.schemas.ComplexSchema.enum).toBeUndefined();
    
    // Reference should remain unchanged
    expect(result.components.schemas.Resource.properties.complex).toEqual({
      $ref: "#/components/schemas/ComplexSchema"
    });
  });

  test("should handle oneOf with enum schemas", () => {
    const api = {
      components: {
        schemas: {
          // Enum-only schemas that will be used in oneOf
          StatusEnum: {
            description: "Status enumeration",
            enum: ["ACTIVE", "INACTIVE"],
            example: "ACTIVE",
            title: "Status",
            type: "string"
          },
          PriorityEnum: {
            description: "Priority enumeration", 
            enum: ["HIGH", "MEDIUM", "LOW"],
            example: "HIGH",
            title: "Priority",
            type: "string"
          },
          // Schema with oneOf referencing enum schemas
          Resource: {
            type: "object",
            properties: {
              id: {
                type: "string"
              },
              category: {
                oneOf: [
                  { $ref: "#/components/schemas/StatusEnum" },
                  { $ref: "#/components/schemas/PriorityEnum" }
                ]
              }
            }
          }
        }
      }
    };

    const result = applyRemoveEnumsTransformations(api);

    // The enum-only schemas should be removed
    expect(result.components.schemas.StatusEnum).toBeUndefined();
    expect(result.components.schemas.PriorityEnum).toBeUndefined();

    // The oneOf should be preserved but with inlined type definitions
    expect(result.components.schemas.Resource.properties.category).toEqual({
      oneOf: [
        {
          description: "Status enumeration",
          example: "ACTIVE",
          title: "Status", 
          type: "string"
        },
        {
          description: "Priority enumeration",
          example: "HIGH", 
          title: "Priority",
          type: "string"
        }
      ]
    });

    // Other properties should remain untouched
    expect(result.components.schemas.Resource.properties.id).toEqual({
      type: "string"
    });
  });
});
