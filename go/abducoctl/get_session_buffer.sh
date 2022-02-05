#!/usr/bin/env bash
set -eou pipefail
SESSION=$1
TIMEOUT=1
tf=$(mktemp)
cmd="timeout $TIMEOUT passh $(command -v abduco-sb) -r -a $SESSION | cat | bleach_text |tail -n 5 > $tf"
#> $tf 2>/dev/null"
#>&2 ansi --yellow --bg-black --italic "$cmd"

set +e
eval "$cmd"
cat $tf
#| bleach_text | cat
unlink $tf

