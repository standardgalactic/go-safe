# go-safe

[![test](https://github.com/kenkyu392/go-safe/workflows/test/badge.svg?branch=main)](https://github.com/kenkyu392/go-safe)
[![codecov](https://codecov.io/gh/kenkyu392/go-safe/branch/main/graph/badge.svg)](https://codecov.io/gh/kenkyu392/go-safe)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-00ADD8?logo=go)](https://pkg.go.dev/github.com/kenkyu392/go-safe)
[![go report card](https://goreportcard.com/badge/github.com/kenkyu392/go-safe)](https://goreportcard.com/report/github.com/kenkyu392/go-safe)
[![license](https://img.shields.io/github/license/kenkyu392/go-safe)](LICENSE)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

This Go package provides a sandbox for the safe execution of panic-inducing programs.

This package is intended for use in development and unit testing, and is not recommended for use in production.

## Installation

```
go get -u github.com/kenkyu392/go-safe
```

## Usage

```go
package main

import (
	"log"

	"github.com/kenkyu392/go-safe"
)

func main() {
	// The panic that occurs in the function is automatically handled as an error.
	if err := safe.Do(func() error {
		// Something that might cause a panic...
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
```

The sandbox function can be used in combination with sync/errgroup.

```go
package main

import (
	"log"

	"github.com/kenkyu392/go-safe"
	"golang.org/x/sync/errgroup"
)

func main() {
	// Create a sandboxed function.
	fn := safe.Func(func() error {
		// Something that might cause a panic...
		return nil
	})

	eg := new(errgroup.Group)
	for i := 0; i < 10; i++ {
		eg.Go(fn)
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
```

## License

[MIT](LICENSE)
