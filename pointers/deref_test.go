package pointers_test

import (
	"fmt"

	"github.com/carlmjohnson/truthy/pointers"
)

func ExampleCoalesce() {
	var np *int
	fmt.Println(pointers.Coalesce(np, 1))
	np = new(int)
	fmt.Println(pointers.Coalesce(np, 1))
	// Output:
	// 1
	// 0
}

func ExampleDeref() {
	var np *int
	fmt.Println(pointers.Deref(np))
	one := 1
	np = &one
	fmt.Println(pointers.Deref(np))
	// Output:
	// 0
	// 1
}
