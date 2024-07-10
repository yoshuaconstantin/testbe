package logging

import (
	"fmt"
	"log"
	"os"
	"log/slog"
)

var logger *slog.Logger
var err error

// Open the log file in append mode or create it if it doesn't exist
func openLogFile() (*os.File, error) {
	return os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}

func CreateLogger() (*slog.Logger, error) {
	logFile, err := openLogFile()
	if err != nil {
		return nil, fmt.Errorf("failed to open or create log file: %w", err)
	}

	handler := slog.NewTextHandler(logFile, nil)
	return slog.New(handler), nil
}

func init() {
	logger, err = CreateLogger()
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
}

// GetError returns the logger initialization error
func GetError() error {
	return err
}

// With creates a logger with additional fields
func With(key string, value interface{}) *slog.Logger {
	if err != nil {
		log.Println("Error initializing logger:", err)
		return nil
	}
	return logger.With(key, value)
}

// Info logs informational messages
func Info(msg string) {
	if err != nil {
		log.Println("Error initializing logger:", err)
		return
	}
	logger.Info(msg)
}

// Warn logs warning messages
func Warn(msg string) {
	if err != nil {
		log.Println("Error initializing logger:", err)
		return
	}
	logger.Warn(msg)
}

// Error logs error messages with variadic arguments
func Error(msg string, args string) {
	if err != nil {
		log.Println("Error initializing logger:", err)
		return
	}
	logger.Error(fmt.Sprintf(msg, args))
}
