# tuissh

*tuissh* is a TUI (terminal User Interface) tool for managing ssh connections.

## Installation

```
go install github.com/linuxexam/tuissh@latest
```

## Configuration
tuissh shares the same config file, ~/.ssh/config, with "ssh" client.
If the *Host* value has "group1:grouper2:server" pattern, fields seperated by : are treated as groups.
