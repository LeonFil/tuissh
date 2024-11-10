package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// entry of ssh.config
type SSHEntry struct {
	displayName string
	serverName  string
	port        int
	user        string
	password    string
	keyFile     string
}

func (entry *SSHEntry) String() string {
	out := fmt.Sprintf("server name: %s\nport: %d\nuser: %s\nprivate key file: %s\n",
		entry.serverName, entry.port, entry.user, entry.keyFile)
	return out
}

func parseSSHEntry(lines []string) (*SSHEntry, error) {
	var defaultKeyFile string // default keyFile
	homeDir, err := os.UserHomeDir()
	if err == nil {
		defaultKeyFile = homeDir + "/.ssh/id_rsa"
	}

	entry := SSHEntry{
		port:    22,
		keyFile: defaultKeyFile,
	}
	for _, line := range lines {
		line = strings.Split(line, "#")[0] // remove comment part
		line = strings.TrimSpace(line)     // trim spaces

		// Host
		if strings.HasPrefix(line, "Host ") {
			parts := strings.Fields(line)
			if len(parts) < 2 {
				continue
			}
			entry.displayName = parts[1]
		}

		// HostName
		if strings.HasPrefix(line, "HostName ") {
			parts := strings.Fields(line)
			if len(parts) < 2 {
				continue
			}
			entry.serverName = parts[1]
		}

		// Port
		if strings.HasPrefix(line, "Port ") {
			parts := strings.Fields(line)
			if len(parts) < 2 {
				continue
			}
			i, err := strconv.Atoi(parts[1])
			if err != nil {
				continue
			}
			entry.port = i
		}

		// IdentityFile
		if strings.HasPrefix(line, "IdentityFile ") {
			parts := strings.Fields(line)
			if len(parts) < 2 {
				continue
			}
			entry.keyFile = parts[1]
		}

		// User
		if strings.HasPrefix(line, "User ") {
			parts := strings.Fields(line)
			if len(parts) < 2 {
				continue
			}
			entry.user = parts[1]
		}
		// Password [ ext ]
		if strings.HasPrefix(line, "Password ") {
			parts := strings.Fields(line)
			if len(parts) < 2 {
				continue
			}
			entry.password = parts[1]
		}
	}
	return &entry, nil
}

func loadSSHConfig(r io.Reader) ([]*SSHEntry, error) {
	var entries []*SSHEntry
	scanner := bufio.NewScanner(r)
	var entryLines []string
	inHost := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "Host ") {
			// until the first Host
			if inHost {
				entry, err := parseSSHEntry(entryLines)
				if err != nil {
					return entries, err
				}
				entries = append(entries, entry)
			}
			entryLines = entryLines[:0]
			inHost = true
		}
		entryLines = append(entryLines, line)
	}

	if err := scanner.Err(); err != nil {
		return entries, err
	}
	entry, err := parseSSHEntry(entryLines)
	if err != nil {
		return entries, err
	}
	entries = append(entries, entry)

	return entries, nil
}

func loadSSHConfigFile(filePath string) ([]*SSHEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return loadSSHConfig(file)
}
