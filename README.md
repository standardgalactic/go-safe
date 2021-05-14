# go-safe

[![test](https://github.com/kenkyu392/go-safe/workflows/test/badge.svg?branch=master)](https://github.com/kenkyu392/go-safe)
[![codecov](https://codecov.io/gh/kenkyu392/go-safe/branch/master/graph/badge.svg)](https://codecov.io/gh/kenkyu392/go-safe)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-00ADD8?logo=go)](https://pkg.go.dev/github.com/kenkyu392/go-safe)
[![go report card](https://goreportcard.com/badge/github.com/kenkyu392/go-safe)](https://goreportcard.com/report/github.com/kenkyu392/go-safe)
[![license](https://img.shields.io/github/license/kenkyu392/go-safe)](LICENSE)

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
		// Do something...
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
```

## License

[MIT](LICENSE)
