#!/bin/bash
set -a
source .env
set +a

DAY=$1

curl "https://adventofcode.com/2023/day/$DAY/input" \
  -b "session=$AOC_SESSION" > inputs/$(printf "%02d" $DAY).txt
