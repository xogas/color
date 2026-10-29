// Package color renders ANSI SGR (Select Graphic Rendition) escape sequences to
// strings and writes them through an io.Writer.
//
// A Color is an immutable style. Methods such as Add and EnableColor return a
// new Color, so a value can be copied, stored and shared between goroutines:
//
//	red := color.New(color.FgRed)
//	bold := red.Add(color.Bold) // red is unchanged
//	fmt.Print(bold.Sprint("hi"))
//
// Package-level helpers print to Output and append a newline:
//
//	color.Red("build failed: %s", err)
//
// # Color policy
//
// Whether escape sequences are emitted is decided by the destination writer
// together with the process environment. The switches are consulted in this
// order, highest priority first:
//
//  1. EnableColor or DisableColor on the Color value
//  2. NoColor
//  3. FORCE_COLOR: a true value forces color, a false value disables it
//  4. NO_COLOR set to a non-empty value
//  5. TERM=dumb
//  6. the writer exposes a terminal file descriptor
//
// "0", "false", "no" and "off" (any case) are false values; any other non-empty
// value is true. NO_COLOR is presence based, so NO_COLOR=false still disables
// color as https://no-color.org specifies.
//
// The policy is evaluated on every call, so environment changes take effect
// immediately. Because the terminal check looks at the writer, writing into a
// bytes.Buffer or a regular file produces plain text unless color is forced
// explicitly. Enabled reports the package level part of the decision. NoColor
// and Output are global; mutating them while other goroutines print is a data
// race.
package color
