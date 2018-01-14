#!/bin/bash

set -euo pipefail

curl -L $1 | pup '.rc-BannerBreadcrumbs .item:nth-of-type(2) a, .about-section-wrapper div.content-inner p, .rc-CreatorInfo div span:nth-of-type(2) json{}' | jq '.[] | .text'

