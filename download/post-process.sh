#!/bin/sh
AWS_PROFILE=moocfetcher
# Update course sizes
moocfetcher us

# Filter downloaded courses
moocfetcher fc -d $COURSE_VOLUME

# Update annotated CSV file
csvjoin --delimiter "," --quotechar "\"" --outer -c "Folder on Disk" courses.csv courses-annotated.csv | csvcut  --delimiter "," --quotechar "\"" -c 1,2,3,7,8,9 | csvsort --delimiter "," --quotechar "\"" -c 4,1 | sponge courses-annotated.csv

# Generate excel file
python csv2excel.py
