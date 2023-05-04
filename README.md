# rbarray &#x1f48e;
[![Go](https://github.com/ppmasa8/rbarray/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/ppmasa8/rbarray/actions/workflows/go.yml)

You can do ruby-like array operations in the golang.&#x1f60e;

https://docs.ruby-lang.org/ja/latest/class/Array.html

# Install
```zsh
$ go install github.com/ppmasa8/rbarray@v1.0.0
```

# Index
- Pop
- Shift
- Push
- Unshift
- Delete
- Sum
- Max
- Min
- Size
- Uniq
- Include
- First
- Last

coming soon....

# Usage
```go
package main

import (
        "fmt"
        "github.com/ppmasa8/rbarray"
)

func main() {
        array := rbarray.Array{IntVals: rbarray.IntArray{1, 3, 5}}
        fmt.Println(array.Pop())

// output
// 5
}
```
