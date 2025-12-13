#!/bin/bash
set -euf -o pipefail

# functions
function echogrey() {
	echo -e "\033[0;90m$1\033[0m"
}

function template() {
	cat <<EOF
package main

import (
    "fmt"
    
	aoc "github.com/D-P-Williams/Advent-of-Code"
)

func main() {
    lines := aoc.ReadLines("./input.txt")

    countPt1 := 0
	countPt2 := 0

	fmt.Println("part 1", countPt1)

	fmt.Println("part 2", countPt2)
}

EOF
}

# two args YEAR and DAY
YEAR="${1:-}"
DAY_RAW="${2:-}"
if [ -z "$YEAR" ] || [ -z "$DAY_RAW" ]; then
	echo "Usage: $0 <YEAR> <DAY>"
	exit 1
    fi
# pad DAY to 2 digits
DAY=$(printf "%02d" $DAY_RAW)
DIR="./$YEAR/$DAY"
# create missing files as needed
if [ ! -d "$DIR" ]; then
	mkdir -p "$DIR"
	echogrey "Created directory $DIR"
fi
if [ ! -f "$DIR/code.go" ]; then
	template >"$DIR/code.go"
	echogrey "Created file code.go"
fi

# After checking if the file is already cached
url="https://adventofcode.com/$YEAR/day/$DAY_RAW/input"
echo $url
curl -H "Cookie: session=$AOC_SESSION" $url
curl -H "Cookie: session=$AOC_SESSION" $url > $DIR/input.txt

cd "$DIR"