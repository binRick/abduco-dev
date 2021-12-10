DVTM_TITLE="Session $ID"
ABDUCO_CMD_file=$(mktemp)

export CMD="$(
	cat <<EOF | tr '\n' ';'
echo -e "<\$\$> \$(date +%s)" >&1
echo -e "[\$\$] \$(date +%s)" >&2
EOF
)"

BD=$(pwd)/bash
