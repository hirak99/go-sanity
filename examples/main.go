package main

import (
	"fmt"

	. "github.com/hirak99/go-sanity"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	odds := Filter(nums, func(_, n int) bool { return n%2 == 1 })
	fmt.Println(odds)
}
