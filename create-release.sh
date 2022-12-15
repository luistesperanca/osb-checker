#!/bin/bash

#create draft
LASTEST_TAG=$(git describe --tags --abbrev=0)
echo "tag: " + $LASTEST_TAG
gh release create "$LASTEST_TAG" --generate-notes --draft

sleep 2
# python parse parse
counter=0
notes="release not found"
until [[ $notes -ne "release not found" ]]
do
  notes=$(gh release view $LASTEST_TAG --json body --jq .body)
  sleep 1
  if [ "$counter" -gt 400 ]; then
      echo "gh could not found release"
      exit 1
  fi
done

notes_edit=$(python parse_release.py "$notes")
echo "$notes_edit"
#
## push new release
gh release edit "$LASTEST_TAG" --notes "$notes_edit" --draft=false --latest