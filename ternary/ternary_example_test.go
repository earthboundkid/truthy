package ternary_test

import (
	"fmt"

	"github.com/carlmjohnson/truthy/is"
	"github.com/carlmjohnson/truthy/ternary"
)

func ExampleEvaluate_lazy() {
	i := 1
	// Evaluate cannot lazily evaluate its arguments,
	// but you can use a closure to fake it.
	s := ternary.Evaluate(
		is.Truthy(i),
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
