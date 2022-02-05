package abducoctl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"github.com/k0kubun/pp"
	gops "github.com/mitchellh/go-ps"
	ps "github.com/shirou/gopsutil/v3/process"
)

const (
	ABDUCO_BINARY_NAME = `abduco-sb`
	DEBUG_MODE         = false
)

type AbducoSessions struct {
	Sessions []AbducoSession
}

type AbducoSession struct {
	PPID        int
	PID         int
	PIDs        []int
	Threads     int
	Session     string
	Executable  string
	Executables []string
	//	Environ        []string
	Started          string
	StartedTime      time.Time
	StartedAgo       int64
	StartedTimestamp int64
	Duration         time.Duration
	Username         string
	Cmdline          string
	Cwd              string
	Status           []string
	ConnectionsQty   int32
	OpenFilesQty     int32
	Terminal         string
	CreateTime       int64
	CPUPercent       float64
	MemoryPercent    float32
}

func NewAbducoSession(pid interface{}, session, started string) AbducoSession {
	return AbducoSession{
		PID:     pid.(int),
		Session: string(session),
		Started: string(started),
	}
}

func NewSessionNameString() string {
	return strings.Split(uuid.NewString(), `-`)[0]
}

func NewNameString() string {
	return uuid.NewString()
}

func Path() string {
	p, err := exec.LookPath(ABDUCO_BINARY_NAME)
	if err != nil {
		panic(err)
	}
	return p
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
func Exists(name string) bool {
	for _, n := range Names() {
		if n == name {
			return true
		}
	}
	return false
}
func Names() []string {
	names := []string{}
	list, _ := List()
	for _, s := range list {
		names = append(names, s.Session)
	}
	return names
}
func GetPids() ([]int, error) {
	pids := []int{}
	return pids, nil
}

func get_cmd() *exec.Cmd {
	c := exec.Command("env", ABDUCO_BINARY_NAME, "-l")
	return c
}

func Ps() {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(Names(), "\n"))
}

func PIDs() string {
	pids, e := GetPids()
	if e != nil {
		panic(e)
	}
	s := []string{}
	for p, _ := range pids {
		s = append(s, fmt.Sprintf(`%d`, p))
	}
	return strings.Join(s, "\n")
}

func JSON() string {
	l, e := List()
	if e != nil {
		panic(e)
	}
	dat, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s", string(dat))
}

func List() ([]AbducoSession, error) {
	var ass []AbducoSession
	cmd := get_cmd()
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
				if DEBUG_MODE {
					pp.Fprintf(os.Stderr, "CL: %s\n", cl)
				}
				if len(cl) != 5 {
					continue
				}
				pid_int, err := strconv.ParseInt(cl[len(cl)-2], 10, 32)
				if err != nil {
					panic(err)
				}
				proc, err := ps.NewProcess(int32(pid_int))
				if err != nil {
					panic(err)
				}

				p, err := gops.FindProcess(int(pid_int))
				if err != nil {
					panic(err)
				}
				pp.Fprintf(os.Stderr, "%s\n", p)
				pp.Fprintf(os.Stderr, "%s\n", proc)
				cmdl, _ := proc.Cmdline()
				cwd, _ := proc.Cwd()
				st, _ := proc.Status()
				as := AbducoSession{
					PID:        int(pid_int),
					PPID:       int(p.PPid()),
					Session:    string(cl[len(cl)-1]),
					Started:    string(fmt.Sprintf(`%s %s`, cl[1], cl[2])),
					Executable: p.Executable(),
					Cmdline:    cmdl,
					Cwd:        cwd,
					Status:     st,
				}
				tm, e := dateparse.ParseLocal(as.Started)
				if e != nil {
					panic(e)
				}
				as.StartedTime = tm
				as.StartedTimestamp = tm.Unix()
				as.StartedAgo = time.Now().Unix() - as.StartedTimestamp
				as.Duration = time.Since(as.StartedTime)
				ass = append(ass, as)
				continue
				//			os.Exit(1)
				/*
					p, err := gops.FindProcess(int(pid_int))
					if err != nil {
						panic(err)
					}
					P, err := getRelevantProcs(int(pid_int))
					if err == nil {
						pids := []int{}
						threads := 0
						executables := []string{}
						for _, _p := range P {
							pids = append(pids, _p.PID)
							executables = append(executables, _p.Comm)
							threads += _p.NumThreads
						}

						proc, err := ps.NewProcess(int32(pid_int))
						if err != nil {
							panic(err)
						}
						if DEBUG_MODE {
							pp.Fprintf(os.Stderr, "C:    %s\n", cl)
						}
						ct, _ := proc.CreateTime()
						mp, _ := proc.MemoryPercent()
						cp, _ := proc.CPUPercent()
						cmdl, _ := proc.Cmdline()
						cwd, _ := proc.Cwd()
						st, _ := proc.Status()
						term, _ := proc.Terminal()
						conns, _ := proc.Connections()
						//					env, _ := proc.Environ()
						un, _ := proc.Username()
						of, _ := proc.OpenFiles()
						ass = append(ass, AbducoSession{
							PID:           int(pid_int),
							PPID:          int(p.PPid()),
							PIDs:          pids,
							Threads:       threads,
							CreateTime:    ct,
							Executables:   executables,
							Cmdline:       cmdl,
							Executable:    p.Executable(),
							MemoryPercent: mp,
							CPUPercent:    cp,
							Cwd:           cwd,
							Terminal:      term,
							Status:        st,
							Username:      un,
							//						Environ:        env,
							OpenFilesQty:   int32(len(of)),
							ConnectionsQty: int32(len(conns)),
							Session:        string(cl[len(cl)-1]),
							Started:        string(fmt.Sprintf(`%s %s`, cl[1], cl[2])),
						})
					}*/
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
	if DEBUG_MODE {
		pp.Fprintf(os.Stderr, "%s\n", ass)
	}
	for _, s := range ass {
		if DEBUG_MODE {
			fmt.Fprintf(os.Stdout, "%s <%d> %s\n", s.Started, s.PID, s.Session)
		}
		//fmt.Fprintf(os.Stderr, "%s <%d> %s\n", s.Started, s.PID, s.Session)
	}
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
