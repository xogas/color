package color

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/term"
)

var (
	// NoColor controls whether color output is enabled. It is automatically set
	// to true when stdout is not a terminal, when TERM=dumb or when NO_COLOR is
	// set. Override it at any time to force-enable or force-disable colors.
	NoColor = noColorIsSet() || os.Getenv("TERM") == "dumb" || !isTerminal()

	// Output is the writer used by all package-level Print functions. Defaults
	// to os.Stdout.
	Output io.Writer = os.Stdout

	// Error is a color-supporting writer for os.Stderr.
	Error io.Writer = os.Stderr

	// colorsCache reduces allocations by reusing Color objects keyed by a
	// single Attribute (used by the helper functions Black/Red/etc.).
	colorsCache   = make(map[Attribute]*Color)
	colorsCacheMu sync.Mutex
)

// Color defines a custom color object composed of one or more SGR parameters.
type Color struct {
	params []Attribute
	// noColor overrides the global NoColor if non-nil.
	noColor *bool
}

// New returns a new Color with the given SGR attributes.
func New(value ...Attribute) *Color {
	c := &Color{params: make([]Attribute, 0)}
	if noColorIsSet() {
		c.noColor = new(true)
	}
	c.Add(value...)
	return c
}

// RGB returns a new foreground color using 24-bit true color (r, g, b in 0-255).
func RGB(r, g, b int) *Color {
	return New(foreground, 2, Attribute(r), Attribute(g), Attribute(b))
}

// BgRGB returns a new background color using 24-bit true color (r, g, b in 0-255).
func BgRGB(r, g, b int) *Color {
	return New(background, 2, Attribute(r), Attribute(g), Attribute(b))
}

// Add appends SGR attributes to an existing Color. Returns the receiver for
// chaining.
func (c *Color) Add(value ...Attribute) *Color {
	c.params = append(c.params, value...)
	return c
}

// AddRGB appends a 24-bit foreground color to an existing Color. Returns the
// receiver for chaining.
func (c *Color) AddRGB(r, g, b int) *Color {
	return c.Add(foreground, 2, Attribute(r), Attribute(g), Attribute(b))
}

// AddBgRGB appends a 24-bit background color to an existing Color. Returns the
// receiver for chaining.
func (c *Color) AddBgRGB(r, g, b int) *Color {
	return c.Add(background, 2, Attribute(r), Attribute(g), Attribute(b))
}

// DisableColor disables color output for this Color instance.
func (c *Color) DisableColor() { c.noColor = new(true) }

// EnableColor re-enables color output for this Color instance.
func (c *Color) EnableColor() { c.noColor = new(false) }

// Equals reports whether two Colors have identical SGR parameters.
func (c *Color) Equals(c2 *Color) bool {
	if c == nil && c2 == nil {
		return true
	}
	if c == nil || c2 == nil {
		return false
	}
	if len(c.params) != len(c2.params) {
		return false
	}
	for _, attr := range c.params {
		if !slices.Contains(c2.params, attr) {
			return false
		}
	}
	return true
}

func (c *Color) isNoColorSet() bool {
	if c.noColor != nil {
		return *c.noColor
	}
	return NoColor
}

// sequence returns the semicolon-joined SGR parameters, e.g. "1;36".
func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}
	return strings.Join(format, ";")
}

// format returns the opening SGR escape sequence.
func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

// unformat returns the closing/reset SGR escape sequence.
func (c *Color) unformat() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(Reset))
		if ra, ok := mapResetAttributes[v]; ok {
			format[i] = strconv.Itoa(int(ra))
		}
	}
	return fmt.Sprintf("%s[%sm", escape, strings.Join(format, ";"))
}

// wrap surrounds s with the opening and closing SGR sequences.
func (c *Color) wrap(s string) string {
	if c.isNoColorSet() {
		return s
	}
	return c.format() + s + c.unformat()
}

// noColorIsSet reports whether the NO_COLOR environment variable is set to a
// non-empty string. See https://no-color.org.
func noColorIsSet() bool {
	return os.Getenv("NO_COLOR") != ""
}

// isTerminal reports whether os.Stdout is connected to a terminal.
func isTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}
