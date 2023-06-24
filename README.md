# testwell

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/rubrikinc/testwell/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/rubrikinc/testwell/assert?status.svg)](https://godoc.org/github.com/rubrikinc/testwell/assert)
[![Go Report Card](https://goreportcard.com/badge/rubrikinc/testwell)](http://goreportcard.com/report/rubrikinc/testwell)

A small set of type-safe convenient testing functions for Go.

Two identical packages (same API) are provided: `expect` and `assert`. Test
functions from the `expect` package will perform the test, and on failure, will
log the result, and mark the test case as failed. Test functions from the
`assert` package will do the same, and in addition, will fail the test case
right away.

## Example

```go
import (
    "github.com/rubrikinc/testwell/expect"
    "github.com/rubrikinc/testwell/assert"
    "testing"
)

func TestEx1(t *testing.T) {
    assert.True(t, 42 == 42, "this got to be true")

    expect.Equal(t, 42, 42)
    expect.NotEqual(t, int8(42), int64(42), "equality includes type!")
    expect.Contains(t, "foo", "zingfoobar", "optional %s", "fmt msg",)
}
```

## Documentation

Use `godoc`. For example:

```shell
godoc -http :9090
```

Then navigate to
[localhost](http://localhost:9090/pkg/github.com/rubrikinc/testwell/assert/).
Alternatively, can view them online
[here](https://godoc.org/github.com/rubrikinc/testwell/assert).

## Code overview

### Get involved

We are happy to receive bug reports, fixes, documentation enhancements, and
other improvements.

Please report bugs via the
[github issue tracker](https://github.com/rubrikinc/testwell/issues).

#### Building and testing

To build and test locally, you need the following dependencies:
 - Go Language: [installation instructions](https://go.dev/doc/install)
 - `goimports`
   ```shell
   # Install
   go install golang.org/x/tools/cmd/goimports@latest

   # You may need to add `goimports` to `PATH` after installation.
   export PATH="$HOME/go/bin:$PATH"
   ```

Before you can use `make`, be sure to have initialized the Go module.
```shell
go mod init github.com/rubrikinc/testwell
```

### Testwell design rules

 - Naming is important, and good names should be used.
 - We must use type-checking when possible even though testing is a runtime
   failure.
 - We must have consistent order of parameters. `*testing.T` first, left
   value, right value, etc...
 - We must print out easy to read test failure message. Nobody has time to
   waste deciphering them.

### Code generation

A small code generator can be found in `internal/codegen`. It takes a Go text
template file, and run it with a list of types. The unique template is
`internal/assert.tmpl`.

The goal of the template is easily generate functions parameterized with
different type. Go doesn't offer generics as of Go 1.9.

`internal/codegen/gen.sh` builds and run the code generator in a single pass.
(We cannot use `go run` because it doesn't support multiple files package).

The packages `assert`, `expect` and `internal/tests` all have a `lib.go` file
with a `go generate` comment for running `internal/codegen/gen.sh`. This
generates `assert.go` (`expect.go` for the package `expect`).

The generated code relies on finding the function
`func failTest(t testing.T, tf fail.TestFailure) bool` in the same package.
This function is responsible for logging out the failed test, and taking the
appropriate action (fail now, mark test as failed or log it for unit testing).

## Motivation

This is our own version of set of helpers function for testing in Go. This
might requires some explanation as to why we needed to invest some effort into
that.

### Testify is not the best choice

[Testify](https://github.com/stretchr/testify) has become more or less the
standard for convenient test functions. However it has a few flaws:

 - In Go, no two different types can be compared (with the notable exception of
   `nil`). Even something that looks trivial at first glance, comparing
   `int8(42)` to `int32(42)` is not possible in Go. If you are wondering why,
   look at the mess of automatic widening coercion rules in C/C++.
 
   Testify however allows comparing of different types to make writing the code
   more convenient. However, when the comparison fails due to mismatched types
   (but not values!), Testify usually doesn't provide enough contextual
   information on the type being compared, making it difficult to understand why
   `42 is not equal to 42`. One could be an `int32` and the other a `string`.

 - Arguments to the different helpers functions are not always consistent.
   Sometimes the expected value is first sometimes second.

 - Naming is poor. One example is the `Contains` function from Testify. Not
   only does it takes its argument in a reversed order compared to most other
   functions, it also has the variable `s` for the container and `contains` for
   the element to search within the container. Another example is the package
   `assert` which doesn't do assertion, but log failure and marks the test case
   as failed.  Which is contrary to what most programmers are used to. While
   the package `require` is doing failing assertion.

 - Go 1.9 improved the situations of test helper functions by adding a
   mechanism to mark them explicitly. We take advantage of this in this
   library.

## Licensing

This library is MIT-licensed.

## Copyright

Copyright © 2022 ([Rubrik Inc.](https://www.rubrik.com))
