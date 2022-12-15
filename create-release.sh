#!/bin/bash

#create draft
LASTEST_TAG=$(git describe --tags --abbrev=0)
echo "tag: " + $LASTEST_TAG
gh release create "$LASTEST_TAG" --generate-notes --draft

# python parse parse
#notes=$(gh release view "$LASTEST_TAG" --json body --jq .body)
#notes_edit=$(python parse_release.py "$notes")
#echo "$notes_edit"
#
## push new release
#gh release edit "$LASTEST_TAG" --notes "$notes_edit" --draft