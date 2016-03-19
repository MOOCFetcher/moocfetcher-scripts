#!/bin/sh
~/courses/my-coursera/coursera/coursera-dl -n --aria2  -- $1
aws s3 cp --storage-class STANDARD_IA --recursive $1 s3://moocfetcher-course-archive/coursera/$1 && rm -rf $1
