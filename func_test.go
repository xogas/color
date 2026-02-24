package color_test

import (
	"bytes"
	"testing"

	"github.com/xogas/color"
)

func TestPrintFunc(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf

	fn := newNoColor(color.Bold).PrintFunc()
	fn("hello", " ", "world")

	got := buf.String()
	want := "hello world"
	if got != want {
		t.Errorf("PrintFunc() wrote %q, want %q", got, want)
	}
}

func TestPrintfFunc(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf

	fn := newNoColor(color.Underline).PrintfFunc()
	fn("x=%d", 7)

	got := buf.String()
	want := "x=7"
	if got != want {
		t.Errorf("PrintfFunc() wrote %q, want %q", got, want)
	}
}

func TestPrintlnFunc(t *testing.T) {
	var buf bytes.Buffer
	color.Output = &buf

	fn := newNoColor(color.FgWhite).PrintlnFunc()
	fn("hello", " ", "world")

	got := buf.String()
	want := "hello world\n"
	if got != want {
		t.Errorf("PrintlnFunc() wrote %q, want %q", got, want)
	}
}

func TestSprintFunc(t *testing.T) {
	fn := newNoColor(color.FgRed).SprintFunc()
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
		got := fn(tt.args...)
		if got != tt.want {
			t.Errorf("SprintFunc()(%v) = %q, want %q", tt.args, got, tt.want)
		}
	}
}

func TestSprintfFunc(t *testing.T) {
	fn := newNoColor(color.FgGreen).SprintfFunc()
	tests := []struct {
		format string
		args   []any
		want   string
	}{
		{"hello %s", []any{"world"}, "hello world"},
		{"%d + %d = %d", []any{1, 2, 3}, "1 + 2 = 3"},
		{"no args", []any{}, "no args"},
	}
	for _, tt := range tests {
		got := fn(tt.format, tt.args...)
		if got != tt.want {
			t.Errorf("SprintfFunc()(%q, %v) = %q, want %q", tt.format, tt.args, got, tt.want)
		}
	}
}

func TestSprintlnFunc(t *testing.T) {
	fn := newNoColor(color.FgBlue).SprintlnFunc()
	tests := []struct {
		args []any
		want string
	}{
		{[]any{"hello"}, "hello\n"},
		{[]any{"hello", " ", "world"}, "hello world\n"},
		{[]any{}, "\n"},
	}
	for _, tt := range tests {
		got := fn(tt.args...)
		if got != tt.want {
			t.Errorf("SprintlnFunc()(%v) = %q, want %q", tt.args, got, tt.want)
		}
	}
}

func TestFprintFunc(t *testing.T) {
	fn := newNoColor(color.FgYellow).FprintFunc()
	var buf bytes.Buffer
	fn(&buf, "hello", " ", "world")
	got := buf.String()
	want := "hello world"
	if got != want {
		t.Errorf("FprintFunc() wrote %q, want %q", got, want)
	}
}

func TestFprintfFunc(t *testing.T) {
	fn := newNoColor(color.FgCyan).FprintfFunc()
	var buf bytes.Buffer
	fn(&buf, "value=%d", 99)
	got := buf.String()
	want := "value=99"
	if got != want {
		t.Errorf("FprintfFunc() wrote %q, want %q", got, want)
	}
}

func TestFprintlnFunc(t *testing.T) {
	fn := newNoColor(color.FgMagenta).FprintlnFunc()
	var buf bytes.Buffer
	fn(&buf, "hello", " ", "world")
	got := buf.String()
	want := "hello world\n"
	if got != want {
		t.Errorf("FprintlnFunc() wrote %q, want %q", got, want)
	}
}
