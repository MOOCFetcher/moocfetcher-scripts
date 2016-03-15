#!/bin/sh
~/courses/my-coursera/coursera/coursera-dl -n --aria2  -- $1
aws s3 cp --recursive $1 s3://moocfetcher-course-archive/coursera/$1 --storage-class STANDARD_IA && rm -rf $1
