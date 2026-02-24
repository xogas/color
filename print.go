package color

import (
	"fmt"
	"io"
)

// Set sets the given SGR attributes on the global Output immediately. Call
// Unset() when done.
func Set(p ...Attribute) *Color {
	c := New(p...)
	c.set()
	return c
}

// Unset resets all SGR attributes on the global Output.
func Unset() {
	if NoColor {
		return
	}
	_, _ = fmt.Fprintf(Output, "%s[%dm", escape, Reset)
}

func (c *Color) set() *Color {
	if c.isNoColorSet() {
		return c
	}
	_, _ = fmt.Fprint(Output, c.format())
	return c
}

func (c *Color) unset() {
	if c.isNoColorSet() {
		return
	}
	Unset()
}

// SetWriter writes the opening SGR sequence to w.
func (c *Color) SetWriter(w io.Writer) *Color {
	if c.isNoColorSet() {
		return c
	}
	_, _ = fmt.Fprint(w, c.format())
	return c
}

// UnsetWriter writes the reset SGR sequence to w.
func (c *Color) UnsetWriter(w io.Writer) {
	if c.isNoColorSet() {
		return
	}
	_, _ = fmt.Fprintf(w, "%s[%dm", escape, Reset)
}

// Print formats and writes to Output with color applied.
func (c *Color) Print(a ...any) (n int, err error) {
	c.set()
	defer c.unset()
	return fmt.Fprint(Output, a...)
}

// Printf formats according to a format specifier and writes to Output with
// color applied.
func (c *Color) Printf(format string, a ...any) (n int, err error) {
	c.set()
	defer c.unset()
	return fmt.Fprintf(Output, format, a...)
}

// Println formats and writes to Output with color applied. A newline is always
// appended.
func (c *Color) Println(a ...any) (n int, err error) {
	return fmt.Fprintln(Output, c.wrap(fmt.Sprint(a...)))
}

// Sprint formats using the default formats and returns the colorized string.
func (c *Color) Sprint(a ...any) string {
	return c.wrap(fmt.Sprint(a...))
}

// Sprintf formats according to a format specifier and returns the colorized
// string.
func (c *Color) Sprintf(format string, a ...any) string {
	return c.wrap(fmt.Sprintf(format, a...))
}

// Sprintln formats using the default formats and returns the colorized string
// with a trailing newline.
func (c *Color) Sprintln(a ...any) string {
	return c.wrap(fmt.Sprint(a...)) + "\n"
}

// Fprint formats and writes to w with color applied. Spaces are added between
// operands when neither is a string.
func (c *Color) Fprint(w io.Writer, a ...any) (n int, err error) {
	c.SetWriter(w)
	defer c.UnsetWriter(w)
	return fmt.Fprint(w, a...)
}

// Fprintf formats according to a format specifier and writes to w with color
// applied.
func (c *Color) Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	c.SetWriter(w)
	defer c.UnsetWriter(w)
	return fmt.Fprintf(w, format, a...)
}

// Fprintln formats and writes to w with color applied. A newline is always
// appended.
func (c *Color) Fprintln(w io.Writer, a ...any) (n int, err error) {
	return fmt.Fprintln(w, c.wrap(fmt.Sprint(a...)))
}
