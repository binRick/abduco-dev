DVTM_TITLE="Session $ID"
ABDUCO_CMD_file=$(mktemp)

export CMD="$(
	cat <<EOF | tr '\n' ';'
echo -e "\$\$- \$(date +%s)"
EOF
)"

BD=$(pwd)/bash
