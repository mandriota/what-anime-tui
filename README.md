# What Anime TUI
A TUI alternative to [irevenko/what-anime-cli](https://github.com/irevenko/what-anime-cli).
Wrapper for [trace.moe](https://trace.moe) API.

## Showcase
https://github.com/mandriota/what-anime-tui/assets/62650188/fc0a4aca-0e20-43b0-a18b-e6b8b9f03694

## Installation

Download and install Go from [go.dev](https://go.dev), then enter the following command in your terminal:
```sh
go install github.com/mandriota/what-anime-tui@latest
```

You may also need to add `go/bin` directory to `PATH` environment variable.
Enter the following command in your terminal to find `go/bin` directory:
```sh
echo `go env GOPATH`/bin
```

### Using Homebrew
```sh
brew tap mandriota/mandriota
brew install what-anime-tui
```

### Using npm
```sh
npm i what-anime-tui
```

## Configuration
Config is read from `$HOME/.config/wat/wat.toml`

### Default config:
```toml
[appearance]
# Specifies background color by hex or ANSI value.
# Examples:
# background = "#0F0"
# background = "#FF006F"
# background = "6"
background = "6"
# Specifies foreground color by hex or ANSI value.
foreground = "15"

[appearance.border]
# Specifies border foreground color by hex or ANSI value.
foreground = "15"
```
