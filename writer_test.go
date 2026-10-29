package color_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/xogas/color"
)

// failingWriter fails every write with errWrite.
type failingWriter struct{}

var errWrite = errors.New("write failed")

func (failingWriter) Write(p []byte) (int, error) { return 0, errWrite }

func TestFprint(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer
	c := color.New(color.FgRed)

	n, err := c.Fprint(&buf, "x")
	if err != nil || n != 1 || buf.String() != "x" {
		t.Fatalf("Fprint to buffer = %q (%d, %v), want %q", buf.String(), n, err, "x")
	}

	buf.Reset()
	want := "\x1b[31mx\x1b[39m"
	n, err = c.EnableColor().Fprint(&buf, "x")
	if err != nil || buf.String() != want || n != len(want) {
		t.Fatalf("forced Fprint = %q (%d, %v), want %q", buf.String(), n, err, want)
	}
}

func TestFprintf(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer

	if _, err := color.New(color.FgGreen).EnableColor().Fprintf(&buf, "ok %d", 1); err != nil {
		t.Fatal(err)
	}
	if got, want := buf.String(), "\x1b[32mok 1\x1b[39m"; got != want {
		t.Fatalf("Fprintf() = %q, want %q", got, want)
	}
}

func TestFprintln(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer

	if _, err := color.New(color.FgRed).EnableColor().Fprintln(&buf, "a", 1); err != nil {
		t.Fatal(err)
	}
	if got, want := buf.String(), "\x1b[31ma 1\x1b[39m\n"; got != want {
		t.Fatalf("Fprintln() = %q, want %q", got, want)
	}

	buf.Reset()
	if _, err := color.New(color.FgRed).EnableColor().Fprintln(&buf); err != nil {
		t.Fatal(err)
	}
	if got := buf.String(); got != "\n" {
		t.Fatalf("Fprintln() without operands = %q, want %q", got, "\n")
	}
}

func TestNilWriterDiscards(t *testing.T) {
	cleanEnv(t)
	c := color.New(color.FgRed)

	n, err := c.Fprint(nil, "x")
	if n != 1 || err != nil {
		t.Fatalf("Fprint(nil) = (%d, %v), want (1, nil)", n, err)
	}
	if _, err := c.EnableColor().Fprintln(nil, "x"); err != nil {
		t.Fatalf("Fprintln(nil) error = %v", err)
	}
}

func TestWriteErrorPropagates(t *testing.T) {
	cleanEnv(t)
	c := color.New(color.FgRed)

	if _, err := c.Fprint(failingWriter{}, "x"); !errors.Is(err, errWrite) {
		t.Fatalf("plain Fprint error = %v, want %v", err, errWrite)
	}
	if _, err := c.EnableColor().Fprintf(failingWriter{}, "%s", "x"); !errors.Is(err, errWrite) {
		t.Fatalf("colored Fprintf error = %v, want %v", err, errWrite)
	}
}

func TestPrintUsesOutput(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer
	useOutput(t, &buf)

	if _, err := color.New(color.FgGreen).Printf("ok %d", 1); err != nil {
		t.Fatal(err)
	}
	if got, want := buf.String(), "ok 1"; got != want {
		t.Fatalf("Printf to Output = %q, want %q", got, want)
	}

	buf.Reset()
	if _, err := color.New(color.FgGreen).EnableColor().Println("ok"); err != nil {
		t.Fatal(err)
	}
	if got, want := buf.String(), "\x1b[32mok\x1b[39m\n"; got != want {
		t.Fatalf("Println to Output = %q, want %q", got, want)
	}
}

func TestNilOutputDiscards(t *testing.T) {
	cleanEnv(t)
	useOutput(t, nil)

	n, err := color.New(color.FgRed).Print("x")
	if n != 1 || err != nil {
		t.Fatalf("Print with nil Output = (%d, %v), want (1, nil)", n, err)
	}
}

func TestSetWriterPair(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer
	c := color.New(color.FgRed).EnableColor()

	c.SetWriter(&buf)
	c.UnsetWriter(&buf)

	if got, want := buf.String(), "\x1b[31m\x1b[39m"; got != want {
		t.Fatalf("SetWriter pair = %q, want %q", got, want)
	}
}

func TestSetWriterUsesPolicy(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer
	c := color.New(color.FgRed)

	c.SetWriter(&buf)
	c.UnsetWriter(&buf)

	if got := buf.String(); got != "" {
		t.Fatalf("SetWriter on a buffer = %q, want no output", got)
	}
}

func TestSetUsesOutput(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer
	useOutput(t, &buf)

	color.New(color.FgRed).EnableColor().Set()
	color.New(color.FgRed).EnableColor().Unset()

	if got, want := buf.String(), "\x1b[31m\x1b[39m"; got != want {
		t.Fatalf("Set/Unset on Output = %q, want %q", got, want)
	}
}

func TestSetEmptyColorWritesNothing(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer

	color.New().EnableColor().SetWriter(&buf)
	color.New().EnableColor().UnsetWriter(&buf)

	if got := buf.String(); got != "" {
		t.Fatalf("empty Color wrote %q, want no output", got)
	}
}
