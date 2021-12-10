type -P sexpect >&/dev/null || exit 1

SEXPECT_SOCKET=/tmp/abduco-sexpect-$$.sock
sexpect=$(command -v sexpect)
TIMEOUT=${TIMEOUT:-10}
TTL=${TTL:-10}
TMPDIR=${TMPDIR:-/tmp}
STDOUT_LOG_FILE=${STDOUT_LOG_FILE:-/tmp/sess-$ID-sexpect-stdout.log}
STDERR_LOG_FILE=${STDERR_LOG_FILE:-/tmp/sess-$ID-sexpect-stderr.log}
export SEXPECT_SOCKFILE=$SEXPECT_SOCKET

sexpect_wrap() {
	CMD="$@"
	SEXPECT_CMD_file=$TMPDIR/sexpect-wrapper-script-$$.sh
	echo -e "#!/usr/bin/env bash" >$SEXPECT_CMD_file
	echo -e "sh -c '$CMD' 2> $STDERR_LOG_FILE" >>$SEXPECT_CMD_file
	chmod +x $SEXPECT_CMD_file
	CMD=$SEXPECT_CMD_file
	#CMD="sh -c '$CMD'"
	local cmd="$sexpect -s $SEXPECT_SOCKET spawn -logfile $STDOUT_LOG_FILE -nonblock -t 10 -ttl 600 $CMD"
	echo -e "$cmd"
}
