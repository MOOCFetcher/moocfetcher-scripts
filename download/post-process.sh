#!/bin/sh

# Update course sizes
moocfetcher us

# Filter downloaded courses
moocfetcher fc -d /Volumes/courses

# Update annotated CSV file
csvjoin --outer -c "Folder on Disk" courses.csv courses-annotated.csv | csvcut -c 1,2,3,7,8,9 | csvsort -c 4,1 |  sponge courses-annotated.csv

