package main

import (
	"bytes"
	"testing"
)

func Test_loadSSHConfig(t *testing.T) {
	sshConfig := `
# comment
Host aws:prod:server1
	HostName usa.linuxexam.net
	User smstong # admin
	Port 22
	IdentityFile ~/.ssh/id_rsa
	Password pass123

Host aws:prod:server2
Host aws:prod:server3
Host local:server1
Host local:server2
Host server1
Host server2
`

	entries, err := loadSSHConfig(bytes.NewReader([]byte(sshConfig)))
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 7 {
		t.Fail()
	}
}
