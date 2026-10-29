package color_test

import (
	"bytes"
	"testing"

	"github.com/xogas/color"
)

func TestNamedHelpers(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer
	useOutput(t, &buf)
	t.Setenv("FORCE_COLOR", "1")

	tests := []struct {
		name  string
		on    string
		off   string
		print func(format string, a ...any)
		str   func(format string, a ...any) string
	}{
		{"Black", "\x1b[30m", "\x1b[39m", color.Black, color.BlackString},
		{"Red", "\x1b[31m", "\x1b[39m", color.Red, color.RedString},
		{"Green", "\x1b[32m", "\x1b[39m", color.Green, color.GreenString},
		{"Yellow", "\x1b[33m", "\x1b[39m", color.Yellow, color.YellowString},
		{"Blue", "\x1b[34m", "\x1b[39m", color.Blue, color.BlueString},
		{"Magenta", "\x1b[35m", "\x1b[39m", color.Magenta, color.MagentaString},
		{"Cyan", "\x1b[36m", "\x1b[39m", color.Cyan, color.CyanString},
		{"White", "\x1b[37m", "\x1b[39m", color.White, color.WhiteString},
		{"HiBlack", "\x1b[90m", "\x1b[39m", color.HiBlack, color.HiBlackString},
		{"HiRed", "\x1b[91m", "\x1b[39m", color.HiRed, color.HiRedString},
		{"HiGreen", "\x1b[92m", "\x1b[39m", color.HiGreen, color.HiGreenString},
		{"HiYellow", "\x1b[93m", "\x1b[39m", color.HiYellow, color.HiYellowString},
		{"HiBlue", "\x1b[94m", "\x1b[39m", color.HiBlue, color.HiBlueString},
		{"HiMagenta", "\x1b[95m", "\x1b[39m", color.HiMagenta, color.HiMagentaString},
		{"HiCyan", "\x1b[96m", "\x1b[39m", color.HiCyan, color.HiCyanString},
		{"HiWhite", "\x1b[97m", "\x1b[39m", color.HiWhite, color.HiWhiteString},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			tt.print("%s", "x")
			if got, want := buf.String(), tt.on+"x"+tt.off+"\n"; got != want {
				t.Fatalf("print = %q, want %q", got, want)
			}
			if got, want := tt.str("%s", "x"), tt.on+"x"+tt.off; got != want {
				t.Fatalf("string = %q, want %q", got, want)
			}
		})
	}
}

func TestHelpersArePrintfLike(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer
	useOutput(t, &buf)
	useNoColor(t, true)

	color.Red("100%%")
	if got, want := buf.String(), "100%\n"; got != want {
		t.Fatalf("Red(\"100%%%%\") = %q, want %q", got, want)
	}
	if got, want := color.RedString("100%%"), "100%"; got != want {
		t.Fatalf("RedString(\"100%%%%\") = %q, want %q", got, want)
	}
}

func TestHelpersAppendOneNewline(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer
	useOutput(t, &buf)
	useNoColor(t, true)

	color.Red("plain")
	if got, want := buf.String(), "plain\n"; got != want {
		t.Fatalf("Red without newline = %q, want %q", got, want)
	}

	buf.Reset()
	color.Red("plain\n")
	if got, want := buf.String(), "plain\n"; got != want {
		t.Fatalf("Red with newline = %q, want %q", got, want)
	}
}

func TestHelpersFormatArguments(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer
	useOutput(t, &buf)
	useNoColor(t, true)

	color.Red("boom %d", 1)
	if got, want := buf.String(), "boom 1\n"; got != want {
		t.Fatalf("Red with an argument = %q, want %q", got, want)
	}
}
