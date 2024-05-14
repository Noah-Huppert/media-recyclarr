#!/usr/bin/env bash
readonly PROG_DIR=$(dirname $(realpath "$0"))

# Arguments
DIR_NAME="$1"
shift
if [[ -z "$DIR_NAME" ]]; then
    echo "Error: DIR argument required" >&2
    exit 1
fi

readonly CLIENT_DIR="$PROG_DIR/../$DIR_NAME"
if [[ ! -d "$CLIENT_DIR" ]]; then
    echo "Error: '$CLIENT_DIR' does not exist" >&2
    exit 1
fi

# Remove
echo "Generating '$CLIENT_DIR'"
openapi-generator generate --config ."$CLIENT_DIR/openapi/openapi-generator.yml"