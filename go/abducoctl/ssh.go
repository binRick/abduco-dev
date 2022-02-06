package abducoctl

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	pp "github.com/k0kubun/pp"
	sftp "github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"local.dev/goph"
)

var (
	REMOTE_PATH = `~/.bin:/usr/local/bin:/usr/bin:/bin:/usr/local/sbin:/usr/sbin:/sbin`
)

type RemoteHost struct {
	Host     string
	Hostname string
	Name     string
	Port     uint
	User     string
	Timeout  time.Duration
	Sessions []AbducoSession
	OS       string
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

func NewRemoteCommand(cmd string) string {
	return fmt.Sprintf(`/usr/bin/env PATH=%s %s`, REMOTE_PATH, cmd)
}

func GetSSHClient(rh RemoteHost) *goph.Client {
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
	return client

}
func GetSftpClient(client *goph.Client) *sftp.Client {
	sftp, err := client.NewSftp()
	if err != nil {
		panic(err)
	}
	return sftp
}

func SftpPutFileContent(rh RemoteHost, f string, content []byte) {
	client := GetSSHClient(rh)
	defer client.Close()
	sftp := GetSftpClient(client)
	defer sftp.Close()
	file, err := sftp.Create(f)
	if err != nil {
		panic(err)
	}
	file.Write(content)
	file.Close()
}

func UploadSB(rh RemoteHost) {
	SftpPutFile(rh, SourcePath(rh), DestPath(rh), 0700)
}

func SftpPutFile(rh RemoteHost, local_file, remote_file string, mode os.FileMode) {
	lf, err := os.Stat(local_file)
	if err != nil {
		panic(err)
	}
	client := GetSSHClient(rh)
	defer client.Close()
	sftp := GetSftpClient(client)
	defer sftp.Close()
	remote_dir := filepath.Dir(remote_file)
	if rf, err := sftp.Lstat(remote_file); (err != nil) || (rf.Size() != lf.Size()) {
		if _, lerr := sftp.Lstat(remote_dir); lerr != nil {
			merr := sftp.Mkdir(remote_dir)
			if merr != nil {
				panic(merr)
			}
		}
		err = client.Upload(local_file, remote_file)
		if err != nil {
			panic(err)
		}
	}
	fi, err := sftp.Lstat(remote_file)
	if err != nil {
		panic(err)
	}
	if fi.Mode() != mode {
		err = sftp.Chmod(remote_file, mode)
		if err != nil {
			panic(err)
		}
	}
}

func NormalizeRemoteHost(rh RemoteHost) {
	//SftpPutFileContent(rh, `/tmp/tt12345`, []byte(`xxxxxxxxxxxxx`))
	//SftpPutFile(rh, `/tmp/src`, `/tmp/dest`, 0700)
	UploadSB(rh)
}

func SSH(rh RemoteHost, cmd string) string {
	client := GetSSHClient(rh)
	defer client.Close()
	ctx, cancel := context.WithTimeout(ctx, rh.Timeout)
	defer cancel()
	out, err := client.RunContext(ctx, NewRemoteCommand(cmd))
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
