package condition_test

import (
	"fmt"

	"github.com/carlmjohnson/truthy/condition"
	"github.com/carlmjohnson/truthy/is"
)

func ExampleEvaluate_lazy() {
	i := 1
	// Cond cannot lazily evaluate its arguments,
	// but you can use a closure to fake it.
	s := condition.Evaluate(
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
