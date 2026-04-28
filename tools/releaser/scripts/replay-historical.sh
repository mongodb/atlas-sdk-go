#!/bin/bash
# Replays the breaking-changes.sh filter logic against every historical
# breaking_changes/v*.md file and prints a comparison table.
#
# Usage:
#   ./replay-historical.sh
#
# The breaking_changes/*.md files capture the post-extractChanges output of
# gorelease for each historical major bump, so we can treat their content as
# the BREAKING_CHANGES variable and re-run the filters in isolation. This is
# a diagnostic tool: handy after editing the filters in breaking-changes.sh
# to confirm which historical bumps would have become minor and how much
# noise gets stripped from the rest.
#
# IMPORTANT: the apply_filters function below must be kept in sync with the
# filters in breaking-changes.sh.
set -ueo pipefail

script_path=$(dirname "$0")
dir="$script_path/../breaking_changes"

apply_filters() {
    local input=$1
    BREAKING_CHANGES="$input"

    set +e
    ADDITIONS=$(echo "$BREAKING_CHANGES" | grep ': added$' || true)
    BREAKING_CHANGES=$(echo "$BREAKING_CHANGES" | grep -v ': added$' || true)
    BREAKING_CHANGES=$(echo "$BREAKING_CHANGES" | grep -v '_Call)' || true)
    if ! echo "$BREAKING_CHANGES" | grep -q '^- '; then
        BREAKING_CHANGES=""
    fi
    set -e
    export ADDITIONS BREAKING_CHANGES
}

count_items() { echo "$1" | grep -c '^- ' || true; }

printf "%-18s | %-6s | %-6s | %-9s | %-6s | %s\n" "version" "before" "after" "additions" "mocks" "verdict"
printf "%-18s | %-6s | %-6s | %-9s | %-6s | %s\n" "------------------" "------" "------" "---------" "------" "-------"

for f in "$dir"/v*.md; do
    name=$(basename "$f" .md)
    raw=$(awk '/^## SDK changes/{c=1;next} /^## API Changelog/{c=0} c' "$f")

    before=$(count_items "$raw")
    additions=$(echo "$raw" | grep -c ': added$' || true)
    mocks=$(echo "$raw" | grep -c '_Call)' || true)

    apply_filters "$raw"
    after=$(count_items "$BREAKING_CHANGES")

    if [ "$before" = "0" ]; then
        verdict="(no items)"
    elif [ -z "$BREAKING_CHANGES" ]; then
        verdict="WOULD BE MINOR"
    elif [ "$additions" -gt 0 ] || [ "$mocks" -gt 0 ]; then
        verdict="reduced ($before -> $after)"
    else
        verdict="unchanged"
    fi

    printf "%-18s | %-6s | %-6s | %-9s | %-6s | %s\n" "$name" "$before" "$after" "$additions" "$mocks" "$verdict"
done
