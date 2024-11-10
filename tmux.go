package main

import (
	_ "embed"
	"log"
	"os"
	"os/exec"
	"strings"
)

//go:embed tmux.conf
var tmuxConfig string

var tmuxSocket = "/tmp/tuissh.sock"
var tmuxConfigFile = "/tmp/tuissh.tmux"

func isTmuxRunning() bool {
	s := os.Getenv("TMUX")
	return strings.Contains(s, tmuxSocket)
}
func runTmux() {
	if isTmuxRunning() {
		log.Printf("tmux is already running...")
		return
	}
	if err := os.WriteFile(tmuxConfigFile, []byte(tmuxConfig), 0600); err != nil {
		log.Fatalf("failed to create tmux file:%s", err)
	}

	cmd := exec.Command("tmux", "-S", tmuxSocket, "-f", tmuxConfigFile, "attach-session")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatalf("run tmux failed: %s", err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatalf("finished with error: %s", err)
	}
}
