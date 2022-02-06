package main

import (
	"io"
	"io/ioutil"
	"net"
	"time"

	//"syscall"
	"fmt"
	"os"

	//"code.google.com/p/go.crypto/ssh"

	//"code.google.com/p/go.crypto/ssh/terminal"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/crypto/ssh/terminal"
	//"github.com/fatih/color"
)

type password string

var (
	TERMINAL_TYPE  = `xterm-256color`
	TERMINAL_MODES = ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	server = os.Getenv(`HOST`)
	port   = os.Getenv(`PORT`)
	user   = os.Getenv(`USER`)
)

func PublicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}

func HasAgent() bool {
	return os.Getenv("SSH_AUTH_SOCK") != ""
}

func UseAgent() (net.Conn, error) {
	sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, fmt.Errorf("could not find ssh agent: %w", err)
	}
	return sshAgent, nil
}

func get_ssh_client_config() *ssh.ClientConfig {
	sshAgent, err := UseAgent()
	if err != nil {
		panic(err)
	}
	return &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers),
		},
	}
}

func main() {
	server = fmt.Sprintf(`%s:%s`, server, port)
	config := get_ssh_client_config()
	conn, err := ssh.Dial("tcp", server, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
		return
	}
	defer terminal.Restore(fd, state)

	w, h, err := terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	if err = session.RequestPty(TERMINAL_TYPE, h, w, TERMINAL_MODES); err != nil {
		panic(err)
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	stdinPipe, err := session.StdinPipe()
	if err != nil {
		panic(err)
	}

	if err = session.Shell(); err != nil {
		panic(err)
	}

	go func() {
		_, err = io.Copy(stdinPipe, os.Stdin)
		if err != nil {
			panic(err)
		}
		session.Close()
	}()

	go func() {
		var (
			ow = w
			oh = h
		)
		for {
			cw, ch, err := terminal.GetSize(fd)
			if err != nil {
				break
			}
			if cw != ow || ch != oh {
				err = session.WindowChange(ch, cw)
				if err != nil {
					break
				}
				ow = cw
				oh = ch
			}
			time.Sleep(time.Second)
		}
	}()
	session.Wait()
}
