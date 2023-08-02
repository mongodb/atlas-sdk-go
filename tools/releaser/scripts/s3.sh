#!/bin/bash
set -ueo pipefail

## Internal script
# Not to be used directly

script_path=$(dirname "$0")
# shellcheck source=/dev/null
source "$script_path/extract-version.sh"

S3_BUCKET="mongodb-golang-url"
LOCAL_FILE_PATH="$script_path/../resources/index.html"
## Indentionally left as manual file to avoid accidental upload/values
NEW_FOLDER_NAME="atlas-sdk/v20230201005"

aws s3api head-object --bucket "$S3_BUCKET" --key "atlas-sdk/index.html"

# Create the new folder with the specified name
aws s3api put-object --bucket "$S3_BUCKET" --key "$NEW_FOLDER_NAME/"

# Upload the local file to the newly created folder
aws s3 cp "$LOCAL_FILE_PATH" "$S3_BUCKET$NEW_FOLDER_NAME/"
