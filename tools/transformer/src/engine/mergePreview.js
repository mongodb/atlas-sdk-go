/**
 * Overlay a preview OpenAPI delta (openapi-preview.yaml) onto a full base spec,
 * so the result has every base endpoint plus the preview ones.
 *
 * For paths, components.*, and tags: preview-only entries are added, and
 * entries with the same key are replaced by the preview version
 * (no field-level merge). Everything else is kept from the base.
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
