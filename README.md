# 🧹 GitClean CLI

A simple CLI tool written in Go to identify and optionally delete **stale local Git branches** that no longer exist on the remote. Keeps your local Git repo clean and manageable.

![Go Version](https://img.shields.io/badge/Go-1.24+-blue?logo=go)
![License](https://img.shields.io/badge/license-Apache--2.0-blue)
![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey)

---

## ✨ Features

- 🔍 Detects stale local branches by comparing against a remote
- 🗑️ Optionally delete stale branches via `--delete`
- 🌐 Supports custom remotes (default: `origin`)
- ⚙️ Pure Go with minimal code and no dependencies
- 🔒 Never deletes `main` or `master` branches

---

## 📦 Installation

### Option 1: Go Install

```bash
go install github.com/hadzicni/gitclean/cmd/gitclean@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

### Option 2: Manual Build

```bash
git clone https://github.com/hadzicni/gitclean.git
cd gitclean/cmd/gitclean
go build -o gitclean
```

---

## 🚀 Usage

```bash
gitclean [flags]
```

### Available Flags

| Flag           | Description                                 |
|----------------|---------------------------------------------|
| `--delete`, `-d` | Delete all stale local branches             |
| `--remote`, `-r` | Specify remote to compare against (default: `origin`) |

---

## 🔧 Examples

List stale branches:

```bash
gitclean
```

List using a custom remote:

```bash
gitclean -r upstream
```

Delete all stale branches:

```bash
gitclean -d
```

---

## 👨‍💻 Author

Made by **Nikola Hadzic**  
GitHub: [@hadzicni](https://github.com/hadzicni)

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
