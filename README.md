# Go Sanity

This is a collection of methods and classes that make it easier to use Golang out of the box.

## Installing
From the root of your package, run this to add a dependency -
```bash
go get -d github.com/hirak99/go-sanity@latest
```

## Using

It is recommended to import go-sanity in the default namespace.

Example -

```golang
package main

import (
  "fmt"

  "github.com/hirak99/go-sanity"
)

func main() {
  nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

  odds := sanity.Filter(nums, func(n int) bool { return n % 2 == 1 })
  fmt.Println(odds)
}
```

## Content

### Methods

| Function      | Example                                                                       |
| ------------- | ----------------------------------------------------------------------------- |
| Ternary if    | `If(a == b, "yes", "no")`                                                     |
| Map           | `Map([]string{"12", "34"}, func(s string) int { return strconv.Atoi(s) })`    |
| Reduce        | `Reduce([]string{"12", "34"}, "", func(x, y string) string { return x + y })` |
| Filter        | `Filter([]int{10, 33, 59, 93}, func(i int) bool { return i > 50 }`            |
| Sum           | `Sum([]float64{1.1, 2.3, 4.8})`                                               |
| ChanToSlice   | `ChanToSlice(c)`                                                              |
| SaneSortSlice | `SaneSortSlice(s, func(a, b string) { return a < b})`                         |

### Collections
| Function | Example                                                         |
| -------- | --------------------------------------------------------------- |
| Set      | `s := MakeSet[int]()`, then use `s.Add(...)`, `s.Has(...)` etc. |

### Testing

| Function                  | Example                                      |
| ------------------------- | -------------------------------------------- |
| Assert                    | `Assert(t, something_that_should_be_true())` |
| AssertEqual               | `Assert(t, a, b)`                            |
| AssertSliceEqual          | `AssertSliceEqual(t, s1, s2)`                |
| AssertSliceEqualUnordered | `AssertSliceEqualUnordered(t, s1, s2)`       |