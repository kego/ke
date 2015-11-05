// exit is a demonstrator for how to capture the Ctrl+C signal and exit gracefully. Unfortunately,
// the cleanup code in the defer statements doesn't get executed, so more work is needed to wire up
// the app for graceful exit on Ctrl+C.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ioutil.WriteFile("./foo.txt", []byte("foo"), 0644)
	defer os.RemoveAll("./foo.txt")

	go func() {
		ioutil.WriteFile("./bar.txt", []byte("bar"), 0644)
		defer os.RemoveAll("./bar.txt")
		for {
			// various long running things
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Received OS interrupt - exiting.")
		os.Exit(0)
	}()

	for {
		// various long running things
	}
}
