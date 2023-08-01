## Internal script
# Not to be used directly

script_path=$(dirname "$0")
# shellcheck source=/dev/null
source "$script_path/extract-version.sh"

S3_BUCKET="s3://mongodb-golang-url/atlas-sdk/"
LOCAL_FILE_PATH="/path/to/local/index.html"
## Indentionally left as manual file to avoid accidental upload/values
NEW_FOLDER_NAME="v20230201005"

# Create the new folder with the specified name
aws s3api put-object --bucket "$S3_BUCKET" --key "$NEW_FOLDER_NAME/" --acl private

# Upload the local file to the newly created folder
aws s3 cp "$LOCAL_FILE_PATH" "$S3_BUCKET$NEW_FOLDER_NAME/"
