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
  applyRemoveObjectAdditonalProperties,
} = require("./additionalPropertiesObject");
const { applyAddExperimentalTag } = require("./tagExperimental");
const { applyAnyOfTransformations } = require("./anyOf");
const applyDecoratePropertiesWithInheritanceContext  = require("./decorateProperties");

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
  applyRemoveObjectAdditonalProperties,
  applyAddExperimentalTag,
  applyAnyOfTransformations,
  applyDecoratePropertiesWithInheritanceContext,
};
