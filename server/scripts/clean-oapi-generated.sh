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

# Options
ARG_DRY_RUN=""
while getopts "d" opt; do
    case "$opt" in
        d) ARG_DRY_RUN="y" ;;
    esac
done

# Remove
RM_CMD="rm"
if [[ -n "$ARG_DRY_RUN" ]]; then
    RM_CMD="echo"
fi

echo "Cleaning '$CLIENT_DIR'"
find "$CLIENT_DIR" -mindepth 1 -type f -not -path "$CLIENT_DIR/openapi/*" | xargs "$RM_CMD"
find "$CLIENT_DIR" -mindepth 1 -type d ! -name openapi | xargs "$RM_CMD" -rf