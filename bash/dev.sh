cp() {
	coproc cp0 (while :; do
		read -r input
		echo "-${input}" | tee $STDOUT
	done)
	coproc cp1 (while :; do
		read -r input
		echo ">${input}" | tee $STDERR
	done)
	echo "The PID of the cp0 coprocess is ${cp0_PID}"
	echo "The PID of the cp1 coprocess is ${cp1_PID}"
}
