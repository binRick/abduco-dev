DVTM_TITLE="Session $ID"
ABDUCO_CMD_file=$(mktemp)

export CMD="$(
	cat <<EOF | tr '\n' ';'
date
>&2 id
seq 1 5
>&2 seq 105 110
EOF
)"

BD=$(pwd)/bash
