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
};
