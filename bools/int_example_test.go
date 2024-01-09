package bools_test

import (
	"fmt"

	"github.com/carlmjohnson/truthy/bools"
)

func ExampleInt() {
	fmt.Println(bools.Int(true))
	fmt.Println(bools.Int(false))
	fmt.Println(bools.Int(bools.Comparable("")))
	fmt.Println(bools.Int('a' < 'b'))
	// Output:
	// 1
	// 0
	// 0
	// 1
}
