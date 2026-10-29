package color

import (
	"strconv"
	"strings"
)

// Attribute is a single SGR (Select Graphic Rendition) parameter code.
//
// The constants below are the attributes whose reset codes this package knows.
// Any other value is written to the terminal verbatim and reset with the
// all-attributes reset code (0).
type Attribute int

// Base attributes.
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Reset counterparts of the base attributes.
const (
	ResetBold Attribute = iota + 22
	ResetItalic
	ResetUnderline
	ResetBlinking
	_ // 26 is reserved by ECMA-48
	ResetReversed
	ResetConcealed
	ResetCrossedOut
)

// Foreground colors.
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Background colors.
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// High intensity foreground colors.
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// High intensity background colors.
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// escape is the CSI (Control Sequence Introducer) prefix of every SGR sequence
// produced by this package.
const escape = "\x1b["

// sgrEnd terminates an SGR sequence.
const sgrEnd = "m"

// Extended color selectors and the reset codes that undo them.
const (
	fgExt   = 38
	bgExt   = 48
	fgReset = 39
	bgReset = 49
)

// step holds one SGR parameter list that enables a style and the list that
// disables exactly that style again.
type step struct {
	on  string // e.g. "1" or "38;2;255;0;0"
	off string // e.g. "22" or "39"
}

// attrStep returns the step for a plain attribute.
func attrStep(a Attribute) step {
	return step{on: strconv.Itoa(int(a)), off: strconv.Itoa(resetCode(a))}
}

// rgbStep returns the step for a 24 bit color, for example
// on="38;2;255;0;0", off="39".
func rgbStep(ext, reset, r, g, b int) step {
	return step{
		on: strconv.Itoa(ext) + ";2;" +
			strconv.Itoa(clamp(r)) + ";" +
			strconv.Itoa(clamp(g)) + ";" +
			strconv.Itoa(clamp(b)),
		off: strconv.Itoa(reset),
	}
}

// paletteStep returns the step for a 256 color palette entry, for example
// on="38;5;196", off="39".
func paletteStep(ext, reset, n int) step {
	return step{
		on:  strconv.Itoa(ext) + ";5;" + strconv.Itoa(clamp(n)),
		off: strconv.Itoa(reset),
	}
}

// appendSGR appends one SGR parameter list to a complete sequence. A zero seq
// becomes the sequence that enables params.
func appendSGR(seq, params string) string {
	if seq == "" {
		return escape + params + sgrEnd
	}
	return strings.TrimSuffix(seq, sgrEnd) + ";" + params + sgrEnd
}

// hasParam reports whether the complete SGR sequence seq already contains
// param as one of its semicolon separated parameter lists.
func hasParam(seq, param string) bool {
	params := strings.TrimSuffix(strings.TrimPrefix(seq, escape), sgrEnd)
	for params != "" {
		p, rest, found := strings.Cut(params, ";")
		if p == param {
			return true
		}
		if !found {
			return false
		}
		params = rest
	}
	return false
}

// resetCode returns the SGR parameter that undoes a.
func resetCode(a Attribute) int {
	switch {
	case a >= FgBlack && a <= FgWhite, a >= FgHiBlack && a <= FgHiWhite:
		return fgReset
	case a >= BgBlack && a <= BgWhite, a >= BgHiBlack && a <= BgHiWhite:
		return bgReset
	}
	switch a {
	case Bold, Faint:
		return int(ResetBold)
	case Italic:
		return int(ResetItalic)
	case Underline:
		return int(ResetUnderline)
	case BlinkSlow, BlinkRapid:
		return int(ResetBlinking)
	case ReverseVideo:
		return int(ResetReversed)
	case Concealed:
		return int(ResetConcealed)
	case CrossedOut:
		return int(ResetCrossedOut)
	default:
		return int(Reset)
	}
}

// clamp limits v to the 0-255 range used by SGR color parameters.
func clamp(v int) int { return min(max(v, 0), 255) }
