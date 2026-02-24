package color

import (
	"fmt"
	"io"
)

// PrintFunc returns a function that writes colorized output to Output using
// Print.
func (c *Color) PrintFunc() func(a ...any) {
	return func(a ...any) { _, _ = c.Print(a...) }
}

// PrintfFunc returns a function that writes colorized output to Output using
// Printf.
func (c *Color) PrintfFunc() func(format string, a ...any) {
	return func(format string, a ...any) { _, _ = c.Printf(format, a...) }
}

// PrintlnFunc returns a function that writes colorized output to Output using
// Println.
func (c *Color) PrintlnFunc() func(a ...any) {
	return func(a ...any) { _, _ = c.Println(a...) }
}

// SprintFunc returns a function that returns the colorized string using
// Sprint.
func (c *Color) SprintFunc() func(a ...any) string {
	return func(a ...any) string { return c.wrap(fmt.Sprint(a...)) }
}

// SprintfFunc returns a function that returns the colorized string using
// Sprintf.
func (c *Color) SprintfFunc() func(format string, a ...any) string {
	return func(format string, a ...any) string { return c.wrap(fmt.Sprintf(format, a...)) }
}

// SprintlnFunc returns a function that returns the colorized string using
// Sprintln (with trailing newline).
func (c *Color) SprintlnFunc() func(a ...any) string {
	return func(a ...any) string { return c.wrap(fmt.Sprint(a...)) + "\n" }
}

// FprintFunc returns a function that writes colorized output to w using
// Fprint.
func (c *Color) FprintFunc() func(w io.Writer, a ...any) {
	return func(w io.Writer, a ...any) { _, _ = c.Fprint(w, a...) }
}

// FprintfFunc returns a function that writes colorized output to w using
// Fprintf.
func (c *Color) FprintfFunc() func(w io.Writer, format string, a ...any) {
	return func(w io.Writer, format string, a ...any) { _, _ = c.Fprintf(w, format, a...) }
}

// FprintlnFunc returns a function that writes colorized output to w using
// Fprintln.
func (c *Color) FprintlnFunc() func(w io.Writer, a ...any) {
	return func(w io.Writer, a ...any) { _, _ = c.Fprintln(w, a...) }
}
