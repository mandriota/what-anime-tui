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

## Usage
```sh
what-anime-tui [-c=/path/to/your/config] [-q] [path]
```
Where
- `-c` specifies the PATH to your configuration file
- `-q`, if set, quits the program immediately after displaying first search result
- `path` can be either URL or path to local file

These flags can be omitted and you can run:
```sh
what-anime-tui
```

## Configuration
By default config is read from `~/.config/wat/wat.toml` if `-c` flag is not set.

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

# Warning: this configuration parameter is unstable and can be changed at any moment
[fetcher]
# API URL to fetch by image/gif URL, which will be passed in the link in place of {{ .Path }}
apiUrlByUrl = "https://api.trace.moe/search?anilistInfo&url={{ .Path }}"
# API URL to fetch by image/gif file, which will be passed in "image" field in multipart
apiUrlByFile = "https://api.trace.moe/search?anilistInfo"
```
