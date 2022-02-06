package abducoctl

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

type RemoteHost struct {
	Host    string
	Port    uint
	User    string
	Timeout time.Duration
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
