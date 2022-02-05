package abducoctl

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Buffer(name string) []string {
	cmd := exec.Command(`./../../abducoctl/get_session_buffer.sh`, name)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	return strings.Split(outStr, "\n")
	/*

		r, _ := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
		done := make(chan struct{})
		scanner := bufio.NewScanner(r)
		lines := []string{}
		go func() {
			for scanner.Scan() {
				line := scanner.Text()
				lines = append(lines, line)
			}
			done <- struct{}{}
		}()
		_ = cmd.Start()
		<-done
		cmd.Wait()
		fmt.Println(lines)*/
	/*
		select {
		case <-time.After(100 * time.Millisecond):
			if err := cmd.Process.Kill(); err != nil {
				log.Fatal("failed to kill process: ", err)
			}
			log.Println("process killed as timeout reached")
		case err := <-done:
			if err != nil {
				log.Fatalf("process finished with error = %v", err)
			}
			log.Print("process finished successfully")
		}*/
}
