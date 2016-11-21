#!/bin/bash
aws s3 cp --quiet s3://moocfetcher/coursera/all.json .
echo "`cat all.json | jq '.courses | map(select(.courseType == "v2.ondemand")) | .[] | .slug ' | wc -l` courses total"
aws s3 cp --quiet s3://moocfetcher/coursera/ondemand/launched.json .
echo "`cat launched.json | jq '.courses[] | .slug' | wc -l` courses launched"
cat launched.json | jq -r  '{courses: [.courses[] | select(.primaryLanguageCodes) | select(.primaryLanguageCodes | map(. == "en") | any)]}' > english.json
echo "`cat english.json | jq '.courses[] | .slug' | wc -l` courses in English"
cat english.json | jq -r ".courses[] |  .slug" | sort > english.txt
aws s3 ls s3://moocfetcher-course-archive/coursera/ | grep PRE | awk '{print $2}' | tr -d '/' | sort > downloaded.txt
echo "`wc -l < downloaded.txt` courses in English downloaded"
comm -23 english.txt downloaded.txt > missing.txt
echo "`wc -l < missing.txt` courses in English missing"

