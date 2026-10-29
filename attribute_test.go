package color_test

import (
	"testing"

	"github.com/xogas/color"
)

func TestSequences(t *testing.T) {
	tests := []struct {
		name    string
		color   color.Color
		on, off string
	}{
		{"attribute", color.New(color.FgRed), "\x1b[31m", "\x1b[39m"},
		{"combination", color.New(color.Bold, color.FgHiCyan), "\x1b[1;96m", "\x1b[22;39m"},
		{"rgb", color.RGB(255, 0, 128), "\x1b[38;2;255;0;128m", "\x1b[39m"},
		{"background rgb", color.BgRGB(0, 0, 0), "\x1b[48;2;0;0;0m", "\x1b[49m"},
		{"foreground palette", color.Fg256(196), "\x1b[38;5;196m", "\x1b[39m"},
		{"background palette", color.Bg256(236), "\x1b[48;5;236m", "\x1b[49m"},
		{"clamped low", color.RGB(-5, 300, 12), "\x1b[38;2;0;255;12m", "\x1b[39m"},
		{"palette clamped low", color.Fg256(-1), "\x1b[38;5;0m", "\x1b[39m"},
		{"palette clamped high", color.Bg256(999), "\x1b[48;5;255m", "\x1b[49m"},
		{"empty", color.New(), "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := render(tt.color, "x"), tt.on+"x"+tt.off; got != want {
				t.Fatalf("render = %q, want %q", got, want)
			}
		})
	}
}

func TestResetCodes(t *testing.T) {
	tests := []struct {
		name    string
		attr    color.Attribute
		on, off string
	}{
		{"Reset", color.Reset, "\x1b[0m", "\x1b[0m"},
		{"Bold", color.Bold, "\x1b[1m", "\x1b[22m"},
		{"Faint", color.Faint, "\x1b[2m", "\x1b[22m"},
		{"Italic", color.Italic, "\x1b[3m", "\x1b[23m"},
		{"Underline", color.Underline, "\x1b[4m", "\x1b[24m"},
		{"BlinkSlow", color.BlinkSlow, "\x1b[5m", "\x1b[25m"},
		{"BlinkRapid", color.BlinkRapid, "\x1b[6m", "\x1b[25m"},
		{"ReverseVideo", color.ReverseVideo, "\x1b[7m", "\x1b[27m"},
		{"Concealed", color.Concealed, "\x1b[8m", "\x1b[28m"},
		{"CrossedOut", color.CrossedOut, "\x1b[9m", "\x1b[29m"},
		{"FgBlack", color.FgBlack, "\x1b[30m", "\x1b[39m"},
		{"FgWhite", color.FgWhite, "\x1b[37m", "\x1b[39m"},
		{"BgBlack", color.BgBlack, "\x1b[40m", "\x1b[49m"},
		{"BgWhite", color.BgWhite, "\x1b[47m", "\x1b[49m"},
		{"FgHiBlack", color.FgHiBlack, "\x1b[90m", "\x1b[39m"},
		{"FgHiWhite", color.FgHiWhite, "\x1b[97m", "\x1b[39m"},
		{"BgHiBlack", color.BgHiBlack, "\x1b[100m", "\x1b[49m"},
		{"BgHiWhite", color.BgHiWhite, "\x1b[107m", "\x1b[49m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := render(color.New(tt.attr), "x"), tt.on+"x"+tt.off; got != want {
				t.Fatalf("render = %q, want %q", got, want)
			}
		})
	}
}

func TestUnknownAttributeUsesFullReset(t *testing.T) {
	got := render(color.New(color.Attribute(26)), "x")
	if want := "\x1b[26mx\x1b[0m"; got != want {
		t.Fatalf("render = %q, want %q", got, want)
	}
}

func TestAdders(t *testing.T) {
	c := color.New(color.Bold).
		AddRGB(1, 2, 3).
		AddFg256(196).
		AddBgRGB(4, 5, 6).
		AddBg256(7)
	want := "\x1b[1;38;2;1;2;3;38;5;196;48;2;4;5;6;48;5;7mx\x1b[22;39;49m"
	if got := render(c, "x"); got != want {
		t.Fatalf("render = %q, want %q", got, want)
	}
}

func TestResetDeduplication(t *testing.T) {
	tests := []struct {
		name  string
		color color.Color
		want  string
	}{
		{"shared reset", color.New(color.Bold, color.Faint), "\x1b[1;2mx\x1b[22m"},
		{"two foregrounds", color.New(color.FgRed, color.FgBlue), "\x1b[31;34mx\x1b[39m"},
		{"two backgrounds", color.New(color.BgRed, color.BgBlue), "\x1b[41;44mx\x1b[49m"},
		{
			"extended colors",
			color.New(color.FgRed).AddRGB(1, 2, 3).AddFg256(7),
			"\x1b[31;38;2;1;2;3;38;5;7mx\x1b[39m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := render(tt.color, "x"); got != tt.want {
				t.Fatalf("render = %q, want %q", got, tt.want)
			}
		})
	}
}
