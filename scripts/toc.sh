#!/usr/bin/env bash
set -o errexit
set -o nounset

if which "./scripts/gh-md-toc" >/dev/null 2>&1; then
  echo "Binary gh-md-toc exists."
else
  wget https://raw.githubusercontent.com/ekalinin/github-markdown-toc/master/gh-md-toc
  chmod a+x gh-md-toc
fi

echo "generating table of contents"
echo "# Atlas SDK GO documentation" > ./docs/README.md
./scripts/gh-md-toc --skip-header ./docs/doc_*.md  >> ./docs/README.md
