#!/usr/bin/env bash
set -eou pipefail
cd $(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
export PATH=/usr/local/bin:$PATH

MODE=${1:-ls}
shift || true
ARGS="${@:-}"

BD=$(pwd)/bash
eval "$(command cat $BD/cmds.sh)"
eval "$(command cat $BD/dev.sh)"


export TIMEOUT=300
export INTERVAL=1
export ID=$(date +%s)
export ACTIVE_FILE=/tmp/sess-$ID-active
export STDOUT=/tmp/sess-$ID-stdout.log
export STDERR=/tmp/sess-$ID-stderr.log

export CMD="$(
	cat <<EOF | tr '\n' ';'
date
>&2 id
seq 1 5
>&2 seq 105 110
EOF
)"

run() {
	RUN_WHILE_LOGIC="[[ -f '$ACTIVE_FILE' ]]"
	eval_prefix="env OK123=1111"
	eval_cmd="eval \"$(
		cat <<EOF
\$(echo $(echo $CMD | base64 -w0)|base64 -d)
EOF
	)\""
	export WHILE_CMD="while $RUN_WHILE_LOGIC; do $eval_cmd; sleep $INTERVAL; done"
	#export ABDUCO_CMD="ls /proc/self/fd/${cp0[1]} || exec ${cp0[1]}>&1; ls /proc/self/fd/${cp1[1]} || exec ${cp1[1]}>&2;
	export ABDUCO_CMD="touch $ACTIVE_FILE && env timeout $TIMEOUT env bash --norc --noprofile -c '{ ${WHILE_CMD}; }'"
	DVTM_CMD="/usr/local/bin/dvtm -t \"SESSION TITLE\" -c /tmp/CMD -s /tmp/STATUS"
	ABDUCO_CMD_file=$(mktemp)
	echo -e "#!/usr/bin/env bash\n" >$ABDUCO_CMD_file
	echo -e "$ABDUCO_CMD" >$ABDUCO_CMD_file
	chmod +x $ABDUCO_CMD_file
	export ABDUCO_CMD="$DVTM_CMD $ABDUCO_CMD_file"
	# > $STDOUT 2>$STDERR"
	#>&"${cp0[1]}" 2>${cp1[1]}'"
	cmd="$(command -v abduco) -n sess-$ID"

	ansi >&2 --blue --bold "$ABDUCO_CMD"
	ansi >&2 --yellow --italic "$cmd"

	ENCODED_DVTM_INFO="ZHZ0bXxkeW5hbWljIHZpcnR1YWwgdGVybWluYWwgbWFuYWdlciwKCWFtLAoJZW8sCgltaXIsCgltc2dyLAoJeGVubCwKCWNvbG9ycyM4LAoJY29scyM4MCwKCWl0IzgsCglsaW5lcyMyNCwKCW5jdkAsCglwYWlycyM2NCwKCWFjc2M9YGBhYWZmZ2dqamtrbGxtbW5ub29wcHFxcnJzc3R0dXV2dnd3eHh5eXp6e3t8fH19fn4sCgliZWw9XkcsCglibGluaz1cRVs1bSwKCWJvbGQ9XEVbMW0sCgljaXZpcz1cRVs/MjVsLAoJY2xlYXI9XEVbSFxFWzJKLAoJY25vcm09XEVbPzI1aCwKCWNyPV5NLAoJY3NyPVxFWyVpJXAxJWQ7JXAyJWRyLAoJY3ViPVxFWyVwMSVkRCwKCWN1YjE9XkgsCgljdWQ9XEVbJXAxJWRCLAoJY3VkMT1eSiwKCWN1Zj1cRVslcDElZEMsCgljdWYxPVxFW0MsCgljdXA9XEVbJWklcDElZDslcDIlZEgsCgljdXU9XEVbJXAxJWRBLAoJY3V1MT1cRVtBLAoJZGw9XEVbJXAxJWRNLAoJZGwxPVxFW00sCgllZD1cRVtKLAoJZWw9XEVbSywKCWVsMT1cRVsxSywKCWVuYWNzPVxFKEJcRSkwLAoJaG9tZT1cRVtILAoJaHBhPVxFWyVpJXAxJWRHLAoJaHQ9XkksCglodHM9XEVILAoJaWNoPVxFWyVwMSVkQCwKCWljaDE9XEVbQCwKCWlsPVxFWyVwMSVkTCwKCWlsMT1cRVtMLAoJaW5kPV5KLAoJaXMxPVxFWz80N2xcRT1cRVs/MWwsCglpczI9XEVbclxFW21cRVsySlxFW0hcRVs/N2hcRVs/MTszOzQ7NmxcRVs0bCwKCWtEQz1cRVszJCwKCWtFTkQ9XEVbOCQsCglrSE9NPVxFWzckLAoJa0lDPVxFWzIkLAoJa0xGVD1cRVtkLAoJa05YVD1cRVs2JCwKCWtQUlY9XEVbNSQsCglrUklUPVxFW2MsCglrYTE9XEVPdywKCWthMz1cRU95LAoJa2IyPVxFT3UsCglrYnM9XDE3NywKCWtjMT1cRU9xLAoJa2MzPVxFT3MsCglrY2J0PVxFW1osCglrY3ViMT1cRVtELAoJa2N1ZDE9XEVbQiwKCWtjdWYxPVxFW0MsCglrY3V1MT1cRVtBLAoJa2RjaDE9XEVbM34sCglrZWw9XEVbOFxeLAoJa2VuZD1cRVs4fiwKCWtlbnQ9XEVPTSwKCWtmMD1cRVsyMX4sCglrZjE9XEVbMTF+LAoJa2YyPVxFWzEyfiwKCWtmMz1cRVsxM34sCglrZjQ9XEVbMTR+LAoJa2Y1PVxFWzE1fiwKCWtmNj1cRVsxN34sCglrZjc9XEVbMTh+LAoJa2Y4PVxFWzE5fiwKCWtmOT1cRVsyMH4sCglrZjEwPVxFWzIxfiwKCWtmMTE9XEVbMjN+LAoJa2YxMj1cRVsyNH4sCglrZjEzPVxFWzI1fiwKCWtmMTQ9XEVbMjZ+LAoJa2YxNT1cRVsyOH4sCglrZjE2PVxFWzI5fiwKCWtmMTc9XEVbMzF+LAoJa2YxOD1cRVszMn4sCglrZjE5PVxFWzMzfiwKCWtmMjA9XEVbMzR+LAoJa2YyMT1cRVsyMyQsCglrZjIyPVxFWzI0JAoJa2ZuZD1cRVsxfiwKCWtob21lPVxFWzd+LAoJa2ljaDE9XEVbMn4sCglraW5kPVxFW2EsCglrbW91cz1cRVtNLAoJa25wPVxFWzZ+LAoJa3BwPVxFWzV+LAoJa3JpPVxFW2IsCglrc2x0PVxFWzR+LAoJb3A9XEVbMzk7NDltLAoJcmM9XEU4LAoJcmV2PVxFWzdtLAoJcmk9XEVNLAoJcml0bT1cRVsyM20sCglybWFjcz1eTywKCXJtY3VwPVxFWzJKXEVbPzQ3bFxFOCwKCXJtaXI9XEVbNGwsCglybXNvPVxFWzI3bSwKCXJtdWw9XEVbMjRtLAoJcnMxPVxFPlxFWz8xOzM7NDs1OzZsXEVbPzdoXEVbbVxFW3JcRVsySlxFW0gsCglyczI9XEVbclxFW21cRVsySlxFW0hcRVs/N2hcRVs/MTszOzQ7NmxcRVs0bFxFPlxFWz8xMDAwbFxFWz8yNWgsCglzMGRzPVxFKEIsCglzMWRzPVxFKDAsCglzYz1cRTcsCglzZXRhYj1cRVs0JXAxJWRtLAoJc2V0YWY9XEVbMyVwMSVkbSwKCXNncj1cRVswJT8lcDYldDsxJTslPyVwMiV0OzQlOyU/JXAxJXAzJXwldDs3JTslPyVwNCV0OzUlO20lPyVwOSV0XDAxNiVlXDAxNyU7LAoJc2dyMD1cRVttXDAxNywKCXNpdG09XEVbM20sCglzbWFjcz1eTiwKCXNtY3VwPVxFN1xFWz80N2gsCglzbWlyPVxFWzRoLAoJc21zbz1cRVs3bSwKCXNtdWw9XEVbNG0sCgl0YmM9XEVbM2csCgl2cGE9XEVbJWklcDElZGQsCgpkdnRtLTI1NmNvbG9yfGR5bmFtaWMgdmlydHVhbCB0ZXJtaW5hbCBtYW5hZ2VyIHdpdGggMjU2IGNvbG9ycywKCXVzZT1kdnRtLAogICAgICAgIGNvbG9ycyMyNTYsCiAgICAgICAgcGFpcnMjMzI3NjcsCiAgICAgICAgc2V0YWI9XEVbJT8lcDElezh9JTwldDQlcDElZCVlJXAxJXsxNn0lPCV0MTAlcDElezh9JS0lZCVlNDg7NTslcDElZCU7bSwKICAgICAgICBzZXRhZj1cRVslPyVwMSV7OH0lPCV0MyVwMSVkJWUlcDElezE2fSU8JXQ5JXAxJXs4fSUtJWQlZTM4OzU7JXAxJWQlO20sCg=="

	CMD_DVTM_INFO="tic -s <(echo -e $ENCODED_DVTM_INFO|base64 -d)"
	#echo -e "$CMD_DVTM_INFO"
	eval "$cmd"
}
ls() {
	abduco -l
}
#killall abduco; ~/ABDUCO.sh; ps axfuw; abduco -l;
latest() {
	LATEST_SESSION="$(abduco -l | grep sess- | tail -n1 | tr -s ' ' | sed 's/[[:space:]]/ /g' | cut -d' ' -f7)"
	echo $LATEST_SESSION
}
kill() {
	killall abduco
}
eval "$MODE"
