# go-iowrap [![Build Status](https://travis-ci.org/LouisBrunner/go-iowrap.svg?branch=master)](https://travis-ci.org/LouisBrunner/go-iowrap) [![Coverage Status](https://coveralls.io/repos/github/LouisBrunner/go-iowrap/badge.svg?branch=master)](https://coveralls.io/github/LouisBrunner/go-iowrap?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/LouisBrunner/go-iowrap)](https://goreportcard.com/report/github.com/LouisBrunner/go-iowrap) [![GoDoc](https://godoc.org/github.com/LouisBrunner/go-iowrap?status.svg)](https://godoc.org/github.com/LouisBrunner/go-iowrap)
Simple gopher client (in Go)

# Install it

```
go get -u github.com/LouisBrunner/go-iowrap
```

# Usage

```go
import (
  "fmt"

  "github.com/LouisBrunner/go-iowrap"
)

func main() {
  content := iowrap.WrapStdout(func() {
    fmt.Printf("123\n")
  })
  if string(content) != "123\n" {
    panic("failed!")
  }
}

```
