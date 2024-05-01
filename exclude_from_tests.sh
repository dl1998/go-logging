#!/bin/bash

# Usage: ./exclude_from_tests.sh [-v] file [paths...]
# -v: verbose mode

verbose=0
if [ "$1" == "-v" ]; then
    verbose=1
    shift
fi

file=$1
shift

tmpfile=$(mktemp)

while read -r line; do
    exclude=0
    for path in "$@"; do
        if [[ $line == *"$path"* ]]; then
            exclude=1
            [ $verbose -eq 1 ] && echo "Excluding: $line"
            break
        fi
    done
    if [ $exclude -eq 0 ]; then
        echo "$line" >> "$tmpfile"
    fi
done < "$file"

mv "$tmpfile" "$file"

echo "Excluded paths from the file."
