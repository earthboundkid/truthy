package truthy_test

import (
	"fmt"

	"github.com/carlmjohnson/truthy"
)

func ExampleCoalesce() {
	var np *int
	fmt.Println(truthy.Coalesce(np, 1))
	np = new(int)
	fmt.Println(truthy.Coalesce(np, 1))
	// Output:
	// 1
	// 0
}

func ExampleDeref() {
	var np *int
	fmt.Println(truthy.Deref(np))
	one := 1
	np = &one
	fmt.Println(truthy.Deref(np))
	// Output:
	// 0
	// 1
}
