DSV parser for Go
====

[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/angu1ss/dsv/blob/master/LICENSE)

DSV is delimiter-separated values parser for GO. It will parse lines and insert data into any type of struct. DSV supports both simple structs and structs with tagging.

```
go get github.com/angu1ss/dsv
```

Quickstart
--

DSV inserts data into struct by fields order.

```go
import (
    "fmt"
    "os"
    "testing"
    )

type TestRow struct {
  Name   string
  Age    int
  Gender string
  Active bool
}

func main() {

  file, _ := os.Open("example.tsv")
  defer file.Close()

  data := TestRow{}
  parser, _ := NewTsvParser(file, &data)

  for {
    eof, err := parser.Next()
    if eof {
      return
    }
    if err != nil {
      panic(err)
    }
    fmt.Println(data)
  }

}

```

You can define tags to struct fields to map values

```go
type TestRow struct {
  Name   string `dsv:"name"`
  Age    int    `dsv:"age"`
  Gender string `dsv:"gender"`
  Active bool   `dsv:"bool"`
}
```

Supported field types
--

Currently, library supports limited fields:

- string
- bool
- int, base 10
- int8, base 10
- int16, base 10
- int32, base 10
- int64, base 10
- unit, base 10
- unit8, base 10
- unit16, base 10
- unit32, base 10
- unit64, base 10
- float32, delimiter is point, comma will be ignored
- float64, delimiter is point, comma will be ignored
- complex64
- complex128

