const fieldTransformations = require("../field.transformations.json");

/**
 * Makes specified fields optional in their respective schemas
 * @param {*} api OpenAPI JSON File
 * @returns OpenAPI JSON File
 */
function applyFieldTransformations(api) {
	const { optionalFields } = fieldTransformations;

	// Make specified fields optional in their schemas
	for (const [schemaName, fields] of Object.entries(optionalFields)) {
		const schema = api.components.schemas[schemaName];
		if (schema && schema.required) {
			schema.required = schema.required.filter(field => !fields.includes(field));
			if (schema.required.length === 0) {
				delete schema.required;
			}
		}
	}

	return api;
}

module.exports = {
	applyFieldTransformations,
}; 