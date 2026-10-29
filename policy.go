package color

import (
	"io"
	"os"
	"strings"

	"golang.org/x/term"
)

// NoColor disables color output for every Color that carries no per-value
// override. It is false by default and is never changed by this package.
//
// NoColor and Output are global; mutating them while other goroutines print is
// a data race.
var NoColor bool

// Enabled reports whether the package policy allows SGR sequences on w. It
// ignores per-value overrides and is evaluated on every call, so environment
// changes take effect immediately.
//
// The precedence is documented in the package documentation.
func Enabled(w io.Writer) bool {
	if NoColor {
		return false
	}
	if force, ok := envBool("FORCE_COLOR"); ok {
		return force
	}
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	if strings.EqualFold(os.Getenv("TERM"), "dumb") {
		return false
	}
	f, ok := w.(fdWriter)
	return ok && term.IsTerminal(int(f.Fd()))
}

// envBool reports the boolean value of the environment variable key and whether
// it is set at all. The empty string counts as unset; "0", "false", "no" and
// "off" (case insensitive) are false; any other value is true.
func envBool(key string) (value, ok bool) {
	v := os.Getenv(key)
	switch {
	case v == "":
		return false, false
	case v == "0",
		strings.EqualFold(v, "false"),
		strings.EqualFold(v, "no"),
		strings.EqualFold(v, "off"):
		return false, true
	default:
		return true, true
	}
}

// fdWriter is implemented by *os.File and by wrappers that can report the file
// descriptor behind a terminal.
type fdWriter interface {
	Fd() uintptr
}

// enabledFor reports whether c may emit escape sequences on w.
func (c Color) enabledFor(w io.Writer) bool {
	switch c.mode {
	case colorAlways:
		return true
	case colorNever:
		return false
	default:
		return Enabled(w)
	}
}
