package config

import (
	"path/filepath"
	"runtime"
	"os"
)

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder of thhis project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)

// Config func to get env value
func Config(key string) string {
	return os.Getenv(key)
}