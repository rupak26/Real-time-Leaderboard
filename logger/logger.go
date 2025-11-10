package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"
	"fmt"
)

const (
	MaxSizeBytes = 5 * 1024 * 1024 // 5 MB
	DateLayout   = "2006-01-02"
	TimeLayout   = "2006-01-02_15-04-05"
)

func SetupLogger() *slog.Logger {
    logsDir := "logs"
	archiveDir := filepath.Join(logsDir, "archive")

    if err := os.MkdirAll(logsDir, 0755); err != nil {
        log.Fatalf("Failed to create logs directory: %v", err)
    }
    
    logPath := filepath.Join(logsDir, "current.log")
    rotateIfNeeded(logsDir,archiveDir,logPath)

    logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }

    multiWriter := io.MultiWriter(logFile, os.Stdout)

    handler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
        Level:     slog.LevelDebug,
        AddSource: true,
    })

    logger := slog.New(handler)
    // Log a test message
 //   logger.Info("Logger initialized", "file", logPath)

    return logger
}


func rotateIfNeeded(logsDir , archiveDir,logFilePath string) {
	info, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		return // no file yet
	}

	if err != nil {
		slog.Warn("Unable to stat log file", "error", err)
		return
	}

	// Get last modified date
	modTime := info.ModTime().Format(DateLayout)
	currentDate := time.Now().Format(DateLayout)

	// Check for rotation conditions
	sizeExceeded := info.Size() >= MaxSizeBytes
	newDay := currentDate != modTime

	if sizeExceeded || newDay {
		timestamp := time.Now().Format(TimeLayout)
		
		archivedName := fmt.Sprintf("log_%s.log", timestamp)
		archivedPath := filepath.Join(archiveDir, archivedName)

		// Close before renaming
		if file, err := os.OpenFile(logFilePath, os.O_WRONLY, 0644); err == nil {
			file.Close()
		}

		// Move to archive folder
		if err := os.Rename(logFilePath, archivedPath); err != nil {
			slog.Warn("Failed to move log to archive", "error", err)
			return
		}
	}
}

