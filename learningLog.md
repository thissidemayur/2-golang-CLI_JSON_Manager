# 📘 Learnings from "CLI JSON Manager" Project

## 🧱 Golang Core Concepts

### 1. Method Receivers
- `(m *Manager)` attaches methods to a struct.
- Works like class methods but simpler.
- Pointer receivers let us modify the original struct instance.

### 2. Constructors in Go
- `NewManager(fileName string) *Manager` is an idiomatic constructor.
- Helps initialize objects safely and consistently.
- Common in standard library: `http.NewRequest()`, `json.NewEncoder()`.

### 3. File I/O
- `os.Stat()` → check if file exists.  
- `os.Create()` → create a new file.  
- `os.WriteFile()` → write data with permissions.  
- Always `defer file.Close()` after creating a file.

### 4. JSON Handling
- `json.MarshalIndent()` → convert Go → JSON.
- `json.Unmarshal()` → convert JSON → Go struct/slice.

### 5. CLI Flags
- `flag.NewFlagSet()` → define command like `add`, `list`.
- `flag.String()` / `flag.Int()` → define flags for commands.
- Use `os.Args` to read the command-line input directly.

### 6. Project Modularity
- Use packages like `internal/commands` and `internal/types`.
- Keeps `main.go` clean (entrypoint only).
- Makes testing and scaling easier.

### 7. Error Handling
- Always check `err != nil`.
- Print user-friendly error messages.
- Return early when error occurs to keep flow clean.

---

## 💡 Reflection
This project taught me how to structure a small Go application like a professional tool:
- Modular structure
- CLI design using `flag`
- CRUD logic using JSON as file storage
- Separation of logic (`commands.Manager`) and entrypoint (`main.go`)
