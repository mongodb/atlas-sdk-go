#!/bin/bash
# Replays the breaking-changes.sh filter logic against every historical
# breaking_changes/v*.md file and prints a comparison table.
#
# Usage:
#   ./replay-historical.sh
#
# The breaking_changes/*.md files capture the post-extractChanges output of
# gorelease for each historical major bump, so we can treat their content as
# the BREAKING_CHANGES variable and re-run the filters in isolation. Useful
# after editing filter-changes.sh to confirm which historical bumps would
# have become minor and how much noise gets stripped from the rest.
set -ueo pipefail

script_path=$(dirname "$0")
dir="$script_path/../breaking_changes"

# shellcheck source=/dev/null
source "$script_path/filter-changes.sh"

count_items() { echo "$1" | grep -c '^- ' || true; }

printf "%-18s | %-6s | %-6s | %-9s | %-6s | %s\n" "version" "before" "after" "additions" "mocks" "verdict"
printf "%-18s | %-6s | %-6s | %-9s | %-6s | %s\n" "------------------" "------" "------" "---------" "------" "-------"

for f in "$dir"/v*.md; do
    name=$(basename "$f" .md)
    BREAKING_CHANGES=$(awk '/^## SDK changes/{c=1;next} /^## API Changelog/{c=0} c' "$f")
    # shellcheck disable=SC2034 # NON_BREAKING_CHANGES is consumed by filterChanges (sourced)
    NON_BREAKING_CHANGES=""

    before=$(count_items "$BREAKING_CHANGES")
    additions=$(echo "$BREAKING_CHANGES" | grep -c ': added$' || true)
    mocks=$(echo "$BREAKING_CHANGES" | grep -c '_Call)' || true)

    filterChanges
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
