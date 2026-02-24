# Color

Color provides colorized terminal output in Go using ANSI escape codes.

## Install

```bash
go get github.com/xogas/color
```

## Examples

Standard colors

```go
color.Cyan("Prints text in cyan.")
color.Blue("Prints %s in blue.", "text")
color.Red("We have red")
color.Yellow("Yellow color too!")
color.Magenta("And many others ..")

// Hi-intensity variants
color.HiGreen("Bright green color.")
color.HiBlack("Bright black means gray..")
color.HiWhite("Shiny white color!")
```

RGB colors

```go
color.RGB(255, 128, 0).Println("foreground orange")
color.RGB(230, 42, 42).Println("foreground red")

color.BgRGB(255, 128, 0).Println("background orange")
color.BgRGB(230, 42, 42).Println("background red")
```

Func-returning variants

```go
red := color.New(color.FgRed).PrintfFunc()
red("Error: %s\n", err)

// Mix up multiple attributes
notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
notice("Don't forget this...")
```

Writing to an io.Writer

```go
color.New(color.FgBlue).Fprintln(myWriter, "blue color!")

blue := color.New(color.FgBlue).FprintFunc()
blue(myWriter, "blue notice")
```

Plug into existing code

```go
color.Set(color.FgYellow)
fmt.Println("This line is yellow")
color.Unset()
```

Disable / Enable color

```go
color.NoColor = true  // disable globally

c := color.New(color.FgCyan)
c.DisableColor()
c.Println("plain text")
c.EnableColor()
c.Println("cyan again")
```

## Licence

The MIT License (MIT) - see [license](./license) for more details.
