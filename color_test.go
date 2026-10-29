package color_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/xogas/color"
)

// cleanEnv clears every environment variable that influences color.Enabled.
func cleanEnv(t *testing.T) {
	t.Helper()
	for _, key := range []string{"NO_COLOR", "FORCE_COLOR", "TERM"} {
		t.Setenv(key, "")
	}
}

// useOutput replaces color.Output for the duration of a test.
func useOutput(t *testing.T, w io.Writer) {
	t.Helper()
	prev := color.Output
	color.Output = w
	t.Cleanup(func() { color.Output = prev })
}

// useNoColor sets color.NoColor for the duration of a test.
func useNoColor(t *testing.T, v bool) {
	t.Helper()
	prev := color.NoColor
	color.NoColor = v
	t.Cleanup(func() { color.NoColor = prev })
}

// render returns s as c writes it to a writer, with color forced on.
func render(c color.Color, s string) string {
	var b strings.Builder
	_, _ = c.EnableColor().Fprint(&b, s)
	return b.String()
}

func TestZeroValue(t *testing.T) {
	cleanEnv(t)
	useOutput(t, &bytes.Buffer{})

	var c color.Color
	if got := c.Sprint("x"); got != "x" {
		t.Errorf("zero Color Sprint() = %q, want %q", got, "x")
	}
	if got := render(c, "x"); got != "x" {
		t.Errorf("zero Color render = %q, want %q", got, "x")
	}
	if c.Enabled() {
		t.Error("the zero Color should not be enabled for a buffer")
	}
}

func TestEmptyColorCannotBeEnabled(t *testing.T) {
	cleanEnv(t)
	useOutput(t, &bytes.Buffer{})
	t.Setenv("FORCE_COLOR", "1")

	if color.New().Enabled() {
		t.Error("an empty Color has nothing to emit and must not report enabled")
	}
	if !color.New(color.FgRed).Enabled() {
		t.Error("a non-empty Color should be enabled when color is forced")
	}
}

func TestAddIsImmutable(t *testing.T) {
	base := color.New(color.FgRed)
	bold := base.Add(color.Bold)

	if !base.SameStyle(color.New(color.FgRed)) {
		t.Error("Add must not modify the receiver")
	}
	if !bold.SameStyle(color.New(color.FgRed, color.Bold)) {
		t.Error("Add must return a color with the attribute appended")
	}
	if base.SameStyle(bold) {
		t.Error("the receiver and the result must differ")
	}
}

func TestEnableDisable(t *testing.T) {
	cleanEnv(t)
	useOutput(t, &bytes.Buffer{})
	c := color.New(color.FgRed)

	if c.Enabled() {
		t.Error("a buffer should be plain by default")
	}
	if !c.EnableColor().Enabled() {
		t.Error("EnableColor should force color on")
	}
	if c.DisableColor().Enabled() {
		t.Error("DisableColor should force color off")
	}
	if c.Enabled() {
		t.Error("EnableColor must not modify the receiver")
	}
}

func TestSameStyle(t *testing.T) {
	tests := []struct {
		name string
		a, b color.Color
		want bool
	}{
		{"same style", color.New(color.FgRed, color.Bold), color.New(color.FgRed, color.Bold), true},
		{"order matters", color.New(color.FgRed, color.Bold), color.New(color.Bold, color.FgRed), false},
		{"different colors", color.New(color.FgRed), color.New(color.FgBlue), false},
		{"overrides are ignored", color.New(color.FgRed).EnableColor(), color.New(color.FgRed), true},
		{"zero colors", color.Color{}, color.New(), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.SameStyle(tt.b); got != tt.want {
				t.Fatalf("SameStyle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSprint(t *testing.T) {
	cleanEnv(t)
	c := color.New(color.FgRed)

	if got, want := c.DisableColor().Sprint("x"), "x"; got != want {
		t.Errorf("disabled Sprint() = %q, want %q", got, want)
	}
	if got, want := c.EnableColor().Sprint("x"), "\x1b[31mx\x1b[39m"; got != want {
		t.Errorf("enabled Sprint() = %q, want %q", got, want)
	}
	if got := c.EnableColor().Sprint(""); got != "" {
		t.Errorf("Sprint(\"\") = %q, want %q", got, "")
	}
	if got, want := c.EnableColor().Sprint("a", "b"), "\x1b[31mab\x1b[39m"; got != want {
		t.Errorf("Sprint two strings = %q, want %q", got, want)
	}
}

func TestSprintf(t *testing.T) {
	cleanEnv(t)
	c := color.New(color.FgGreen).EnableColor()

	if got, want := c.Sprintf("ok %d", 42), "\x1b[32mok 42\x1b[39m"; got != want {
		t.Errorf("Sprintf() = %q, want %q", got, want)
	}
}

func TestSprintln(t *testing.T) {
	cleanEnv(t)
	c := color.New(color.FgRed).EnableColor()

	if got, want := c.Sprintln("x"), "\x1b[31mx\x1b[39m\n"; got != want {
		t.Errorf("Sprintln() = %q, want %q", got, want)
	}
	// An empty line is never wrapped in escape sequences.
	for _, tt := range []struct {
		name string
		got  string
	}{
		{"no operands", c.Sprintln()},
		{"empty operand", c.Sprintln("")},
	} {
		if tt.got != "\n" {
			t.Errorf("Sprintln %s = %q, want %q", tt.name, tt.got, "\n")
		}
	}
}

func TestStringFunctionsFollowOutputPolicy(t *testing.T) {
	cleanEnv(t)
	useOutput(t, &bytes.Buffer{})
	c := color.New(color.FgRed)

	if got := c.Sprint("x"); got != "x" {
		t.Errorf("Sprint over a buffer = %q, want %q", got, "x")
	}

	t.Setenv("FORCE_COLOR", "1")
	if got, want := c.Sprint("x"), "\x1b[31mx\x1b[39m"; got != want {
		t.Errorf("forced Sprint() = %q, want %q", got, want)
	}
}
