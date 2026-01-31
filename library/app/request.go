package app

import (
	"github.com/everyday-items/gin-example/library/logging"
)

// LogError logs error message
func LogError(key, message string) {
	logging.Error(key, message)
}
