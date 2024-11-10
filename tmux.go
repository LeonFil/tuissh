package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

//go:embed tmux.conf
var tmuxConfig string

var tmuxSocket string
var tmuxConfigFile string

func isTmuxRunning() bool {
	s := os.Getenv("TMUX")
	return strings.Contains(s, tmuxSocket)
}

func runTmux() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	tmuxSocket = homeDir + "/tuissh.sock"
	tmuxConfigFile = homeDir + "/tuissh.tmux"

	// is this is run by tmux
	if isTmuxRunning() {
		log.Printf("tmux is already running...")
		return
	}

	// start tmux and exit tuissh itself
	if err := os.WriteFile(tmuxConfigFile, []byte(tmuxConfig), 0600); err != nil {
		log.Printf("failed to create tmux file:%s", err)
		return
	}

	cmd := exec.Command("tmux", "-S", tmuxSocket, "-f", tmuxConfigFile, "attach-session")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Printf("run tmux failed: %s", err)
	}
	if err := cmd.Wait(); err != nil {
		log.Printf("finished with error: %s", err)
	}
	fmt.Printf("tmux started.")
	os.Exit(0)
}
