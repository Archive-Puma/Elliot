// +build !windows

// -----------------------------------------------------------------------------
// PACKAGE (Windows Platform Excluded)
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
	fmt.Fprintf(std, "\033[?25l\r\033[2K") // removes all the content in the current line and hide the cursor
}

func newLine(std io.Writer) {
	fmt.Fprintf(std, "\033[?25h\n") // appends a new line and show the cursor
}
