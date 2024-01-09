package pointers_test

import (
	"fmt"

	"github.com/carlmjohnson/truthy/pointers"
)

func ExampleGetFirst() {
	type config struct{ string }
	userInput := func() *config {
		return nil
	}
	someConfig := pointers.GetFirst(
		userInput(),
		&config{"default config"},
	)
	fmt.Println(someConfig)
	// Output:
	// &{default config}
}
