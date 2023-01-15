#!/bin/bash

PROJECT_ROOT="$(readlink -f $(git rev-parse --git-dir) | sed 's/\/\.git.*//g')"
ANNICT_GRAPHQL_ENDPOINT=https://api.annict.com/graphql

if [ $# != 1 ]; then
    echo "Usage: $0 your-annict-api-token"
    exit 1
fi

API_TOKEN="$1"

set -o pipefall
set -o nounset
set -o errexit

get-graphql-schema -h "Authorization=Bearer "${API_TOKEN}" "${ANNICT_GRAPHQL_ENDPOINT}" > "${PROJECT_ROOT}/annict/schema.graphql"
