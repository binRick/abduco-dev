#!/usr/bin/env bash
set -eou pipefail
cd $(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
export PATH=/usr/local/bin:$PATH

MODE=${1:-ls}
shift || true
ARGS="${@:-}"
export ID=$(date +%s)

for f in cmds dev paths vars env run sexpect; do eval "$(command cat $(pwd)/bash/$f.sh)"; done

a() {
	cmd="$abduco -l| grep sess-1|tr -s ' ' | sed 's/[[:space:]]/ /g' |tr -s ' '|sed 's/^[[:space:]]//g'"
	eval "$cmd"
}

pids() {
	cmd="a|cut -d' ' -f4|sort -u"
	eval "$cmd"
}

ps() {
	pids | while read -r pid; do eval "echo -n $pid\ && pstree -n $pid"; done
}

ls() {
	abduco -l
}

latest() {
	LATEST_SESSION="$($abduco -l | grep sess- | tail -n1 | tr -s ' ' | sed 's/[[:space:]]/ /g' | cut -d' ' -f7)"
	echo $LATEST_SESSION
}

kill() {
	killall abduco
}

eval "$MODE"
