/*
Package color provides ANSI escape code–based colored terminal output for Go.
The API can be used in several ways; pick the style that suits you.

# Simple helper functions

Use predefined foreground-color helpers that automatically append a newline:

	color.Cyan("Prints text in cyan.")
	color.Blue("Prints %s in blue.", "text")
	color.Red("We have red")
	color.Yellow("Yellow color too!")
	color.Magenta("And many others ..")

	// Hi-intensity variants
	color.HiGreen("Bright green color.")
	color.HiBlack("Bright black means gray..")
	color.HiWhite("Shiny white color!")

# Custom Color objects

For reusable, mixed-attribute styles:

	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")

	// Or chain attributes directly in New()
	c = color.New(color.Bold, color.FgGreen)
	c.Printf("Bold green: %s\n", "hello")

# 24-bit RGB colors

When your terminal supports true color:

	color.RGB(255, 128, 0).Println("foreground orange")
	color.BgRGB(0, 0, 128).Println("dark blue background")

# Func-returning variants (for reuse / embedding)

	red := color.New(color.FgRed).PrintfFunc()
	red("Error: %s\n", err)

	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("this is a %s\n", yellow("warning"))

# Writing to an io.Writer

	blue := color.New(color.FgBlue).FprintFunc()
	blue(myWriter, "blue notice")

# Plug into existing code (Set / Unset)

	color.Set(color.FgYellow)
	fmt.Println("This line is yellow")
	color.Unset()

# Disable / Enable color

Color output is automatically disabled when stdout is not a terminal, when
TERM=dumb, or when the NO_COLOR environment variable is set (see
https://no-color.org). You can also control it programmatically:

	color.NoColor = true  // disable globally

	c := color.New(color.FgCyan)
	c.DisableColor()
	c.Println("plain text")
	c.EnableColor()
	c.Println("cyan again")
*/
package color
