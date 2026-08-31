#!/bin/bash

set -e

if [ -z "$HERO_ID" ]; then
    echo "HERO_ID is not set" >&2
    exit 1
fi

curl -fsSL https://01.tomorrow-school.ai/assets/superhero/all.json \
    | jq -r --argjson id "$HERO_ID" '.[] | select(.id == $id) | .connections.relatives' \
    | tr -d '"'
