package logging

import (
	"log"
	"os"
	"sync"
)

var (
	// InfoLogger logging info
	InfoLogger *log.Logger

	// ErrorLogger logging errors
	ErrorLogger *log.Logger

	// FatalLogger logging fatal
	FatalLogger *log.Logger

	once sync.Once
)

// GetInfoLogger returns the InfoLogger
func GetInfoLogger() *log.Logger {
	once.Do(func() {
		InfoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	})
	return InfoLogger
}

// GetErrorLogger returns the ErrorLogger
func GetErrorLogger() *log.Logger {
	once.Do(func() {
		ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	})
	return ErrorLogger
}

// GetFatalLogger returns the FatalLogger
func GetFatalLogger() *log.Logger {
	once.Do(func() {
		FatalLogger = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	})
	return FatalLogger
}
