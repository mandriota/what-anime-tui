# What Anime TUI
A TUI alternative to https://github.com/irevenko/what-anime-cli.

## Showcase
https://github.com/mandriota/what-anime-tui/assets/62650188/92ccf79f-d246-4308-a5e0-1764e6c5c578

## Installation

Download and install Go from https://go.dev, then enter the following command in your terminal:
```
go install github.com/mandriota/what-anime-tui@latest
```

You may also need to add go/bin directory to PATH environment variable.
Enter the following command in your terminal to find go/bin directory:
```
echo `go env GOPATH`/bin
```

## Configuration
Config is read from `$HOME/.config/wat/wat.toml`

### Default config:
```toml
[appearance]
# Background ANSI Color. Must be from 0 to 15.
background = 6
```