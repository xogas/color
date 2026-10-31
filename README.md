# color

Small ANSI SGR color library for Go. A `Color` is an immutable style: build it,
then render it to a string or write it to any `io.Writer`. The only direct
dependency is [`golang.org/x/term`](https://pkg.go.dev/golang.org/x/term).

## Install

```sh
go get github.com/xogas/color
```

Requires Go 1.27 or later.

## Usage

```go
// Package-level helpers write to color.Output and append a newline.
color.Red("build failed: %s", err)

// A Color is a value, so it can be built once and reused.
danger := color.New(color.Bold, color.FgHiRed)
msg := danger.Sprint("boom")

// Write to any writer; color is emitted only for a terminal.
color.New(color.FgYellow).Fprintf(os.Stderr, "warn: %d\n", n)

// 24 bit and 256 color palettes.
color.RGB(255, 128, 0).Println("orange")
color.Bg256(236).Println("dark panel")
```

Methods that change a color return a new value, so a `Color` can be copied and
shared between goroutines without locking:

```go
var danger = color.New(color.Bold, color.FgHiRed) // safe to share

red := color.New(color.FgRed)
boldRed := red.Add(color.Bold) // red is unchanged
```

The package-level helpers are printf-like: escape a literal percent sign as
`%%`. A trailing newline is appended unless the format already ends with one.

## When is color emitted

The decision combines the process environment with the target writer. Switches
are consulted in this order:

| Priority | Switch | Effect |
| --- | --- | --- |
| 1 | `EnableColor` / `DisableColor` | per value override |
| 2 | `NoColor` | disables everything |
| 3 | `FORCE_COLOR` | a true value forces color, a false value disables it |
| 4 | `NO_COLOR` | any non-empty value disables color |
| 5 | `TERM=dumb` | disables color |
| 6 | target writer is a terminal | otherwise plain text |

`0`, `false`, `no` and `off` (any case) are false values; any other non-empty
value is true. `NO_COLOR` is presence based, so `NO_COLOR=false` still disables
color, as [no-color.org](https://no-color.org) specifies. A per-value override
wins over the environment. `color.Enabled(w)` reports the package-level decision
for an explicit writer.

The policy is evaluated on every call, so environment changes take effect
immediately. Because the terminal check is applied to the writer, writing into a
`bytes.Buffer` or a regular file stays free of escape sequences unless color is
forced explicitly. Tests do not need a pty.

`color.Output` and `color.NoColor` are global; do not mutate them while other
goroutines are printing.

## Design

| File | Domain | Responsibility |
| --- | --- | --- |
| `attribute.go` | encoding | `Attribute` constants, SGR parameter encoding, reset codes |
| `color.go` | style | `Color`, constructors, `Add`, `EnableColor`, `Sprint` |
| `policy.go` | policy | `NoColor`, `Enabled`, environment and terminal detection |
| `writer.go` | output | `Output`, `Fprint`, `Print`, `Set`/`Unset`, streaming |
| `sugar.go` | sugar | `Red`, `RedString` and friends |
| `doc.go` | docs | package documentation |

`attribute.go` is pure data and encoding with no IO or global state. `writer.go`
is the only file that performs IO, `policy.go` owns the color decision, and
`color.go` composes the two. Tests and benchmarks mirror these domains, and the
examples run as documentation tests.

## Platform support

Any platform where `golang.org/x/term` reports a terminal: Linux, macOS, the
BSDs and Windows. On Windows the console must support ANSI escape sequences
(Windows 10 or later, or Windows Terminal); legacy conhost may need virtual
terminal processing enabled. No build tags, no separate Windows console adapter.

## License

MIT - see [LICENSE](./LICENSE) for details.
