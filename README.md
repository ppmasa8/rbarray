# rbarray
It can handle arrays like ruby.

# Install
```zsh
$ go install github.com/ppmasa8/rbarr@v0.1.1
```

# Usage
```go
package main

import (
        "fmt"
        "github.com/ppmasa8/rbarray"
)

func main() {
        fmt.Println("Hello, 世界")
        array := rbarray.Array{IntVals: rbarray.IntArray{1, 3, 5}}
        fmt.Println(array.Pop())
        
// output
// Hello, 世界
// 5
}
```
