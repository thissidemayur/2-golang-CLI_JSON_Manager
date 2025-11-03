# 📘 Learnings from "CLI JSON Manager" Project

## 🧱 Golang Core Concepts

### 1. Method Receivers
- `(m *Manager)` attaches methods to a struct.
- Similar to class methods in OOP.
- Pointer receivers let methods modify the original struct instance.

### 2. Constructors in Go
- `NewManager(fileName string) *Manager` is an idiomatic constructor.
- Ensures safe, consistent initialization.
- Common in stdlib: `http.NewRequest()`, `json.NewEncoder()`.

### 3. File I/O
- `os.Stat()` → check if file exists.  
- `os.Create()` → create a new file.  
- `os.WriteFile()` → write data with permissions.  
- Always `defer file.Close()` to release resources.

### 4. JSON Handling
- `json.MarshalIndent()` → convert Go struct → JSON.
- `json.Unmarshal()` → convert JSON → Go struct/slice.
- Use `encoding/json` package — standard and reliable.

### 5. CLI Flags
- `flag.NewFlagSet()` → defines subcommands like `add`, `list`, etc.
- `flag.String()` / `flag.Int()` → define CLI flags.
- `os.Args` → reads raw command-line input.

### 6. Project Modularity
- Organized packages: `internal/commands`, `logger`, `types`.
- Keeps `main.go` clean as the **entrypoint only**.
- Improves maintainability and testing.

### 7. Error Handling
- Always check `if err != nil`.
- Log detailed info for developers.
- Show user-friendly CLI messages.
- Return early to keep control flow clean.

---

## 🧪 Testing in Go
Go has built-in testing support — **no external libraries required**.

- Files ending in `_test.go` are auto-detected by `go test`.
- Tests receive `t *testing.T`, a small test engine for each test case.

**Common Methods:**
```go
t.Log("this is a log message")           // prints info
t.Error("something went wrong")          // marks test failed but continues
t.Errorf("expected %d, got %d", 1, 2)    // formatted version
t.Fatal("stop test immediately")         // fails and stops test here
```

## 🪵 Logger Levels Explained
Go’s slog provides structured logging with levels:
```
| Level | Constant          | Description            |
| ----- | ----------------- | ---------------------- |
| Debug | `slog.LevelDebug` | Detailed logs for devs |
| Info  | `slog.LevelInfo`  | High-level information |
| Warn  | `slog.LevelWarn`  | Potential issues       |
| Error | `slog.LevelError` | Critical problems only |
```

#⚙️ Exit Codes (for CLI status)
```
| Code | Meaning                                                  |
| ---- | -------------------------------------------------------- |
| `0`  | ✅ Success                                                |
| `1`  | ❌ General error (missing/unknown command)                |
| `2`  | ⚠️ Bad usage (invalid flags or syntax)                   |
| `3`  | 🧱 Runtime error (manager or file issues)                |
| `4`  | 💥 Operation-specific error (add/delete/update failures) |
```

## 📦 GoReleaser (for Future)

Automates packaging & releasing binaries for:
- All OS → Linux, Windows, macOS
- All architectures → amd64, arm64