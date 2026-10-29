package color

import (
	"fmt"
	"strings"
)

// Color is an immutable SGR style.
//
// Methods that change a color return a new value, so a Color can be copied
// freely and shared between goroutines. The zero Color is valid and applies no
// styling; it renders text unchanged.
type Color struct {
	on   string // CSI sequence that enables the style; "" for the zero Color
	off  string // CSI sequence that disables it again
	mode colorMode
}

// colorMode is a per-value override of the color policy.
type colorMode uint8

const (
	colorAuto   colorMode = iota // follow the package policy
	colorAlways                  // always emit escape sequences
	colorNever                   // never emit escape sequences
)

// New returns a color that enables attrs in order.
func New(attrs ...Attribute) Color { return Color{}.Add(attrs...) }

// RGB returns a 24 bit foreground color. Components are clamped to 0-255.
func RGB(r, g, b int) Color { return Color{}.AddRGB(r, g, b) }

// BgRGB returns a 24 bit background color. Components are clamped to 0-255.
func BgRGB(r, g, b int) Color { return Color{}.AddBgRGB(r, g, b) }

// Fg256 returns a foreground color taken from the 256 color palette.
func Fg256(n int) Color { return Color{}.AddFg256(n) }

// Bg256 returns a background color taken from the 256 color palette.
func Bg256(n int) Color { return Color{}.AddBg256(n) }

// Add returns c with attrs appended.
func (c Color) Add(attrs ...Attribute) Color {
	for _, a := range attrs {
		c = c.add(attrStep(a))
	}
	return c
}

// AddRGB returns c with a 24 bit foreground color appended.
func (c Color) AddRGB(r, g, b int) Color { return c.add(rgbStep(fgExt, fgReset, r, g, b)) }

// AddBgRGB returns c with a 24 bit background color appended.
func (c Color) AddBgRGB(r, g, b int) Color { return c.add(rgbStep(bgExt, bgReset, r, g, b)) }

// AddFg256 returns c with a foreground color from the 256 color palette appended.
func (c Color) AddFg256(n int) Color { return c.add(paletteStep(fgExt, fgReset, n)) }

// AddBg256 returns c with a background color from the 256 color palette appended.
func (c Color) AddBg256(n int) Color { return c.add(paletteStep(bgExt, bgReset, n)) }

// EnableColor returns c with color output forced on, overriding the package
// policy for this value.
func (c Color) EnableColor() Color { c.mode = colorAlways; return c }

// DisableColor returns c with color output forced off, overriding the package
// policy for this value.
func (c Color) DisableColor() Color { c.mode = colorNever; return c }

// Enabled reports whether c is allowed to emit escape sequences on Output. The
// zero Color never is, because it has nothing to emit.
func (c Color) Enabled() bool { return c.on != "" && c.enabledFor(Output) }

// SameStyle reports whether c and other enable and disable the same SGR
// parameters. Attribute order matters and EnableColor or DisableColor overrides
// are ignored; use == to compare the complete state of two Colors.
func (c Color) SameStyle(other Color) bool { return c.on == other.on && c.off == other.off }

// Sprint formats its operands and returns the colorized result. Color is
// applied only when c is enabled for Output.
func (c Color) Sprint(a ...any) string {
	return c.render(fmt.Sprint(a...), c.Enabled())
}

// Sprintf formats according to a format specifier and returns the colorized result.
func (c Color) Sprintf(format string, a ...any) string {
	return c.render(fmt.Sprintf(format, a...), c.Enabled())
}

// Sprintln is like Sprint but appends a newline after the reset sequence.
func (c Color) Sprintln(a ...any) string {
	return c.renderLine(strings.TrimSuffix(fmt.Sprintln(a...), "\n"), c.Enabled())
}

// add returns c with s appended. Reset codes are deduplicated because several
// attributes share the same reset, for example Bold and Faint both use 22.
func (c Color) add(s step) Color {
	c.on = appendSGR(c.on, s.on)
	if !hasParam(c.off, s.off) {
		c.off = appendSGR(c.off, s.off)
	}
	return c
}

// render wraps s in c's escape sequences when enabled is true. The zero Color,
// a disabled color and an empty string are returned unchanged.
func (c Color) render(s string, enabled bool) string {
	if !enabled || s == "" || c.on == "" {
		return s
	}
	return c.on + s + c.off
}

// renderLine is render for one line: the trailing newline is never wrapped in
// escape sequences.
func (c Color) renderLine(s string, enabled bool) string {
	return c.render(s, enabled) + "\n"
}
