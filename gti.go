package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	drawer := newDrawer()
	drawer.draw()

	moveToTop()

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
