## Floyd Rivest

This library implements the [Floyd Rivest](https://en.wikipedia.org/wiki/Floyd%E2%80%93Rivest_algorithm) algorithm.

### Installation

    go get github.com/furstenheim/FloydRivest

### Import

    import (
        "github.com/furstenheim/FloydRivest"
    )

### Methods


    s := // some slice
    FloydRivest.Buckets(sort.Interface(s), 5)

    // s is now sorted into buckets of size 5
    // max(s[0:5]) < min(s[5:10])
    // max(s[10: 15]) < min(s[15:20])
    // ...



