#!/bin/sh
# This wrapper can be used to run the codegen from a go generate comment.
# We cannot use go run here because we have more than one file to build.

pkg="$(dirname "$0")"
set -x
go build -o "${pkg}/codegen" "${pkg}" && exec "${pkg}/codegen" $@
