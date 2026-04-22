# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```shell
# Build, generate, and test everything
make

# Run all tests
go test -vet off ./internal/tests/

# Run a single test
go test -vet off -run TestEqual ./internal/tests/

# Regenerate assert.go, expect.go, and internal/tests/tests.go from the template
make generate

# Compile-check public packages without running tests
make testbuild

# Remove all generated files
make clean
```

`goimports` must be on `PATH` (`go install golang.org/x/tools/cmd/goimports@latest`).

Tests always require `-vet off` because they intentionally pass mismatched types (e.g. `int8` vs `int32`) to exercise the library's type-safety checks, which triggers go vet warnings.

## Architecture

The library exposes two public packages with identical APIs:

- **`assert/`** — calls `t.FailNow()` on failure (stops the test immediately)
- **`expect/`** — calls `t.Fail()` on failure (marks failed but continues)

Both packages are **entirely code-generated** from a single template:

```
internal/assert.tmpl  →  codegen  →  assert/assert.go
                                  →  expect/expect.go
                                  →  internal/tests/tests.go
```

The code generator (`internal/codegen/`) applies `internal/assert.tmpl` (a `text/template`) once per target package, pipes the output through `goimports`, and writes `<pkg>.go`. `gen.sh` builds the codegen binary and runs it in one step; each `lib.go` has the `//go:generate` directive that invokes it.

The only code that differs between `assert`, `expect`, and the internal test package is the hand-written `lib.go` in each, which defines:

```go
func failTest(t testing.T, tf fail.TestFailure) bool
```

- `assert/lib.go` → calls `t.Log(...)` then `t.FailNow()`
- `expect/lib.go` → calls `t.Log(...)` then `t.Fail()`
- `internal/tests/lib.go` → records `tf` into a `tWrapper` struct without calling `FailNow` or `Fail`

### How the library tests itself

`internal/tests/` wraps `*testing.T` in a `tWrapper` that captures `fail.TestFailure` objects rather than aborting the test. Tests call assertions on `wrap(t)` and then inspect `wt.LastFailure()` to verify that a failure occurred (or didn't). This allows testing that assertions fail correctly without aborting the Go test runner.

### Public API

Functions generated from the template, where `<Type>` is a capitalized type name (e.g. `Int32`, `String`, `Bool`). The untyped `interface{}` variant uses the bare name.

| Bare (interface{}) | Typed variants |
|----|---|
| `Equal` / `NotEqual` | `Equal<Type>` / `NotEqual<Type>` |
| `Contains` / `NotContains` | `Contains<Type>` / `NotContains<Type>` |
| `EqualTypes` / `NotEqualTypes` | — |
| `Nil` / `NotNil` | — |
| `Empty` / `NotEmpty` | — |
| `DeepEqual` / `NotDeepEqual` | — |
| `True` / `False` | — |
| `Failed` | — |
| `Panics` / `PanicsWith` | — |

`Contains` takes `(t, element, container)` — element first, container second. This is the reverse of Testify and is intentional (see README).

## Design rules

- Type-safe: no cross-type comparisons (e.g. `int8(42)` and `int32(42)` are not equal)
- Parameter order: `testing.T` first, then left/expected value, right/actual value, then optional `msg ...interface{}`
- Template type functions (`NumericalTypes`, `PrimitiveTypes`, `ErrorType`, `GenericType`) can be piped together with `|` to form type set unions

## What not to edit directly

`assert/assert.go`, `expect/expect.go`, and `internal/tests/tests.go` are generated — edit `internal/assert.tmpl` and re-run `make generate`.
