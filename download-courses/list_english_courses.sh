#!/bin/bash
cat launched.json | jq -r  '.courses[] | select(.primaryLanguageCodes) | select(.primaryLanguageCodes | map(. == "en") | any) |  .slug' | sort
