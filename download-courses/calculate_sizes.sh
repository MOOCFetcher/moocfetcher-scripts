#!/bin/sh

coursesSize() {
  cat $1 | jq  ".courses | map(.size) | add  | ./(1024*1024*1024) | floor"
}

echo "All courses: `coursesSize launched.json`GB"
echo "English courses: `coursesSize english.json`GB"
