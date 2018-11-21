package main

import (
	"fmt"
	"os/exec"

	// "github.com/zapkub/gorogoso"
	"gorogoso"
)

func developmentRunner() {

	serverPID := gorogoso.Monit("pkg/**/*.go", "pkg/main.go")
	// wait until server boots
	<-serverPID
	views := exec.Command("make", "start-views")
	viewPID := make(chan int)
	go gorogoso.CMDLogHandler(viewPID, views)

	for {
		select {
		case pid := <-serverPID:
			fmt.Printf("\n\n Server PID: %d \n\n", pid)
		}
	}

}
