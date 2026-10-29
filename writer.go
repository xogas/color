package color

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Output is the destination of Print, Printf and Println. It is also the writer
// whose terminal status decides whether Sprint, Sprintf and Sprintln emit
// escape sequences. A nil Output behaves like io.Discard.
//
// NoColor and Output are global; mutating them while other goroutines print is
// a data race.
var Output io.Writer = os.Stdout

// Fprint writes the colorized result of fmt.Sprint to w. It returns the number
// of bytes written and any write error encountered.
func (c Color) Fprint(w io.Writer, a ...any) (int, error) {
	return c.write(w, fmt.Sprint(a...))
}

// Fprintf writes the colorized result of fmt.Sprintf to w.
func (c Color) Fprintf(w io.Writer, format string, a ...any) (int, error) {
	return c.write(w, fmt.Sprintf(format, a...))
}

// Fprintln is like Fprint but appends a newline after the reset sequence.
func (c Color) Fprintln(w io.Writer, a ...any) (int, error) {
	return c.writeLine(w, strings.TrimSuffix(fmt.Sprintln(a...), "\n"))
}

// Print writes to Output.
func (c Color) Print(a ...any) (int, error) { return c.Fprint(writerOr(Output), a...) }

// Printf writes to Output.
func (c Color) Printf(format string, a ...any) (int, error) {
	return c.Fprintf(writerOr(Output), format, a...)
}

// Println writes to Output.
func (c Color) Println(a ...any) (int, error) {
	return c.Fprintln(writerOr(Output), a...)
}

// SetWriter writes the escape sequence that enables c to w. It is meant for
// streaming a block of output and should be paired with UnsetWriter.
func (c Color) SetWriter(w io.Writer) { c.writeSequence(w, c.on) }

// UnsetWriter writes the escape sequence that disables c to w.
func (c Color) UnsetWriter(w io.Writer) { c.writeSequence(w, c.off) }

// Set writes the escape sequence that enables c to Output.
func (c Color) Set() { c.SetWriter(Output) }

// Unset writes the escape sequence that disables c to Output.
func (c Color) Unset() { c.UnsetWriter(Output) }

// write writes s to w, wrapped in c's escape sequences when enabled.
// A nil w discards the output.
func (c Color) write(w io.Writer, s string) (int, error) {
	return io.WriteString(writerOr(w), c.render(s, c.enabledFor(w)))
}

// writeLine writes s to w followed by a newline outside the escape sequences.
func (c Color) writeLine(w io.Writer, s string) (int, error) {
	return io.WriteString(writerOr(w), c.renderLine(s, c.enabledFor(w)))
}

// writeSequence writes seq to w when c is enabled for w. Write errors are discarded.
func (c Color) writeSequence(w io.Writer, seq string) {
	if seq == "" || !c.enabledFor(w) {
		return
	}
	_, _ = io.WriteString(writerOr(w), seq)
}

// writerOr returns w, or io.Discard when w is nil.
func writerOr(w io.Writer) io.Writer {
	if w == nil {
		return io.Discard
	}
	return w
}
