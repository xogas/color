package color

import (
	"fmt"
	"strings"
)

// Styles used by the package level helpers.
var (
	blackStyle     = New(FgBlack)
	redStyle       = New(FgRed)
	greenStyle     = New(FgGreen)
	yellowStyle    = New(FgYellow)
	blueStyle      = New(FgBlue)
	magentaStyle   = New(FgMagenta)
	cyanStyle      = New(FgCyan)
	whiteStyle     = New(FgWhite)
	hiBlackStyle   = New(FgHiBlack)
	hiRedStyle     = New(FgHiRed)
	hiGreenStyle   = New(FgHiGreen)
	hiYellowStyle  = New(FgHiYellow)
	hiBlueStyle    = New(FgHiBlue)
	hiMagentaStyle = New(FgHiMagenta)
	hiCyanStyle    = New(FgHiCyan)
	hiWhiteStyle   = New(FgHiWhite)
)

// printWith formats text with c and writes it to Output followed by a newline.
// The newline is written after the reset sequence.
func printWith(c Color, format string, a ...any) {
	text := strings.TrimSuffix(fmt.Sprintf(format, a...), "\n")
	_, _ = c.writeLine(writerOr(Output), text)
}

// Black prints in black foreground and appends a newline.
func Black(format string, a ...any) { printWith(blackStyle, format, a...) }

// Red prints in red foreground and appends a newline.
func Red(format string, a ...any) { printWith(redStyle, format, a...) }

// Green prints in green foreground and appends a newline.
func Green(format string, a ...any) { printWith(greenStyle, format, a...) }

// Yellow prints in yellow foreground and appends a newline.
func Yellow(format string, a ...any) { printWith(yellowStyle, format, a...) }

// Blue prints in blue foreground and appends a newline.
func Blue(format string, a ...any) { printWith(blueStyle, format, a...) }

// Magenta prints in magenta foreground and appends a newline.
func Magenta(format string, a ...any) { printWith(magentaStyle, format, a...) }

// Cyan prints in cyan foreground and appends a newline.
func Cyan(format string, a ...any) { printWith(cyanStyle, format, a...) }

// White prints in white foreground and appends a newline.
func White(format string, a ...any) { printWith(whiteStyle, format, a...) }

// HiBlack prints in high intensity black foreground and appends a newline.
func HiBlack(format string, a ...any) { printWith(hiBlackStyle, format, a...) }

// HiRed prints in high intensity red foreground and appends a newline.
func HiRed(format string, a ...any) { printWith(hiRedStyle, format, a...) }

// HiGreen prints in high intensity green foreground and appends a newline.
func HiGreen(format string, a ...any) { printWith(hiGreenStyle, format, a...) }

// HiYellow prints in high intensity yellow foreground and appends a newline.
func HiYellow(format string, a ...any) { printWith(hiYellowStyle, format, a...) }

// HiBlue prints in high intensity blue foreground and appends a newline.
func HiBlue(format string, a ...any) { printWith(hiBlueStyle, format, a...) }

// HiMagenta prints in high intensity magenta foreground and appends a newline.
func HiMagenta(format string, a ...any) { printWith(hiMagentaStyle, format, a...) }

// HiCyan prints in high intensity cyan foreground and appends a newline.
func HiCyan(format string, a ...any) { printWith(hiCyanStyle, format, a...) }

// HiWhite prints in high intensity white foreground and appends a newline.
func HiWhite(format string, a ...any) { printWith(hiWhiteStyle, format, a...) }

// BlackString returns the formatted text in black foreground.
func BlackString(format string, a ...any) string { return blackStyle.Sprintf(format, a...) }

// RedString returns the formatted text in red foreground.
func RedString(format string, a ...any) string { return redStyle.Sprintf(format, a...) }

// GreenString returns the formatted text in green foreground.
func GreenString(format string, a ...any) string { return greenStyle.Sprintf(format, a...) }

// YellowString returns the formatted text in yellow foreground.
func YellowString(format string, a ...any) string { return yellowStyle.Sprintf(format, a...) }

// BlueString returns the formatted text in blue foreground.
func BlueString(format string, a ...any) string { return blueStyle.Sprintf(format, a...) }

// MagentaString returns the formatted text in magenta foreground.
func MagentaString(format string, a ...any) string { return magentaStyle.Sprintf(format, a...) }

// CyanString returns the formatted text in cyan foreground.
func CyanString(format string, a ...any) string { return cyanStyle.Sprintf(format, a...) }

// WhiteString returns the formatted text in white foreground.
func WhiteString(format string, a ...any) string { return whiteStyle.Sprintf(format, a...) }

// HiBlackString returns the formatted text in high intensity black foreground.
func HiBlackString(format string, a ...any) string { return hiBlackStyle.Sprintf(format, a...) }

// HiRedString returns the formatted text in high intensity red foreground.
func HiRedString(format string, a ...any) string { return hiRedStyle.Sprintf(format, a...) }

// HiGreenString returns the formatted text in high intensity green foreground.
func HiGreenString(format string, a ...any) string { return hiGreenStyle.Sprintf(format, a...) }

// HiYellowString returns the formatted text in high intensity yellow foreground.
func HiYellowString(format string, a ...any) string { return hiYellowStyle.Sprintf(format, a...) }

// HiBlueString returns the formatted text in high intensity blue foreground.
func HiBlueString(format string, a ...any) string { return hiBlueStyle.Sprintf(format, a...) }

// HiMagentaString returns the formatted text in high intensity magenta foreground.
func HiMagentaString(format string, a ...any) string { return hiMagentaStyle.Sprintf(format, a...) }

// HiCyanString returns the formatted text in high intensity cyan foreground.
func HiCyanString(format string, a ...any) string { return hiCyanStyle.Sprintf(format, a...) }

// HiWhiteString returns the formatted text in high intensity white foreground.
func HiWhiteString(format string, a ...any) string { return hiWhiteStyle.Sprintf(format, a...) }
