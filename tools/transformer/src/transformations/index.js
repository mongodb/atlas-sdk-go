const { applyAllOfTransformations, transformAllOf } = require("./allOf");
const { applyModelNameTransformations } = require("./name");
const { applyOneOfTransformations, transformOneOf } = require("./oneOf");
const { applyDiscriminatorTransformations } = require("./discriminator");
const { applyArrayTransformations } = require("./swapArray");

module.exports = {
  applyModelNameTransformations,
  transformAllOf,
  transformOneOf,
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyDiscriminatorTransformations,
  applyArrayTransformations,
};
