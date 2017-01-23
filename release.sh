#!/bin/bash

set -e # exit on error

SCRIPT="$0"

usage() {
    echo "${SCRIPT} VERSION" && exit 1
}

if [ $# -lt 1 ]; then usage; fi

VERSION="$1"

# change version in branch
git checkout "${VERSION}"
echo -n "${VERSION}" > VERSION
git add VERSION
git commit -m "[release] v${VERSION}"

# merge commits into master && create tag
git checkout master
git merge "${VERSION}"
git tag "v${VERSION}"
git push origin master
git push --tags
