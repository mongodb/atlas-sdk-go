/**
 * Merge a preview OpenAPI delta spec onto a full base (stable) spec.
 *
 * The published `openapi-preview.yaml` only contains the operations that have
 * preview content (a small delta), not the full API surface. To produce a
 * usable preview SDK we overlay that delta on top of the latest stable spec so
 * the result is a superset: every stable endpoint plus the preview ones.
 *
 * Merge rules:
 *  - paths: preview paths/methods are added; a method present in both is
 *    overridden by the preview version (so preview media types win).
 *  - components.*: preview entries (schemas, parameters, responses,
 *    requestBodies, headers, securitySchemes, ...) are added/overridden so that
 *    $refs used by preview operations resolve.
 *  - tags: preview tags not already present are appended.
 *  - info/servers/security and everything else are kept from the base.
 *
 * @param {*} base full stable OpenAPI JSON doc (mutated and returned)
 * @param {*} preview preview delta OpenAPI JSON doc
 * @returns merged base doc
 */
function mergePreview(base, preview) {
  if (!preview) {
    return base;
  }

  // Merge paths (operation-level override).
  if (preview.paths) {
    base.paths = base.paths || {};
    for (const [path, previewItem] of Object.entries(preview.paths)) {
      if (base.paths[path]) {
        base.paths[path] = { ...base.paths[path], ...previewItem };
      } else {
        base.paths[path] = previewItem;
      }
    }
  }

  // Merge components subsections (entry-level override).
  if (preview.components) {
    base.components = base.components || {};
    for (const [section, entries] of Object.entries(preview.components)) {
      if (entries && typeof entries === "object") {
        base.components[section] = {
          ...(base.components[section] || {}),
          ...entries,
        };
      }
    }
  }

  // Append preview tags that are not already declared.
  if (Array.isArray(preview.tags)) {
    base.tags = base.tags || [];
    const existing = new Set(base.tags.map((t) => t && t.name));
    for (const tag of preview.tags) {
      if (tag && !existing.has(tag.name)) {
        base.tags.push(tag);
      }
    }
  }

  return base;
}

module.exports = { mergePreview };
