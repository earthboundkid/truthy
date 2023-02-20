package truthy_test

import (
	"fmt"

	"github.com/carlmjohnson/truthy"
)

func ExampleCond_lazy() {
	i := 1
	// Cond cannot lazily evaluate its arguments,
	// but you can use a closure to fake it.
	s := truthy.Cond(
		truthy.Value(i),
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
