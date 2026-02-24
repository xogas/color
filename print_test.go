package color_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/xogas/color"
)

// newNoColor returns a Color with color output disabled so test assertions
// can compare plain text without ANSI escape sequences.
func newNoColor(attrs ...color.Attribute) *color.Color {
	c := color.New(attrs...)
	c.DisableColor()
	return c
}

// newWithColor returns a Color with color output force-enabled, regardless of
// the global NoColor setting or whether stdout is a terminal.
func newWithColor(attrs ...color.Attribute) *color.Color {
	c := color.New(attrs...)
	c.EnableColor()
	return c
}

func TestSet(t *testing.T) {
	tests := []struct {
		name    string
		noColor bool
		wantESC bool
	}{
		{"with color", false, true},
		{"no color", true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prev := color.NoColor
			color.NoColor = tt.noColor
			defer func() { color.NoColor = prev }()

			var buf bytes.Buffer
			color.Output = &buf

			color.Set(color.FgRed)
			hasESC := strings.Contains(buf.String(), "\x1b[")
			if hasESC != tt.wantESC {
				t.Errorf("Set() wrote %q, wantESC=%v", buf.String(), tt.wantESC)
			}
		})
	}
}

func TestUnset(t *testing.T) {
	tests := []struct {
		name    string
		noColor bool
		wantESC bool
	}{
		{"with color", false, true},
		{"no color", true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prev := color.NoColor
			color.NoColor = tt.noColor
			defer func() { color.NoColor = prev }()

			var buf bytes.Buffer
			color.Output = &buf

			color.Unset()
			hasESC := strings.Contains(buf.String(), "\x1b[0m")
			if hasESC != tt.wantESC {
				t.Errorf("Unset() wrote %q, wantESC=%v", buf.String(), tt.wantESC)
			}
		})
	}
}

func TestSetWriter(t *testing.T) {
	tests := []struct {
		name    string
		c       *color.Color
		wantESC bool
	}{
		{"with color", newWithColor(color.FgRed), true},
		{"no color", newNoColor(color.FgRed), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			tt.c.SetWriter(&buf)
			hasESC := strings.Contains(buf.String(), "\x1b[")
			if hasESC != tt.wantESC {
				t.Errorf("SetWriter() wrote %q, wantESC=%v", buf.String(), tt.wantESC)
			}
		})
	}
}

func TestUnsetWriter(t *testing.T) {
	tests := []struct {
		name    string
		c       *color.Color
		wantESC bool
	}{
		{"with color", newWithColor(color.FgRed), true},
		{"no color", newNoColor(color.FgRed), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			tt.c.UnsetWriter(&buf)
			hasESC := strings.Contains(buf.String(), "\x1b[")
			if hasESC != tt.wantESC {
				t.Errorf("UnsetWriter() wrote %q, wantESC=%v", buf.String(), tt.wantESC)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	tests := []struct {
		name    string
		c       *color.Color
		noColor bool
		wantOut string
	}{
		{
			"no color",
			newNoColor(color.Bold),
			true,
			"hello world",
		},
		{
			"with color",
			newWithColor(color.FgRed),
			false,
			"\x1b[31m" + "hello world" + "\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prev := color.NoColor
			color.NoColor = tt.noColor
			defer func() { color.NoColor = prev }()

			var buf bytes.Buffer
			color.Output = &buf
			_, err := tt.c.Print("hello", " ", "world")
			if err != nil {
				t.Fatalf("Print error: %v", err)
			}
			if got := buf.String(); got != tt.wantOut {
				t.Errorf("Print() = %q, want %q", got, tt.wantOut)
			}
		})
	}
}

func TestPrintf(t *testing.T) {
	tests := []struct {
		name    string
		c       *color.Color
		noColor bool
		wantOut string
	}{
		{
			"no color",
			newNoColor(color.Underline),
			true,
			"x=7",
		},
		{
			"with color",
			newWithColor(color.FgRed),
			false,
			"\x1b[31m" + "x=7" + "\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prev := color.NoColor
			color.NoColor = tt.noColor
			defer func() { color.NoColor = prev }()

			var buf bytes.Buffer
			color.Output = &buf
			_, err := tt.c.Printf("x=%d", 7)
			if err != nil {
				t.Fatalf("Printf error: %v", err)
			}
			if got := buf.String(); got != tt.wantOut {
				t.Errorf("Printf() = %q, want %q", got, tt.wantOut)
			}
		})
	}
}

func TestPrintln(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf

	n, err := newNoColor(color.FgWhite).Println("hello", " ", "world")
	if err != nil {
		t.Fatalf("Println error: %v", err)
	}
	got, want := buf.String(), "hello world\n"
	if got != want {
		t.Errorf("Println() wrote %q, want %q", got, want)
	}
	if n != len(want) {
		t.Errorf("Println() n = %d, want %d", n, len(want))
	}
}

func TestSprint(t *testing.T) {
	tests := []struct {
		args []any
		want string
	}{
		{[]any{"hello"}, "hello"},
		{[]any{"hello", " ", "world"}, "hello world"},
		{[]any{42}, "42"},
		{[]any{}, ""},
	}
	for _, tt := range tests {
		got := newNoColor(color.FgRed).Sprint(tt.args...)
		if got != tt.want {
			t.Errorf("Sprint(%v) = %q, want %q", tt.args, got, tt.want)
		}
	}
}

func TestSprintf(t *testing.T) {
	tests := []struct {
		format string
		args   []any
		want   string
	}{
		{"hello %s", []any{"world"}, "hello world"},
		{"%d + %d = %d", []any{1, 2, 3}, "1 + 2 = 3"},
		{"no args", nil, "no args"},
	}
	for _, tt := range tests {
		got := newNoColor(color.FgGreen).Sprintf(tt.format, tt.args...)
		if got != tt.want {
			t.Errorf("Sprintf(%q, %v) = %q, want %q", tt.format, tt.args, got, tt.want)
		}
	}
}

func TestSprintln(t *testing.T) {
	tests := []struct {
		args []any
		want string
	}{
		{[]any{"hello"}, "hello\n"},
		{[]any{"hello", " ", "world"}, "hello world\n"},
		{[]any{}, "\n"},
	}
	for _, tt := range tests {
		got := newNoColor(color.FgBlue).Sprintln(tt.args...)
		if got != tt.want {
			t.Errorf("Sprintln(%v) = %q, want %q", tt.args, got, tt.want)
		}
	}
}

func TestFprint(t *testing.T) {
	tests := []struct {
		args []any
		want string
	}{
		{[]any{"hello"}, "hello"},
		{[]any{"hello", " ", "world"}, "hello world"},
		{[]any{42}, "42"},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		n, err := newNoColor(color.FgYellow).Fprint(&buf, tt.args...)
		if err != nil {
			t.Errorf("Fprint(%v) error: %v", tt.args, err)
		}
		got := buf.String()
		if got != tt.want {
			t.Errorf("Fprint(%v) wrote %q, want %q", tt.args, got, tt.want)
		}
		if n != len(tt.want) {
			t.Errorf("Fprint(%v) n = %d, want %d", tt.args, n, len(tt.want))
		}
	}
}

func TestFprintf(t *testing.T) {
	var buf bytes.Buffer
	n, err := newNoColor(color.FgCyan).Fprintf(&buf, "value=%d", 99)
	if err != nil {
		t.Fatalf("Fprintf error: %v", err)
	}
	got, want := buf.String(), "value=99"
	if got != want {
		t.Errorf("Fprintf() wrote %q, want %q", got, want)
	}
	if n != len(want) {
		t.Errorf("Fprintf() n = %d, want %d", n, len(want))
	}
}

func TestFprintln(t *testing.T) {
	tests := []struct {
		args []any
		want string
	}{
		{[]any{"hello"}, "hello\n"},
		{[]any{"hello", " ", "world"}, "hello world\n"},
		{[]any{42}, "42\n"},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		n, err := newNoColor(color.FgMagenta).Fprintln(&buf, tt.args...)
		if err != nil {
			t.Errorf("Fprintln(%v) error: %v", tt.args, err)
		}
		got := buf.String()
		if got != tt.want {
			t.Errorf("Fprintln(%v) wrote %q, want %q", tt.args, got, tt.want)
		}
		if n != len(tt.want) {
			t.Errorf("Fprintln(%v) n = %d, want %d", tt.args, n, len(tt.want))
		}
	}
}
