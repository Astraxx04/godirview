# GoDirView

GoDirView is a lightweight CLI tool written in Go that displays the directory structure visually, with **folder/file icons** and **colored output**. It also supports JSON output for automation or scripting.

### Features

  * Visual tree view of directories
  * Folder (📁) and file (📄) icons
  * Colorized output: folders in **blue**, files in **white**
  * Option to show hidden files
  * Depth control (`-L`) to limit recursion
  * JSON output (`--json`) for programmatic use
  * Cross-platform (Linux, macOS, Windows with Go installed)

-----

### Installation

#### From Source (Recommended for Latest Changes)

Clone the repository:

```sh
git clone https://github.com/Astraxx04/godirview.git
cd godirview
```

Install the CLI locally:

```sh
go install ./cmd/godirview
```

> **Note:** Make sure `$GOPATH/bin` or `~/go/bin` is in your `PATH`.

#### From GitHub (Latest Released Version)

```sh
go install github.com/Astraxx04/godirview/cmd/godirview@latest
```

If you want to bypass Go module caching and ensure the latest commit is installed:

```sh
GOPROXY=direct go install github.com/Astraxx04/godirview/cmd/godirview@latest
```

-----

### Usage

```sh
godirview [path] [flags]
```

#### Flags

  * `-L <number>`: Descend only this many directory levels deep (default: unlimited)
  * `-a`: Show hidden files
  * `--json`: Output directory structure as JSON instead of a tree

#### Examples

Display the current directory with icons:

```sh
godirview .
```

Limit depth to 2 levels:

```sh
godirview . -L 2
```

Show hidden files:

```sh
godirview . -a
```

Output as JSON:

```sh
godirview . --json
```

-----

### Screenshot

```
├── 📁 cmd
│   └── 📁 godirview
│       └── 📄 main.go
├── 📄 go.mod
├── 📄 go.sum
└── 📁 internal
    └── 📁 visualizer
        └── 📄 visualizer.go
```

-----

### Dependencies

  * [`fatih/color`](https://www.google.com/search?q=%5Bhttps://github.com/fatih/color%5D\(https://github.com/fatih/color\)) – For colored CLI output
  * Standard Go packages: `os`, `filepath`, `strings`, `encoding/json`

-----

### Development

Clone the repo and install dependencies:

```sh
git clone https://github.com/Astraxx04/godirview.git
cd godirview
go mod tidy
```

Build locally:

```sh
go install ./cmd/godirview
```

Test your changes:

```sh
godirview .
```

-----

### License

[MIT License](https://github.com/Astraxx04/godirview/blob/main/LICENSE) © 2025 [Gagan S.](https://github.com/Astraxx04)