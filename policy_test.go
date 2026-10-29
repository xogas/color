package color_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/xogas/color"
)

func TestEnabled(t *testing.T) {
	var buf bytes.Buffer
	tests := []struct {
		name    string
		env     map[string]string
		noColor bool
		want    bool
	}{
		{"buffer is not a terminal", nil, false, false},
		{"FORCE_COLOR", map[string]string{"FORCE_COLOR": "1"}, false, true},
		{"FORCE_COLOR=true", map[string]string{"FORCE_COLOR": "true"}, false, true},
		{"FORCE_COLOR beats NO_COLOR", map[string]string{"FORCE_COLOR": "1", "NO_COLOR": "1"}, false, true},
		{"NO_COLOR", map[string]string{"NO_COLOR": "1"}, false, false},
		{"NO_COLOR=false is still set", map[string]string{"NO_COLOR": "false"}, false, false},
		{"FORCE_COLOR=0 is off", map[string]string{"FORCE_COLOR": "0"}, false, false},
		{"FORCE_COLOR=false is off", map[string]string{"FORCE_COLOR": "false"}, false, false},
		{"FORCE_COLOR=no is off", map[string]string{"FORCE_COLOR": "no"}, false, false},
		{"FORCE_COLOR=off is off", map[string]string{"FORCE_COLOR": "OFF"}, false, false},
		{"TERM=dumb", map[string]string{"TERM": "dumb"}, false, false},
		{"TERM=DUMB is case insensitive", map[string]string{"TERM": "DUMB"}, false, false},
		{"NoColor beats force", map[string]string{"FORCE_COLOR": "1"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanEnv(t)
			for key, value := range tt.env {
				t.Setenv(key, value)
			}
			useNoColor(t, tt.noColor)
			if got := color.Enabled(&buf); got != tt.want {
				t.Fatalf("Enabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnabledIsEvaluatedPerCall(t *testing.T) {
	cleanEnv(t)
	var buf bytes.Buffer

	if color.Enabled(&buf) {
		t.Fatal("Enabled() = true before the environment changed")
	}
	t.Setenv("FORCE_COLOR", "1")
	if !color.Enabled(&buf) {
		t.Fatal("Enabled() = false, want the new environment to take effect")
	}
}

func TestEnabledNilWriter(t *testing.T) {
	cleanEnv(t)
	if color.Enabled(nil) {
		t.Error("Enabled(nil) = true, want false")
	}
}

func TestEnabledFileDescriptorWriter(t *testing.T) {
	cleanEnv(t)
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close() //nolint: errcheck
	defer w.Close() //nolint: errcheck

	if color.Enabled(w) {
		t.Error("a pipe is not a terminal")
	}
}

func TestColorOverride(t *testing.T) {
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

	t.Run("DisableColor beats a forced environment", func(t *testing.T) {
		t.Setenv("FORCE_COLOR", "1")
		if color.New(color.FgRed).DisableColor().Enabled() {
			t.Error("DisableColor should win over FORCE_COLOR")
		}
	})
	t.Run("EnableColor beats NoColor", func(t *testing.T) {
		useNoColor(t, true)
		if !color.New(color.FgRed).EnableColor().Enabled() {
			t.Error("EnableColor should win over NoColor")
		}
		if color.New(color.FgRed).Enabled() {
			t.Error("NoColor should disable a color without an override")
		}
	})
}
