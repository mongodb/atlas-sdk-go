#!/usr/bin/env bash
set -o errexit
set -ox nounset

MOCK_FOLDER=${MOCK_FOLDER:-../admin}
MOQ_INSTALLED=$(command -v moq)

if [[ -z $MOQ_INSTALLED ]]; then
  echo "moq is not installed. Installing moq to continue."
  go install github.com/matryer/moq@latest
fi


FILES=$(find "$MOCK_FOLDER" -type f -name "api*.go")

for FILE in $FILES; do
  FILENAME=$(basename "$FILE")
  echo "$FILENAME"
  FILENAME_API=$(echo "$FILENAME" | awk -F'[_.]' '{print $2}' | awk '{print toupper(substr($0, 1, 1)) tolower(substr($0, 2))}')Api
  FILENAME_MOCK="${FILENAME%.*}_mock.go"
  echo "Generating mock for $FILENAME_API"
  moq -out "$MOCK_FOLDER"/"$FILENAME_MOCK" "$MOCK_FOLDER" "$FILENAME_API"
  echo "Finished generation of mock for $FILENAME_API"
done
