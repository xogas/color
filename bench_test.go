package color_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/xogas/color"
)

var (
	stringSink string
	intSink    int
	boolSink   bool
	colorSink  color.Color
)

func BenchmarkSprintColored(b *testing.B) {
	c := color.New(color.Bold, color.FgHiRed).EnableColor()
	b.ReportAllocs()
	for b.Loop() {
		stringSink = c.Sprint("hello world")
	}
}

func BenchmarkSprintPlain(b *testing.B) {
	c := color.New(color.Bold, color.FgHiRed).DisableColor()
	b.ReportAllocs()
	for b.Loop() {
		stringSink = c.Sprint("hello world")
	}
}

func BenchmarkFprint(b *testing.B) {
	c := color.New(color.Bold, color.FgHiRed).EnableColor()
	b.ReportAllocs()
	for b.Loop() {
		intSink, _ = c.Fprint(io.Discard, "hello world")
	}
}

func BenchmarkPolicy(b *testing.B) {
	var buf bytes.Buffer
	b.ReportAllocs()
	for b.Loop() {
		boolSink = color.Enabled(&buf)
	}
}

func BenchmarkBuildColor(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		colorSink = color.New(color.Bold, color.FgRed)
	}
}

func BenchmarkBuildRGB(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		colorSink = color.RGB(255, 128, 0)
	}
}
