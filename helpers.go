package color

import "strings"

// getCachedColor returns (or creates and caches) a *Color for a single
// Attribute. This avoids allocating a new Color for every helper call.
func getCachedColor(p Attribute) *Color {
	colorsCacheMu.Lock()
	defer colorsCacheMu.Unlock()
	c, ok := colorsCache[p]
	if !ok {
		c = New(p)
		colorsCache[p] = c
	}
	return c
}

func colorPrint(format string, p Attribute, a ...any) {
	c := getCachedColor(p)
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	if len(a) == 0 {
		_, _ = c.Print(format)
	} else {
		_, _ = c.Printf(format, a...)
	}
}

func colorString(format string, p Attribute, a ...any) string {
	c := getCachedColor(p)
	if len(a) == 0 {
		return c.SprintFunc()(format)
	}
	return c.SprintfFunc()(format, a...)
}

// Black prints with black foreground; a newline is appended automatically.
func Black(format string, a ...any) { colorPrint(format, FgBlack, a...) }

// Red prints with red foreground; a newline is appended automatically.
func Red(format string, a ...any) { colorPrint(format, FgRed, a...) }

// Green prints with green foreground; a newline is appended automatically.
func Green(format string, a ...any) { colorPrint(format, FgGreen, a...) }

// Yellow prints with yellow foreground; a newline is appended automatically.
func Yellow(format string, a ...any) { colorPrint(format, FgYellow, a...) }

// Blue prints with blue foreground; a newline is appended automatically.
func Blue(format string, a ...any) { colorPrint(format, FgBlue, a...) }

// Magenta prints with magenta foreground; a newline is appended automatically.
func Magenta(format string, a ...any) { colorPrint(format, FgMagenta, a...) }

// Cyan prints with cyan foreground; a newline is appended automatically.
func Cyan(format string, a ...any) { colorPrint(format, FgCyan, a...) }

// White prints with white foreground; a newline is appended automatically.
func White(format string, a ...any) { colorPrint(format, FgWhite, a...) }

// BlackString returns a string with black foreground.
func BlackString(format string, a ...any) string { return colorString(format, FgBlack, a...) }

// RedString returns a string with red foreground.
func RedString(format string, a ...any) string { return colorString(format, FgRed, a...) }

// GreenString returns a string with green foreground.
func GreenString(format string, a ...any) string { return colorString(format, FgGreen, a...) }

// YellowString returns a string with yellow foreground.
func YellowString(format string, a ...any) string { return colorString(format, FgYellow, a...) }

// BlueString returns a string with blue foreground.
func BlueString(format string, a ...any) string { return colorString(format, FgBlue, a...) }

// MagentaString returns a string with magenta foreground.
func MagentaString(format string, a ...any) string { return colorString(format, FgMagenta, a...) }

// CyanString returns a string with cyan foreground.
func CyanString(format string, a ...any) string { return colorString(format, FgCyan, a...) }

// WhiteString returns a string with white foreground.
func WhiteString(format string, a ...any) string { return colorString(format, FgWhite, a...) }

// HiBlack prints with hi-intensity black foreground; a newline is appended.
func HiBlack(format string, a ...any) { colorPrint(format, FgHiBlack, a...) }

// HiRed prints with hi-intensity red foreground; a newline is appended.
func HiRed(format string, a ...any) { colorPrint(format, FgHiRed, a...) }

// HiGreen prints with hi-intensity green foreground; a newline is appended.
func HiGreen(format string, a ...any) { colorPrint(format, FgHiGreen, a...) }

// HiYellow prints with hi-intensity yellow foreground; a newline is appended.
func HiYellow(format string, a ...any) { colorPrint(format, FgHiYellow, a...) }

// HiBlue prints with hi-intensity blue foreground; a newline is appended.
func HiBlue(format string, a ...any) { colorPrint(format, FgHiBlue, a...) }

// HiMagenta prints with hi-intensity magenta foreground; a newline is appended.
func HiMagenta(format string, a ...any) { colorPrint(format, FgHiMagenta, a...) }

// HiCyan prints with hi-intensity cyan foreground; a newline is appended.
func HiCyan(format string, a ...any) { colorPrint(format, FgHiCyan, a...) }

// HiWhite prints with hi-intensity white foreground; a newline is appended.
func HiWhite(format string, a ...any) { colorPrint(format, FgHiWhite, a...) }

// HiBlackString returns a string with hi-intensity black foreground.
func HiBlackString(format string, a ...any) string { return colorString(format, FgHiBlack, a...) }

// HiRedString returns a string with hi-intensity red foreground.
func HiRedString(format string, a ...any) string { return colorString(format, FgHiRed, a...) }

// HiGreenString returns a string with hi-intensity green foreground.
func HiGreenString(format string, a ...any) string { return colorString(format, FgHiGreen, a...) }

// HiYellowString returns a string with hi-intensity yellow foreground.
func HiYellowString(format string, a ...any) string { return colorString(format, FgHiYellow, a...) }

// HiBlueString returns a string with hi-intensity blue foreground.
func HiBlueString(format string, a ...any) string { return colorString(format, FgHiBlue, a...) }

// HiMagentaString returns a string with hi-intensity magenta foreground.
func HiMagentaString(format string, a ...any) string { return colorString(format, FgHiMagenta, a...) }

// HiCyanString returns a string with hi-intensity cyan foreground.
func HiCyanString(format string, a ...any) string { return colorString(format, FgHiCyan, a...) }

// HiWhiteString returns a string with hi-intensity white foreground.
func HiWhiteString(format string, a ...any) string { return colorString(format, FgHiWhite, a...) }
