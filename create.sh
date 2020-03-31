#!/usr/bin/env bash

# stricter bash
# see http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -uo pipefail
IFS=$'\n\t'
RED='\033[0;31m'
NC='\033[0m' # No Color

SECTION=${1:-}
TITLE=${2:-}
CONTENT_DIR="./content"

# e.g. - ./create.sh posts "2019 - year in review"
if [[ -z "$SECTION" || -z "$TITLE" ]]; then
    echo "Usage: $0 SECTION TITLE"
    exit 1
fi

DATE="$(date +"%Y-%m-%d")"
TITLE_SLUG="$(echo -n "$TITLE" | sed -e 's/[^[:alnum:]]/-/g' | tr -s '-' | tr A-Z a-z)"
DIR="${CONTENT_DIR}/${1}/${DATE}-${TITLE_SLUG}"

if [ -d "${DIR}" ]
then
  echo "${DIR} already exists."
  echo "Aborting..."
  exit
fi

echo "Following directory will be created - ${DIR}"
read -p "Press [Enter] key to confirm..."

mkdir ${DIR}
echo "Creating ${DIR}/index.md in $SECTION..."
hugo new -k $SECTION ${DIR}/index.md
retVal=$?

if [ $retVal -ne 0 ]
then
  rm -r ${DIR}
  echo -e "${RED}Some error occured"
else
  echo ""
  echo -e "${RED}Do not forget to update the title in the frontmatter."
fi
