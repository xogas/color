package color_test

import (
	"bytes"
	"fmt"
	"io"

	"github.com/xogas/color"
)

func ExampleNew() {
	c := color.New(color.Bold, color.FgRed)
	fmt.Printf("%q\n", c.EnableColor().Sprint("x"))
	fmt.Printf("%q\n", c.DisableColor().Sprint("x"))
	// Output:
	// "\x1b[1;31mx\x1b[22;39m"
	// "x"
}

func ExampleColor_Fprintln() {
	var buf bytes.Buffer
	_, _ = color.New(color.FgYellow).EnableColor().Fprintln(&buf, "x:", 3)
	fmt.Printf("%q\n", buf.String())
	// Output: "\x1b[33mx: 3\x1b[39m\n"
}

func ExampleColor_Sprintf() {
	fmt.Printf("%q\n", color.New(color.FgGreen).EnableColor().Sprintf("ok %d", 42))
	// Output: "\x1b[32mok 42\x1b[39m"
}

func ExampleRGB() {
	fmt.Printf("%q\n", color.RGB(255, 128, 0).EnableColor().Sprint("x"))
	// Output: "\x1b[38;2;255;128;0mx\x1b[39m"
}

func ExampleRed() {
	prevOutput, prevNoColor := color.Output, color.NoColor
	var buf bytes.Buffer
	color.Output, color.NoColor = &buf, true
	defer func() { color.Output, color.NoColor = prevOutput, prevNoColor }()

	color.Red("build failed: %d", 1)
	fmt.Print(buf.String())
	// Output: build failed: 1
}

func ExampleEnabled() {
	prev := color.NoColor
	color.NoColor = true
	defer func() { color.NoColor = prev }()

	fmt.Println(color.Enabled(io.Discard))
	// Output: false
}
