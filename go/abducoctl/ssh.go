package abducoctl

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pp "github.com/k0kubun/pp"
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

type RemoteHost struct {
	Host     string
	Hostname string
	Port     uint
	User     string
	Timeout  time.Duration
	Sessions []AbducoSession
}

func (rh *RemoteHost) ParseList(lines string) {
	host := ``
	on_active := false
	var pid_int int64 = -1
	for _, line := range strings.Split(lines, "\n") {
		if len(line) < 1 {
			continue
		}
		started := ``
		session := ``
		active := false
		if strings.HasPrefix(line, `* `) {
			active = true
			line = strings.TrimLeft(line, `* `)
		}
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
			_pi, err := strconv.ParseInt(cl[len(cl)-2], 10, 32)
			if err != nil {
				panic(err)
			}
			pid_int = _pi
			started = string(fmt.Sprintf(`%s %s`, cl[1], cl[2]))
			session = strings.Split(line, ` `)[len(strings.Split(line, ` `))-1]
		}
		if strings.Contains(line, `Active sessions (on host`) {
			host = strings.Replace(strings.Split(line, ` `)[4], `)`, ``, 1)
			on_active = true
		}
		if pid_int > 0 {
			rh.Sessions = append(rh.Sessions, AbducoSession{
				Session:  session,
				Started:  started,
				PID:      int(pid_int),
				Hostname: host,
				Active:   active,
			})
			if false {
				fmt.Fprintf(os.Stderr, `|host:%s|pid:%d|started:%s|active:%v|sess:%s|
`,
					host,
					pid_int,
					started,
					active,
					session,
				)
			}
		}
	}
}

func SSH(rh RemoteHost, cmd string) string {
	auth, err := goph.UseAgent()
	if err != nil {
		panic(err)
	}
	client, err := goph.NewConn(&goph.Config{
		User:     rh.User,
		Addr:     rh.Host,
		Port:     rh.Port,
		Auth:     auth,
		Callback: ssh.InsecureIgnoreHostKey(),
		Timeout:  rh.Timeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(ctx, rh.Timeout)
	defer cancel()
	out, err := client.RunContext(ctx, fmt.Sprintf(`/usr/bin/env PATH=/usr/local/bin:/usr/bin:/bin:/usr/local/sbin:/usr/sbin:/sbin %s`, cmd))
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
