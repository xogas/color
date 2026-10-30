# AGENTS.md

## Project

ANSI SGR terminal color library for Go. Feature-frozen: changes are bug fixes
and refactors only. Adding or changing public API requires discussion first --
tests and docs form a contract that must be kept in sync.

## Layer discipline (required)

- `attribute.go` -- pure data and encoding. No IO, no global state.
- `color.go` -- immutable `Color` values. Every "modifying" method returns a new
  value; never mutate the receiver.
- `policy.go` -- the only file that decides color on/off.
- `writer.go` -- the only file that performs stream IO.
- `sugar.go` -- pure forwarding to `Color`; no new logic.

## Invariants (pinned by tests; changing one means updating its test)

- `Color` fields `on`/`off` are complete CSI sequences; the zero `Color` is
  valid and produces no output.
- Unknown `Attribute` values are written verbatim and reset with SGR 0.
- Reset codes shared by multiple attributes are deduplicated (`hasParam`).
- Newlines always live outside the escape sequences.
- `NO_COLOR` is non-empty-value based (`NO_COLOR=false` still disables color);
  `FORCE_COLOR` is boolean parsed (`0`/`false`/`no`/`off` are false).
- The policy is re-evaluated on every call. Never snapshot environment
  variables at init time.

## Testing

- External test package `color_test`; shared helpers live in `color_test.go`
  and must not be redefined elsewhere.
- The oracle for expected byte sequences is the ECMA-48 spec, not the current
  implementation's output.
- Global state (`Output`, `NoColor`) is only touched through the
  borrow-and-restore helpers `useOutput` / `useNoColor`.
- Any test touching environment variables calls `cleanEnv(t)` first.

## Documentation

- Every exported symbol has a doc comment.
- Contract statements live in three places that must stay in sync when changed:
  code comments, `doc.go`, and `README.md`.
