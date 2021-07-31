# envim - Go environment variables for human

![Build status](https://github.com/cuonglm/envim/actions/workflows/ci.yaml/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/cuonglm/envim)](https://goreportcard.com/report/github.com/cuonglm/envim)
[![GoDoc](https://godoc.org/github.com/cuonglm/envim?status.svg)](https://godoc.org/github.com/cuonglm/envim)

# Why envim?

Just another (better) way for managing environment variables from Go code, but written for human.

# Installation
```sh
go get -u github.com/cuonglm/envim
```

# Usage

```go
package main

import (
	"fmt"

	env "github.com/cuonglm/envim"
)

func main() {

	// Clean environment variables
	env.Clear()

	// Set a variable
	_ = env.Set("foo", "foo")
	_ = env.Set("fooo", "")
	_ = env.Set("GOPATH", "/home/cuonglm/go")
	_ = env.Set("GOROOT", "/home/cuonglm/sources/go")

	// Get a variable
	fmt.Println(env.Get("foo"))
	fmt.Println(env.Get("fooo"))

	// Check environment variable is set
	fmt.Println(env.IsSet("foo"))
	fmt.Println(env.IsSet("NotSetVar"))

	// Unset a variable
	_ = env.Unset("fooo")

	// Get all variables into map
	fmt.Printf("%+v\n", env.Map())

	// Like Map(), but variable with prefix only
	fmt.Printf("%+v\n", env.MapWithPrefix("GO"))

	// Update environment variables from a map
	// skipped invalid
	m := map[string]string{"bar": "bar", "=": "equal"}
	env.FromMap(m)
	fmt.Printf("%+v\n", env.Map())
}
```

Run that file give you:

```sh
$ go run envim_example.go
foo

true
false
map[foo:foo GOPATH:/home/cuonglm/go GOROOT:/home/cuonglm/sources/go]
map[GOROOT:/home/cuonglm/sources/go GOPATH:/home/cuonglm/go]
map[foo:foo GOPATH:/home/cuonglm/go GOROOT:/home/cuonglm/sources/go bar:bar]
```

# Author

Cuong Manh Le <cuong.manhle.vn@gmail.com>

# License

See [LICENSE](https://github.com/cuonglm/envim/blob/master/LICENSE)
