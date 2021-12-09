module basic

go 1.17

require (
	github.com/binRick/abduco-dev/go/abducoctl v0.0.0-00010101000000-000000000000
	github.com/k0kubun/pp v3.0.1+incompatible
)

require (
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
)

replace github.com/binRick/abduco-dev/go/abducoctl => ./../../abducoctl/.
