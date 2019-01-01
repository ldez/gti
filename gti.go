package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"syscall"
)

var version = "dev"

func main() {
	displayVersion()

	drawer := newDrawer()
	drawer.draw()

	if gitPath := os.Getenv("GIT"); len(gitPath) > 0 {
		execCommand(gitPath)
	} else {
		execCommand("git")
	}
}

func execCommand(name string) {
	cmd := exec.Command(name, os.Args[1:]...)
	output, err := cmd.CombinedOutput()

	fmt.Println(string(output))

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				os.Exit(status.ExitStatus())
			}
		} else {
			log.Println(err)
		}
	}
}

func displayVersion() {
	if raw, found := os.LookupEnv("GTI_VERBOSE"); found {
		if v, _ := strconv.ParseBool(raw); v {
			fmt.Printf("gti version %s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
		}
	}
}
