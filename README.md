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


### Benchmark
Algorithm is specially fast when the number of buckets is small (ie bucketSize is big)

    -- Plain sorting
    BenchmarkSort200kSize5-4           	     100	  71359551 ns/op
    BenchmarkBuckets200knBuckets5-4    	     300	  24358098 ns/op
    BenchmarkBuckets200knBuckets16-4   	     200	  35642710 ns/op
    BenchmarkBuckets200knBuckets32-4   	     200	  40671283 ns/op

For example, for an array of 200k elements, organizing the elements into 5 buckets with Floyd Rivest is 3x times faster than plain sort.


### Acknowldegment
The library is based on the insight from [Mourner](https://github.com/mourner/rbush/blob/master/index.js#L547) that Floyd-Rivest performs better than plain sort
