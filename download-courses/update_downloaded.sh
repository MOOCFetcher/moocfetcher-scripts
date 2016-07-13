#!/bin/bash
aws s3 cp --quiet s3://moocfetcher/coursera/ondemand/launched.json .
cat launched.json | jq -r ".courses[] |  .slug" | sort > launched.txt
echo "`wc -l < launched.txt` courses launched"
aws s3 ls s3://moocfetcher-course-archive/coursera/ | grep PRE | awk '{print $2}' | tr -d '/' | sort > downloaded.txt
echo "`wc -l < downloaded.txt` courses downloaded"
cat launched.json | jq -r  '{courses: [.courses[] | select(.primaryLanguageCodes) | select(.primaryLanguageCodes | map(. == "en") | any)]}' > english.json
echo "`cat english.json | jq '.courses[] | .slug' | wc -l` courses in English"
comm -23 launched.txt downloaded.txt > missing.txt
echo "`wc -l < missing.txt` courses missing"

