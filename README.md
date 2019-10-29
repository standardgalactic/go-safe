# go-safe

[![GoDoc](https://godoc.org/github.com/kenkyu392/go-safe?status.svg)](https://godoc.org/github.com/kenkyu392/go-safe)
[![Go Report Card](https://goreportcard.com/badge/github.com/kenkyu392/go-safe)](https://goreportcard.com/report/github.com/kenkyu392/go-safe)

Provides functionality to help you execute methods safely.

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

[MIT](https://github.com/kenkyu392/go-safe/blob/master/LICENSE)
