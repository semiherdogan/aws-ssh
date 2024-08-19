package utils

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
)

func RunShellCommand(command string) {
	/// Disable "ctrl + c" to close application on ssh
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
