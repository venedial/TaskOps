package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	slogmulti "github.com/samber/slog-multi"

	"github.com/venedial/taskops/config"
)

var logFile *os.File

func getLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func setupStdoutLogger(cfg *config.StdoutLogConfig) slog.Handler {
	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: getLogLevel(cfg.Level)})
}

func setupFileLogger(cfg *config.FileLogConfig) (slog.Handler, error) {
	if cfg.Path == nil {
		return nil, fmt.Errorf("Path must not be nil!")
	}

	file, err := os.OpenFile(*cfg.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	logFile = file
	return slog.NewJSONHandler(file, &slog.HandlerOptions{
			Level: getLogLevel(cfg.Level),
		}),
		nil
}

func Setup(cfg *config.Configuration) {
	handlers := []slog.Handler{
		setupStdoutLogger(&cfg.Log.Stdout),
	}

	if cfg.Log.File != nil {
		fileHandler, err := setupFileLogger(cfg.Log.File)

		if err != nil {
			slog.Error("Unable to attach file logger. Error: ", "file", cfg.Log.File, slog.With("error", err))
		} else {
			handlers = append(handlers, fileHandler)
		}
	}

	logger := slog.New(slogmulti.Fanout(handlers...))

	slog.SetDefault(logger)
	slog.Debug("Logger sucessfully initialized!")
}

func Cleanup() {
	slog.Debug("Cleaning up logger...")

	if logFile != nil {
		slog.Debug("Found one open log file. Syncing and closing...")

		if err := logFile.Sync(); err != nil {
			slog.Warn("Failed to sync log file", slog.With("error", err))
		}

		if err := logFile.Close(); err != nil {
			slog.Warn("Failed to close log file", slog.With("error", err))
		} else {
			slog.Debug("Log file successfully closed!")
		}
	}

	slog.Debug("Logger sucessfully cleaned up!")
}
