const { applyAllOfTransformations } = require("./allOf");
const { applyOneOfTransformations } = require("./oneOf");
const { applyAnyOfTransformations } = require("./anyOf");
const { applyModelNameTransformations } = require("./name");
const { applyDiscriminatorTransformations } = require("./discriminator");
const { applyRemoveEnumsTransformations } = require("./removeEnums");
const { applyRemoveObjectAdditionalProperties } = require("./additionalPropertiesObject");
const { applyRemoveNullableTransformations } = require("./removeNullable");
const { removeRefsFromParameters } = require("./removeInvalidParams");
const { reorderResponseBodies } = require("./reorderResponseBodies");
const { applyFieldTransformations } = require("./fields");

module.exports = {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyAnyOfTransformations,
  applyModelNameTransformations,
  applyDiscriminatorTransformations,
  applyRemoveEnumsTransformations,
  applyRemoveObjectAdditionalProperties,
  applyRemoveNullableTransformations,
  removeRefsFromParameters,
  reorderResponseBodies,
  applyFieldTransformations,
};
