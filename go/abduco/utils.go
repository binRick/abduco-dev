package abduco

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"unicode"
)

type AbducoSession struct {
	PID     int32
	Session string
	Started string
}

func NewAbducoSession(pid interface{}, session, started string) AbducoSession {
	return AbducoSession{
		PID:     pid.(int32),
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
				ass = append(ass, AbducoSession{
					PID:     int32(pid_int),
					Session: string(cl[4]),
					Started: string(fmt.Sprintf(`%s %s`, cl[1], cl[2])),
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
