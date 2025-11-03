# 🧰 CLI JSON Manager

A lightweight command-line tool written in **Go** to manage JSON records (Add, List, Update, Delete) — built to demonstrate CLI design, modular architecture, and production-style logging.

---

## 🚀 Features

- 🗂️ Manage JSON records (CRUD operations)
- ⚙️ Configurable file paths using flags (`--file`)
- 🧠 Environment-aware logging (`APP_ENV=dev` or `APP_ENV=prod`)
- 🪵 Structured logging using Go’s `slog`
- 🧩 Modular architecture (clean folder structure)
- 🧾 Supports Linux, macOS, and Windows

---

## 🏗️ Project Structure
```
cli-json-manager/
├── cmd/
│ └── cli-json-manager/
│ └── main.go # CLI entrypoint
├── internal/
│ └── cli/
│ └── commands/ # CRUD logic & manager
│ ├── crud.go
│ ├── manager.go
│ └── file_utils.go
├── logger/
│ └── logger.go # Environment-based logger setup
├── logs/
│ └── app.log # Prod logs stored here
├── data.json # Default data store
├── go.mod / go.sum
├── Makefile # (Optional) build shortcuts
└── README.md
```

---

## ⚙️ Installation

### 1️⃣ Clone the repository

```bash
git clone https://github.com/thissidemayur/cli-json-manager.git
cd cli-json-manager
```

### 2️⃣ Build the binary
go build -o /cli-json-manager ./cmd/cli-json-manager

### 3️⃣ (Optional) Move to global path
`sudo mv cli-json-manager /usr/local/bin/`


# 🧠 Usage
### General Syntax
` cli-json-manager [--file <fileName>] <command> [options] `
### 🧩 Commands

| Command  | Description         | Example                                                |
| -------- | ------------------- | ------------------------------------------------------ |
| `add`    | Add a new record    | `cli-json-manager add --name "Mayur"`                  |
| `list`   | List all records    | `cli-json-manager list`                                |
| `delete` | Delete record by ID | `cli-json-manager delete --id 3`                       |
| `update` | Update record name  | `cli-json-manager update --id 2 --name "Bhawaniputra"` |

---

### 🧱 Environment Modes

| Mode        | Variable       | Description                                        |
| ------------ | -------------- | -------------------------------------------------- |
| Development  | `APP_ENV=dev`  | Shows debug & info logs                            |
| Production   | `APP_ENV=prod` | Logs only errors and writes them to `logs/app.log` |

# Example:
`APP_ENV=prod cli-json-manager add --name "TaraPutra"`

## 🧪 Development Notes
- Built with Go 1.23+
- No external dependencies except standard library
- Followed clean architecture principles for modular design

## 📘 Learning Goals
This project was built to learn:
- How Go CLI tools work
- Handling subcommands and flags
- File operations (read/write JSON)
- Environment-based logging
- Structuring real-world Go projects
## 💡 Future Enhancements
- Add config file support (--config path)
- Implement search/filter feature
- Add colored output & progress bars
- Cross-platform installer scripts