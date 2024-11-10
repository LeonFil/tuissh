# tuissh

```
tuissh0:31- 1:tuissh*                                                           
╔════available servers═══╗┌─────────────────────server info────────────────────┐
║servers                 ║│server name: 192.168.0.131                          │
║├──w0                   ║│port: 22                                            │
║├──31                   ║│user: smstong                                       │
║├──30                   ║│private key file: /Users/jonathan/.ssh/id_rsa       │
║├──20                   ║│                                                    │
║├──19                   ║│                                                    │
║├──x                    ║│                                                    │
║├──homelab              ║│                                                    │
║│  ├──dev               ║│                                                    │
║│  │  ├──31             ║└────────────────────────────────────────────────────┘
║│  │  └──19             ║┌───────────────────────message──────────────────────┐
║│  └──prod              ║│2024-11-10 15:51:45 connecting to 31...             │
║│     ├──30             ║│                                                    │
║│     └──20             ║│                                                    │
║└──linuxexam            ║│                                                    │
║   └──usa               ║│                                                    │
║      └──x              ║│                                                    │
║                        ║│                                                    │
║                        ║│                                                    │
║                        ║│                                                    │
║                        ║│                                                    │
╚════════════════════════╝└────────────────────────────────────────────────────┘
```

*tuissh* is a TUI (terminal User Interface) tool for managing ssh connections.

## Installation

```
go install github.com/linuxexam/tuissh@latest
```

## Configuration
tuissh shares the same config file, ~/.ssh/config, with "ssh" client.
If the *Host* value has "group1:grouper2:server" pattern, fields seperated by : are treated as groups.
