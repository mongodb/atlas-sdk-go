#!/usr/bin/env bash
set -euo pipefail

OUT_DIR="compliance"
PURL_FILE="${OUT_DIR}/purls.txt"
mkdir -p "$OUT_DIR"

# Generate purls for all dependencies (including indirect)
echo "==> Generating purls from go.mod dependencies"
go list -m all | tail -n +2 | awk '{print "pkg:golang/" $1 "@" $2}' | LC_ALL=C sort | uniq > "$PURL_FILE"

echo "PURLs written to $PURL_FILE"
