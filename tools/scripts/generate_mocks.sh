#!/bin/sh
rm -Rf ../mockadmin

# install only if not already present
if ! which mockery ; then
	docker run --rm -v "$(dirname "$PWD")":/src --workdir=/src vektra/mockery:v2.42.1
else
	mockery --dir ../mockadmin
fi
