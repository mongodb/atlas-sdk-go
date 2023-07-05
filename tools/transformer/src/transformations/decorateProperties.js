/**
 * For each discriminator mapping, add the discriminator value to the children property description
 * 
 * @param {*} openapi 
 * @returns 
 */
function applyDecoratePropertiesWithInheritanceContext(openapi) {
    const schemas = openapi.components.schemas;
    
    Object.keys(schemas).forEach(schemaName => {
      const schema = schemas[schemaName];
      const discriminator = schema.discriminator;
      
      if (discriminator && discriminator.mapping) {
        const mapping = discriminator.mapping;
        Object.keys(mapping).forEach(mappingKey => {
          if (!mapping[mappingKey]) return;
         
          const subtypeRef = mapping[mappingKey];
          const subtypeName = subtypeRef.substring(subtypeRef.lastIndexOf('/') + 1);
          const subtypeProperties = schemas[subtypeName].properties;
          if(!subtypeProperties) return;
          Object.keys(subtypeProperties).forEach(propertyName => {
            prefixFormat = `[${discriminator.propertyName}=${mappingKey}]`

            const property = subtypeProperties[propertyName];
            property.description = property.description
              ? prefixFormat + ' ' + property.description 
              : mappingKey;
          });
        });
      }
    });
    
    return openapi;
  }
  
module.exports = applyDecoratePropertiesWithInheritanceContext;
