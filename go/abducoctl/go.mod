module abducoctl

go 1.17

require (
	github.com/AlecAivazis/survey/v2 v2.3.2
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/charmbracelet/lipgloss v0.4.0
	github.com/creack/pty v1.1.17
	github.com/fatih/color v1.13.0
	github.com/google/uuid v1.3.0
	github.com/gosuri/uitable v0.0.4
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/leaanthony/go-ansi-parser v1.2.0
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d
	github.com/mitchellh/go-ps v1.0.0
	github.com/pkg/sftp v1.13.4
	github.com/prometheus/procfs v0.7.3
	github.com/shirou/gopsutil/v3 v3.22.1
	github.com/wayneashleyberry/terminal-dimensions v1.1.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
	local.dev/go-fuzzyfinder v0.0.0-00010101000000-000000000000
	local.dev/goph v0.0.0-00010101000000-000000000000
)

require (
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/gdamore/tcell/v2 v2.4.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/ktr0731/go-fuzzyfinder v0.5.1 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/melbahja/goph v1.3.0 // indirect
	github.com/muesli/reflow v0.2.1-0.20210115123740-9e1d0d53df68 // indirect
	github.com/muesli/termenv v0.9.0 // indirect
	github.com/nsf/termbox-go v0.0.0-20201124104050-ed494de23a00 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/tklauser/numcpus v0.3.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	golang.org/x/sys v0.0.0-20220204135822-1c1b9b1eba6a // indirect
	golang.org/x/text v0.3.6 // indirect
)

replace local.dev/go-fuzzyfinder => ./go-fuzzyfinder

replace local.dev/goph => ./goph
