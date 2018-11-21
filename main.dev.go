package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func cmdLogHandler(cmd *exec.Cmd) {

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	var stdoutBuf, stderrBuf bytes.Buffer
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

func developmentRunner() {

	go func() {
		views := exec.Command("make", "start-views")
		cmdLogHandler(views)
	}()
	server := exec.Command("go", "run", "./pkg/main.go")
	cmdLogHandler(server)

}
