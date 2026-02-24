package color_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xogas/color"
)

func TestNewWithNoColorEnv(t *testing.T) {
	tests := []struct {
		name    string
		envVal  string
		wantESC bool
	}{
		{"NO_COLOR set", "1", false},
		{"NO_COLOR absent", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("NO_COLOR", tt.envVal)

			prev := color.NoColor
			color.NoColor = false
			defer func() { color.NoColor = prev }()

			got := color.New(color.FgRed).Sprint("hello")
			hasESC := strings.Contains(got, "\x1b[")
			if hasESC != tt.wantESC {
				t.Errorf("Sprint() = %q, wantESC=%v", got, tt.wantESC)
			}
		})
	}
}

func TestRGBAndBgRGB(t *testing.T) {
	// visual only – just make sure these don't panic
	tests := []struct{ r, g, b int }{
		{255, 128, 0},
		{230, 42, 42},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			_, _ = color.RGB(tt.r, tt.g, tt.b).Println("foreground")
			_, _ = color.BgRGB(tt.r, tt.g, tt.b).AddBgRGB(0, 0, 0).Println("with background")
			_, _ = color.BgRGB(tt.r, tt.g, tt.b).Println("background")
			_, _ = color.RGB(tt.r, tt.g, tt.b).AddRGB(255, 255, 255).Println("with foreground")
		})
	}
}

func TestEquals(t *testing.T) {
	tests := []struct {
		name string
		a, b *color.Color
		want bool
	}{
		{
			"both nil",
			nil, nil,
			true,
		},
		{
			"one nil",
			color.New(color.FgRed), nil,
			false,
		},
		{
			"nil receiver",
			nil, color.New(color.FgRed),
			false,
		},
		{
			"same single attr",
			color.New(color.FgRed), color.New(color.FgRed),
			true,
		},
		{
			"different single attr",
			color.New(color.FgRed), color.New(color.FgBlue),
			false,
		},
		{
			"same multiple attrs",
			color.New(color.FgRed, color.Bold), color.New(color.FgRed, color.Bold),
			true,
		},
		{
			"different length",
			color.New(color.FgRed), color.New(color.FgRed, color.Bold),
			false,
		},
		{
			"same attrs different order",
			color.New(color.FgRed, color.Bold), color.New(color.Bold, color.FgRed),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Equals(tt.b)
			if got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnformatUsesSpecificResetCode(t *testing.T) {
	tests := []struct {
		name      string
		attr      color.Attribute
		wantClose string
	}{
		{"Bold uses ResetBold(22)", color.Bold, "\x1b[22m"},
		{"Italic uses ResetItalic(23)", color.Italic, "\x1b[23m"},
		{"Underline uses ResetUnderline(24)", color.Underline, "\x1b[24m"},
		{"FgRed uses generic Reset(0)", color.FgRed, "\x1b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := color.New(tt.attr)
			c.EnableColor()
			got := c.Sprint("x")
			if !strings.HasSuffix(got, tt.wantClose) {
				t.Errorf("Sprint() = %q, want closing sequence %q", got, tt.wantClose)
			}
		})
	}
}
