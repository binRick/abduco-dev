type -P sexpect >&/dev/null || exit 1

SEXPECT_SOCKET=/tmp/abduco-sexpect-$$.sock
sexpect=$(command -v sexpect)
TIMEOUT=${TIMEOUT:-10}
TTL=${TTL:-10}
STDOUT_LOG_FILE=${STDOUT_LOG_FILE:-/tmp/sess-$ID-sexpect-stdout.log}
export SEXPECT_SOCKFILE=$SEXPECT_SOCKET

sexpect_wrap() {
	CMD="$@"
	SEXPECT_CMD_file=$(mktemp)
	echo -e "#!/usr/bin/env bash\n" >$SEXPECT_CMD_file
	echo -e "$CMD" >>$SEXPECT_CMD_file
	chmod +x $SEXPECT_CMD_file

	local cmd="$sexpect -s $SEXPECT_SOCKET spawn -logfile $STDOUT_LOG_FILE -nonblock -t 10 -ttl 600  $SEXPECT_CMD_file"
	echo -e "$cmd"
}
