// -----------------------------------------------------------------------------
// PACKAGE
// -----------------------------------------------------------------------------

package logger

// -----------------------------------------------------------------------------
// IMPORTS
// -----------------------------------------------------------------------------

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

// -----------------------------------------------------------------------------
// STRUCT
// -----------------------------------------------------------------------------

type sLogger struct {
	colored    bool
	verbose    bool
	inprogress bool
}

// Logger TODO
var Logger = &sLogger{
	colored:    true,
	verbose:    false,
	inprogress: false,
}

// -----------------------------------------------------------------------------
// PUBLIC METHODS
// -----------------------------------------------------------------------------

// Command displays the current command along with a brief description in a beautiful way
func Command(name string, msg string, a ...interface{}) {
	s := fmt.Sprintf(msg, a...)
	if Logger.colored {
		name = color.FgLightMagenta.Render(name)
	}
	write(false, false, fmt.Sprintf("\n%s ❖ %s", name, s))
}

// Put displays or writes text without color or formatting
func Put(msg string, a ...interface{}) {
	write(false, false, fmt.Sprintf(msg, a...))
}

// Warning displays a light yellow message
func Warning(msg string, a ...interface{}) {
	s := fmt.Sprintf(msg, a...)
	if Logger.colored {
		s = color.Gray.Render(fmt.Sprintf("• %s", s))
	}
	write(false, false, s)
}

// Progress TODO
func Progress(msg string, a ...interface{}) {
	s := fmt.Sprintf(msg, a...)
	write(false, true, color.FgGray.Render(fmt.Sprintf("○ %s", s)))
}

// -----------------------------------------------------------------------------
// PRIVATE METHODS
// -----------------------------------------------------------------------------

func write(iserr bool, inprogress bool, a ...interface{}) {
	// set standard stream
	std := os.Stdout
	if iserr {
		std = os.Stderr
	}
	// check if a clear is needed
	if Logger.inprogress {
		clearLine(std)
	}
	Logger.inprogress = inprogress // remove line in next print
	fmt.Fprint(std, a...)          // print message
	// new line if not in progress
	if !Logger.inprogress {
		newLine(std)
	}
}
