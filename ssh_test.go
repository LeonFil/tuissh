package main

import "testing"

func Test_sshShell(t *testing.T) {
	_ = &SSHEntry{
		serverName: "usa.linuxexam.net",
		port:       22,
		user:       "smstong",
		keyFile:    "/Users/jonathan/.ssh/id_rsa",
	}

	//sshShell(entry)
}
