#!/usr/bin/env bash
set -o errexit
set -o nounset

BIN_FOLDER=./bin
DOCS_FOLDER=./docs

pushd $BIN_FOLDER

if which "./gh-md-toc" >/dev/null 2>&1; then
  echo "Binary gh-md-toc exists."
else
  echo Downloading scripts

  curl -O https://raw.githubusercontent.com/ekalinin/github-markdown-toc/master/gh-md-toc
  chmod a+x gh-md-toc
fi

docsbin=$(readlink -f "./gh-md-toc")

popd

echo "generating table of contents"
echo "# Atlas SDK GO documentation" > ./docs/README.md

pushd $DOCS_FOLDER

$docsbin --skip-header ./doc_*.md  >> ./README.md

popd
