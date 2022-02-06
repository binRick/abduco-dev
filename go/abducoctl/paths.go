package abducoctl

import "fmt"

var BIN = `sb`
var (
	SRC_SB_PATH  = `./../../abducoctl/files/sb-%s`
	_DST_SB_PATH = `/usr/local/bin/sb-linux`
	DST_SB_PATH  = `/root/.bin/sb`
)

func DestPath(rh RemoteHost) string {
	prefix := `/tmp`
	switch rh.OS {
	case "linux":
		switch rh.User {
		case "root":
			prefix = `/root`
		default:
			prefix = fmt.Sprintf(`/home/%s`, rh.User)
		}
	case "darwin":
		prefix = fmt.Sprintf(`/Users/%s`, rh.User)
	}
	return fmt.Sprintf(`%s/%s`, prefix, BIN)
}

func SourcePath(rh RemoteHost) string {
	return fmt.Sprintf(SRC_SB_PATH, rh.OS)
}

func init() {
	//	p, _ := Expand(SRC_SB_PATH)
	//	SRC_SB_PATH = p
	//	p, _ = Expand(DST_SB_PATH)
	//	DST_SB_PATH = p
}
