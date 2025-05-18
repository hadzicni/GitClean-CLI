# 🧹 GitClean CLI

A simple and powerful command-line tool to clean up your local Git repositories by removing untracked files and directories — safely and efficiently.

![Go Version](https://img.shields.io/badge/Go-1.20+-blue?logo=go)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey)

---

## ✨ Features

- 🗑️ Cleans untracked files and directories (similar to `git clean -fd`)
- 🕵️ Dry-run mode for safe previews before deleting
- 📁 Target specific Git repositories
- 🧪 Built with the Go standard library — no external dependencies
- 🧼 Useful for monorepos or cleaning multiple repositories in bulk

---

## 📦 Installation

### Option 1: Go Install (recommended)

```bash
go install github.com/hadzicni/gitclean-cli@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

### Option 2: Manual Build

```bash
git clone https://github.com/hadzicni/gitclean-cli.git
cd gitclean-cli
go build -o gitclean ./cmd/gitclean
```

---

## 🚀 Usage

```bash
gitclean [flags]
```

### Available Flags

| Flag              | Description                                      | Example                          |
|-------------------|--------------------------------------------------|----------------------------------|
| `--path`, `-p`    | Path to the Git repo (default: current directory) | `gitclean -p ./my-project`       |
| `--dry-run`, `-d` | Show what would be deleted without deleting      | `gitclean -d`                    |
| `--force`, `-f`   | Actually perform deletion                        | `gitclean -f`                    |
| `--help`, `-h`    | Show help message                                | `gitclean -h`                    |

---

## 🔧 Examples

Dry run on current directory:

```bash
gitclean
```

Clean a specific path (force mode):

```bash
gitclean -p ./some-repo -f
```

Preview changes on a specific directory:

```bash
gitclean -p ./some-repo -d
```

---

## 🧪 Development

Run the CLI locally:

```bash
go run ./cmd/gitclean
```

Run tests:

```bash
go test ./...
```

---

## 📁 Project Structure

```
.
├── cmd/
│   └── gitclean/       # CLI logic and argument parsing
├── internal/           # Reusable internal packages
│   └── cleaner/        # Cleaning logic
├── go.mod              # Go module definition
├── LICENSE             # License file
└── README.md           # Project documentation
```

---

## 👨‍💻 Author

Made with ❤️ by **Nikola Hadzic**

- GitHub: [@hadzicni](https://github.com/hadzicni)

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
