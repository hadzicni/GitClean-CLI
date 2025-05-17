# GitClean

A simple CLI tool to find and optionally delete stale local Git branches that no longer exist on the remote.

## 🚀 Usage

```bash
gitclean             # list stale branches
gitclean -d          # delete them locally
gitclean -r upstream # use a different remote
```

## 🧾 Flags

| Flag | Description                   |
| ---- | ----------------------------- |
| `-d` | Delete stale branches locally |
| `-r` | Remote name (default: origin) |

## 🛠 Build

```bash
go build -o gitclean.exe ./cmd/gitclean
```

## 📄 License

MIT License
