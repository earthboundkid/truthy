package bools_test

import (
	"fmt"

	"github.com/carlmjohnson/truthy/bools"
)

func ExampleTernary_lazy() {
	i := 1
	// Ternary cannot lazily evaluate its arguments,
	// but you can use a closure to fake it.
	s := bools.Ternary(
		bools.Comparable(i),
		func() string {
			// do some calculation
			return "true"
		},
		func() string {
			// do some calculation
			return "false"
		})()
	fmt.Printf("%q\n", s)
	// Output:
	// "true"
}
