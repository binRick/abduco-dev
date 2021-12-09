package abducoctl

import (
	"bufio"
	"fmt"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	gops "github.com/mitchellh/go-ps"
)

type AbducoSession struct {
	PPID        int
	PID         int
	PIDs        []int
	Threads     int
	Session     string
	Executable  string
	Executables []string
	Started     string
}

func NewAbducoSession(pid interface{}, session, started string) AbducoSession {
	return AbducoSession{
		PID:     pid.(int),
		Session: string(session),
		Started: string(started),
	}
}

func TabToSpace(input string) string {
	var result []string
	for _, i := range input {
		switch {
		case unicode.IsSpace(i):
			result = append(result, " ")
		case !unicode.IsSpace(i):
			result = append(result, string(i))
		}
	}
	return strings.Join(result, "")
}

type AbducoSessions struct {
	Sessions []AbducoSession
}

func List() ([]AbducoSession, error) {
	var ass []AbducoSession
	cmd := exec.Command("env", "abduco", "-l")
	r, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	done := make(chan struct{})
	scanner := bufio.NewScanner(r)
	on_active := false
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			line = TabToSpace(line)
			line = strings.Trim(line, ` `)
			if len(strings.Split(line, ` `)) < 1 {
				continue
			}
			spl := strings.Split(line, ` `)
			if on_active {
				cl := []string{}
				for _, c := range spl {
					c = strings.TrimSpace(c)
					if len(c) > 0 {
						cl = append(cl, c)
					}
				}
				pid_int, err := strconv.ParseInt(cl[3], 10, 32)
				if err != nil {
					panic(err)
				}
				p, err := gops.FindProcess(int(pid_int))
				if err != nil {
					panic(err)
				}
				P, err := getRelevantProcs(int(pid_int))
				pids := []int{}
				threads := 0
				executables := []string{}
				for _, _p := range P {
					pids = append(pids, _p.PID)
					executables = append(executables, _p.Comm)
					threads += _p.NumThreads
				}
				if err == nil {

				}
				//				pp.Println(p, P)
				ass = append(ass, AbducoSession{
					PID:         int(pid_int),
					PPID:        int(p.PPid()),
					PIDs:        pids,
					Threads:     threads,
					Executables: executables,
					Executable:  p.Executable(),
					Session:     string(cl[4]),
					Started:     string(fmt.Sprintf(`%s %s`, cl[1], cl[2])),
				})
			} else {
				if spl[0] == `Active` {
					on_active = true
				}
			}
		}
		done <- struct{}{}
	}()
	_ = cmd.Start()
	<-done
	_ = cmd.Wait()
	return ass, nil
}

func ReverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
