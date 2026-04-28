#!/bin/bash
# Sourceable: defines filterChanges, which post-processes the BREAKING_CHANGES
# and NON_BREAKING_CHANGES variables produced by extractChanges to drop
# false-positive incompatibilities reported by gorelease.
#
# In/out variables (mutated in place):
#   BREAKING_CHANGES      - lines under "### incompatible changes"
#   NON_BREAKING_CHANGES  - lines under "### compatible changes"
#
# Sourced by breaking-changes.sh (release pipeline) and replay-historical.sh
# (diagnostic tool) so the filtering logic stays in one place.

filterChanges() {
    # Filter 1: additions (`X: added`) are not breaking for SDK consumers.
    # They only break implementors of exported interfaces, and the only
    # implementors are mockery-generated mocks, which are regenerated
    # alongside the SDK. Move these lines into the compatible changes section.
    local additions
    additions=$(echo "$BREAKING_CHANGES" | grep ': added$' || true)
    BREAKING_CHANGES=$(echo "$BREAKING_CHANGES" | grep -v ': added$' || true)
    if [ -n "$additions" ]; then
        if [ -z "$NON_BREAKING_CHANGES" ]; then
            NON_BREAKING_CHANGES="### compatible changes"
        fi
        NON_BREAKING_CHANGES="${NON_BREAKING_CHANGES}
### additions (moved from incompatible)
${additions}"
    fi

    # Filter 2: mockery-generated `_Call` types (e.g. `(*FooApi_Bar_Call).Run`)
    # are internal mock plumbing, not part of the consumer-facing surface.
    # Drop them from breaking changes entirely.
    BREAKING_CHANGES=$(echo "$BREAKING_CHANGES" | grep -v '_Call)' || true)

    # If only the section header is left after filtering, treat as no breaking
    # changes so the release pipeline picks a minor bump.
    if ! echo "$BREAKING_CHANGES" | grep -q '^- '; then
        BREAKING_CHANGES=""
    fi
}
