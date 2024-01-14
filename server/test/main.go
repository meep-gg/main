package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	services := []string{
		// riot
		"/riot/api",
		"/riot/scylla",
		"/riot/gateway",

		// lol
		"/lol/api",
		"/lol/scylla",
		"/lol/gateway",
		"/lol/nexus",
	}

	var cmds []*exec.Cmd

	for _, service := range services {
		cmd := runService(service)
		cmds = append(cmds, cmd)
	}

	// Handle interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-c

	// Gracefully stop each service
	for _, cmd := range cmds {
		if cmd.Process != nil {
			err := cmd.Process.Signal(syscall.SIGTERM)
			if err != nil {
				log.Printf("Failed to terminate process: %s\n", err)
			}
		}
	}
}

func runService(service string) *exec.Cmd {
	cmd := exec.Command("make", "-C", ".."+service, "server")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start() // Use Start instead of Run to return immediately
	if err != nil {
		panic(err)
	}
	return cmd
}
