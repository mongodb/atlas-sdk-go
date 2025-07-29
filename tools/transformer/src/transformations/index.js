const { applyAllOfTransformations, transformAllOf } = require("./allOf");
const { applyModelNameTransformations } = require("./name");
const {
  applyOneOfTransformations,
  transformOneOf,
  transformOneOfProperties,
} = require("./oneOf");
const { applyDiscriminatorTransformations } = require("./discriminator");
const { applyArrayTransformations } = require("./swapArray");
const { applyRemoveEnumsTransformations } = require("./removeEnums");
const {
  applyRemoveObjectAdditionalProperties,
} = require("./additionalPropertiesObject");
const { applyAnyOfTransformations } = require("./anyOf");
const { applyRemoveNullableTransformations } = require("./removeNullable");
const { removeRefsFromParameters } = require("./removeInvalidParams");
const { reorderResponseBodies } = require("./reorderResponseBodies");
const { applyFieldTransformations } = require("./fields");
const { applyOperationIdOverrides } = require("./operationId");

module.exports = {
  applyModelNameTransformations,
  transformAllOf,
  transformOneOf,
  transformOneOfProperties,
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyDiscriminatorTransformations,
  applyArrayTransformations,
  applyRemoveEnumsTransformations,
  applyRemoveObjectAdditionalProperties,
  applyAnyOfTransformations,
  applyRemoveNullableTransformations,
  removeRefsFromParameters,
  reorderResponseBodies,
  applyFieldTransformations,
  applyOperationIdOverrides,
};
