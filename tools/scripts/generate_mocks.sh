#!/usr/bin/env bash
set -o errexit
set -o nounset

MOCK_FOLDER=${MOCK_FOLDER:-../admin}
MOQ_INSTALLED=$(command -v moq)

if [[ -z $MOQ_INSTALLED ]]; then
  echo "moq is not installed. Installing moq to continue."
  go install github.com/matryer/moq@latest
fi


FILES=$(find "$MOCK_FOLDER" -type f -name "api*.go")

for FILE in $FILES; do
  FILENAME=$(basename "$FILE")
  FILENAME_API="${FILENAME%.*}API"
  FILENAME_MOCK="${FILENAME%.*}_mock.go"
  
  moq -out "$MOCK_FOLDER"/"$FILENAME_MOCK" "$MOCK_FOLDER" "$FILENAME_API"
done
