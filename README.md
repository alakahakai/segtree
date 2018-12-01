A segment tree implementation in Go with update function

It supports different data types by using generic interface{}; It provides a custom combine function, so it can support different use cases, such as sum, product, max, min, or any others; It uses a specified emptyValue - not always 0, e.g. for product, it should be 1.

To build with debug output, use: go build -tags debug
