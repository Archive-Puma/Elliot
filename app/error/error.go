package error

import (
	"fmt"
	"os"
)

type tLevel int

const (
	// CRITICAL TODO: Doc
	CRITICAL tLevel = 10
	// WARNING TODO: Doc
	WARNING tLevel = 20
)

// MrRobotError TODO: Doc
type MrRobotError struct {
	message string
	level   tLevel
}

// NewError TODO: Doc
func NewError(level tLevel, message string, a ...interface{}) *MrRobotError {
	msg := fmt.Sprintf(message, a...)
	return &MrRobotError{message: msg, level: level}
}

// NewWarning TODO: Doc
func NewWarning(message string, a ...interface{}) *MrRobotError {
	return NewError(WARNING, message, a...)
}

// NewCritical TODO: Doc
func NewCritical(message string, a ...interface{}) *MrRobotError {
	return NewError(CRITICAL, message, a...)
}

// Resolve TODO: Doc
func (err MrRobotError) Resolve(verbosity bool) {
	if verbosity {
		if err.level == CRITICAL {
			fmt.Printf("[!] %s\n", err.message)
			os.Exit(1)
		} else {
			fmt.Printf("[-] %s\n", err.message)
		}
	}
}
