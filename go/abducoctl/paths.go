package abducoctl

var (
	SRC_SB_PATH        = `./../../abducoctl/files/sb-linux`
	_DST_SB_PATH       = `/usr/local/bin/sb-linux`
	DST_SB_PATH        = `/root/.bin/sb-linux`
	ABDUCO_BINARY_NAME = DST_SB_PATH
)

func init() {
	//	p, _ := Expand(SRC_SB_PATH)
	//	SRC_SB_PATH = p
	//	p, _ = Expand(DST_SB_PATH)
	//	DST_SB_PATH = p
}
