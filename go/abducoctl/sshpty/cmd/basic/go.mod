module basic

go 1.17

replace local.dev/sshpty => ./../../

require local.dev/sshpty v0.0.0-00010101000000-000000000000

require (
	golang.org/x/crypto v0.0.0-20220131195533-30dcbda58838 // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1 // indirect
)
