// +build windows

// -----------------------------------------------------------------------------
// PACKAGE (Windows Platform Only)
// -----------------------------------------------------------------------------

package logger

// -----------------------------------------------------------------------------
// IMPORTS
// -----------------------------------------------------------------------------

import (
	"fmt"
	"io"
)

// -----------------------------------------------------------------------------
// PRIVATE METHODS
// -----------------------------------------------------------------------------

func clearLine(std io.Writer) {
	fmt.Fprintf(std, "\r\r") // removes all the content in the current line and hide the cursor
}

func newLine(std io.Writer) {
	fmt.Fprintf(std, "\n") // appends a new line and show the cursor
}
