#!/bin/sh
rm -Rf ../mockadmin

# install only if not already present
if ! which mockery; then
	docker run --rm -u "$(id -u):$(id -g)" -v "$(dirname "$PWD")":/src --workdir=/src vektra/mockery:v2.42.1
else
	mockery --dir ../mockadmin
fi

# Use any instead of interface{} in genereated mocks
npm install
npm exec -c "replace-in-file /interface{}/g any  ../mockadmin/**/*.go --isRegex"
