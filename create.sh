#!/usr/bin/env bash

# stricter bash
#  see http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail
IFS=$'\n\t'

SECTION=${1:-}
TITLE=${2:-}

if [[ -z "$SECTION" || -z "$TITLE" ]]; then
    echo "Usage: $0 SECTION TITLE"
    exit 1
fi

TITLE_SLUG="$(echo -n "$TITLE" | sed -e 's/[^[:alnum:]]/-/g' | tr -s '-' | tr A-Z a-z)"
DATE="$(date +"%Y-%m-%d")"
STUB="$SECTION/$DATE-$TITLE_SLUG"

hugo new "$STUB.md"
