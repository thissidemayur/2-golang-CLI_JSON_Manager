package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

var Log *slog.Logger

func InitLogger(env string) {
	var handler slog.Handler

	switch env {
	case "prod", "production":
		// Use minimal logging — only show errors
		handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelError,
		})
	default:
		// Development mode — show colorful structured logs
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	if env == "prod" || env == "production" {
		// ensure log folder exist
		logDir := "logs"
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			fmt.Println("⚠️  Failed to create log directory.")
		}

		logFile := filepath.Join(logDir, "app.log")
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("⚠️  Failed to open log file.")
			file = os.Stdout //fallback
		}
		handler = slog.NewJSONHandler(file, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})

	} else {
		// Development mode — show colorful structured logs
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
		fmt.Println("🧰 Development logs will print to console.")

	}

	Log = slog.New(handler)
	slog.SetDefault(Log)

	slog.Info("✅ Logger initialized successfully", "env", env)
}

// Info logs informational messages
func Info(msg string, args ...any) {
	Log.Info(msg, args...)
}

// Error logs error messages
func Error(msg string, args ...any) {
	Log.Error(msg, args...)
}

// Debug logs debugging message
func Debug(msg string, args ...any) {
	Log.Debug(msg, args...)

}

// Warn logs warning messages
func Warn(msg string, args ...any) {
	Log.Warn(msg, args...)
}
