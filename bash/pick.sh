#!/usr/bin/env bash

FILENAME=$1
L1=$2
L2=$3

function help() {
    echo "Usage:"
    echo "- \"$0\" <filename_to_pick> [head_lines]"
    echo "- \"$0\" <filename_to_pick> [tail_lines] (use negative number \`-n\` to specify last \`n\` lines)"
    echo "- \"$0\" <filename_to_pick> [start_line] [end_line]"
}

case $# in
    1)
        cat "$FILENAME"
        ;;
    2)
        if [[ "$L1" -gt 0 ]]; then
            head -n "$L1" "$FILENAME"
        elif [[ "$L1" -lt 0 ]]; then
            tail -n "$L1" "$FILENAME"
        fi
        ;;
    3)
        if (! [[ "$L1" -gt 0 ]]) || (! [[ "$L2" -gt 0 ]]); then
            echo "Error: start line number and end line number should both be greater than 0"
            exit 1
        elif [[ "$L1" -gt "$L2" ]]; then
            echo "Error: start line number should be less than end line number"
            exit 1
        fi
        sed -n "${L1},${L2}p" "$FILENAME"
        ;;
    *)
        help
        exit 1
        ;;
esac
