#!/usr/bin/env bash
set -o errexit
set -o nounset

if which "./scripts/gh-md-toc" >/dev/null 2>&1; then
  echo "Binary gh-md-toc exists."
else
  echo Downloading scripts
  pushd ./scripts 
  wget https://raw.githubusercontent.com/ekalinin/github-markdown-toc/master/gh-md-toc
  chmod a+x gh-md-toc
  popd
fi

echo "generating table of contents"
echo "# Atlas SDK GO documentation" > ./docs/README.md
cd  ./docs/ && ../scripts/gh-md-toc --skip-header ./doc_*.md  >> ./README.md
