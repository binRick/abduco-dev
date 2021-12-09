export ACTIVE_FILE=/tmp/sess-$ID-active
export STDOUT=/tmp/sess-$ID-stdout.log
export STDERR=/tmp/sess-$ID-stderr.log
export STATUS_SOCK=/tmp/sess-$ID-dvtm-status.sock
export CMD_SOCK=/tmp/sess-$ID-dvtm-cmd.sock

export PASSH_STDOUT_LOG_FILE=${PASSH_STDOUT_LOG_FILE:-/tmp/sess-$ID-passh-stdout.log}
export PASSH_STDIN_LOG_FILE=${PASSH_STDIN_LOG_FILE:-/tmp/sess-$ID-passh-stdin.log}
