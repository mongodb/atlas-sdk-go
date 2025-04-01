#!/bin/sh
rm -Rf ../mockadmin

# Use the locally installed mockery
if which mockery >/dev/null 2>&1; then
	mockery --dir ../mockadmin
else
	echo "mockery not found; please ensure it is installed."
	exit 1
fi

# Use any instead of interface{} in generated mocks
npm install
npm exec -c "replace-in-file /interface{}/g any  ../mockadmin/**/*.go --isRegex"
