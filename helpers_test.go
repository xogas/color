package color_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/xogas/color"
)

func TestColorString(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string, ...any) string
	}{
		{"BlackString", color.BlackString},
		{"RedString", color.RedString},
		{"GreenString", color.GreenString},
		{"YellowString", color.YellowString},
		{"BlueString", color.BlueString},
		{"MagentaString", color.MagentaString},
		{"CyanString", color.CyanString},
		{"WhiteString", color.WhiteString},
		{"HiBlackString", color.HiBlackString},
		{"HiRedString", color.HiRedString},
		{"HiGreenString", color.HiGreenString},
		{"HiYellowString", color.HiYellowString},
		{"HiBlueString", color.HiBlueString},
		{"HiMagentaString", color.HiMagentaString},
		{"HiCyanString", color.HiCyanString},
		{"HiWhiteString", color.HiWhiteString},
	}

	t.Run("String/no_color_passthrough", func(t *testing.T) {
		prev := color.NoColor
		color.NoColor = true
		defer func() { color.NoColor = prev }()

		for _, tt := range tests {
			got := tt.fn("hello %s", "world")
			if got != "hello world" {
				t.Errorf("%s() = %q, want %q", tt.name, got, "hello world")
			}
		}
	})

	t.Run("String/with_color_has_ansi", func(t *testing.T) {
		prev := color.NoColor
		color.NoColor = false
		defer func() { color.NoColor = prev }()

		for _, tt := range tests {
			got := tt.fn("hello")
			if !strings.Contains(got, "\x1b[") {
				t.Errorf("%s() = %q, want ANSI escape sequence", tt.name, got)
			}
			if !strings.Contains(got, "hello") {
				t.Errorf("%s() = %q, missing text", tt.name, got)
			}
		}
	})
}

func TestColor(t *testing.T) {

	tests := []struct {
		name string
		fn   func(string, ...any)
	}{
		{"Black", color.Black},
		{"Red", color.Red},
		{"Green", color.Green},
		{"Yellow", color.Yellow},
		{"Blue", color.Blue},
		{"Magenta", color.Magenta},
		{"Cyan", color.Cyan},
		{"White", color.White},
		{"HiBlack", color.HiBlack},
		{"HiRed", color.HiRed},
		{"HiGreen", color.HiGreen},
		{"HiYellow", color.HiYellow},
		{"HiBlue", color.HiBlue},
		{"HiMagenta", color.HiMagenta},
		{"HiCyan", color.HiCyan},
		{"HiWhite", color.HiWhite},
	}

	t.Run("Print/appends_newline", func(t *testing.T) {
		prev := color.NoColor
		color.NoColor = true
		defer func() { color.NoColor = prev }()

		for _, tt := range tests {
			var buf bytes.Buffer
			color.Output = &buf

			tt.fn("hello")
			got := buf.String()
			if got != "hello\n" {
				t.Errorf("%s() wrote %q, want %q", tt.name, got, "hello\\n")
			}
		}
	})

	t.Run("Print/format_args", func(t *testing.T) {
		prev := color.NoColor
		color.NoColor = true
		defer func() { color.NoColor = prev }()

		for _, tt := range tests {
			var buf bytes.Buffer
			color.Output = &buf

			tt.fn("x=%d", 42)
			got := buf.String()
			if got != "x=42\n" {
				t.Errorf("%s() wrote %q, want %q", tt.name, got, "x=42\\n")
			}
		}
	})

	t.Run("Print/already_has_newline", func(t *testing.T) {
		prev := color.NoColor
		color.NoColor = true
		defer func() { color.NoColor = prev }()

		for _, tt := range tests {
			var buf bytes.Buffer
			color.Output = &buf

			tt.fn("hello\n")
			got := buf.String()
			if got != "hello\n" {
				t.Errorf("%s() wrote %q, want exactly one newline", tt.name, got)
			}
		}
	})
}
