module.exports = function reorderAndCleanResponseBodies(api) {
    const hasResponses = api && api.paths;
    if (!hasResponses) {
      throw new Error("Missing paths in OpenAPI");
    }
  
    // Helper function to process media types
    function processContent(content) {
      // Sort vnd.atlas types and remove application/json
      const sortedContent = {};
  
      // Sort by date (vnd.atlas first), and remove application/json
      Object.keys(content)
        .sort((a, b) => {
          // Check for vnd.atlas and prioritize by date
          const atlasA = a.match(/application\/vnd\.atlas\.(\d{4}-\d{2}-\d{2})\+json/);
          const atlasB = b.match(/application\/vnd\.atlas\.(\d{4}-\d{2}-\d{2})\+json/);
          
          if (atlasA && atlasB) {
            return new Date(atlasB[1]) - new Date(atlasA[1]); // Newer date first
          } else if (atlasA) {
            return -1; // Atlas should be first
          } else if (atlasB) {
            return 1;
          } else {
            return 0;
          }
        })
        .forEach(key => {
          if (key !== "application/json") {
            sortedContent[key] = content[key]; // Keep only non-application/json
          }
        });
  
      return sortedContent;
    }
  
    // Recursive function to traverse the OpenAPI object
    function reorderAndClean(obj) {
      if (typeof obj !== "object" || obj === null) {
        return;
      }
  
      if (Array.isArray(obj)) {
        for (let i = 0; i < obj.length; i++) {
          reorderAndClean(obj[i]);
        }
      } else {
        // Process response bodies
        if (obj.content) {
          obj.content = processContent(obj.content);
        }
  
        // Traverse nested properties
        for (const prop in obj) {
          if (obj.hasOwnProperty(prop)) {
            reorderAndClean(obj[prop]);
          }
        }
      }
    }
  
    // Iterate over each path and method in the OpenAPI spec
    Object.keys(api.paths).forEach((path) => {
      const methods = api.paths[path];
      Object.keys(methods).forEach((method) => {
        const operation = methods[method];
        if (operation.responses) {
          reorderAndClean(operation.responses);
        }
      });
    });
    return api;
  };
  