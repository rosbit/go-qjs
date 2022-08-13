# go-qjs, makes QuickJS be embedded easily

[QuickJS](https://bellard.org/quickjs/) is a small and embeddable Javascript engine written by [Fabrice Bellard](https://bellard.org).

This package is intended to provide a wrapper to interact `QuickJS` with application written in golang.
With some helper functions, `go-qjs` makes it simple to calle QuickJS from Golang, and `go-qjs` can be
treated as an embeddable JavaScript.

### Install

The package is fully go-getable, So, just type

  `go get github.com/rosbit/go-qjs`

to install.

### Usage

Suppose there's a Javascript file named `a.js` like this:

```javascript
function add(a, b) {
    return a+b
}
```

one can call the Javascript function `add()` in Go code like the following:

```go
package main

import (
  "github.com/rosbit/go-qjs"
  "fmt"
)

var add func(int, int)int

func main() {
  ctx, err := qjs.NewQuickJS("/path/to/quickjs-exe/qjs", "a.js")
  if err != nil {
     fmt.Printf("%v\n", err)
     return
  }
  defer ctx.Quit()

  // method 1: bind JS function with a golang var
  if err := ctx.BindFunc("add", &add); err != nil {
     fmt.Printf("%v\n", err)
     return
  }
  res := add(1, 2)

  // method 2: call JS function using Call
  res, err := ctx.CallFunc("add", 1, 2)
  if err != nil {
     fmt.Printf("%v\n", err)
     return
  }

  fmt.Println("result is:", res)
}
```

### Status

The package is not fully tested, so be careful.

### Contribution

Pull requests are welcome! Also, if you want to discuss something send a pull request with proposal and changes.

__Convention:__ fork the repository and make changes on your fork in a feature branch.
