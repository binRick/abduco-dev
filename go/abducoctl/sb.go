package abducoctl

import (
	"bufio"
	"fmt"
	"strings"
)

func Buffer(name string) []byte {
	cmd := SbList(name)
	r, _ := cmd.StdoutPipe()
	//pp.Println(cmd)
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
	fmt.Println(lines)
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
	return []byte(strings.Join(lines, "\n"))
}
