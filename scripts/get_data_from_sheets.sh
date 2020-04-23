#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

if [[ -z "${GSHEET_ID}" ]]; then
  echo "'${GSHEET_ID}' is not set."
  exit
else
  SHEET="${GSHEET_ID}"
fi

declare -a sheets_to_fetch=("read" "unread" "queued" "currently_reading" "recently_finished" )

for val in "${sheets_to_fetch[@]}"; do
  gsheets --key="${SHEET}" --title="${val}" --pretty > "${val}".json
  if [ $? -ne 0 ]; then
    exit
  fi
done

mkdir -p data

jq -s '[.[] | to_entries] | flatten | reduce .[] as $dot ({}; .[$dot.key] += $dot.value)' currently_reading.json queued.json recently_finished.json unread.json read.json | jq '.data' > ./data/library.json

rm currently_reading.json queued.json recently_finished.json unread.json read.json