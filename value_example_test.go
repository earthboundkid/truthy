package truthy_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/carlmjohnson/truthy"
)

func ExampleValue() {
	var err error
	fmt.Println(truthy.Value(err))

	err = errors.New("hi")
	fmt.Println(truthy.Value(err))

	var n int
	fmt.Println(truthy.Value(n))

	n = 1
	fmt.Println(truthy.Value(n))

	var p *int
	fmt.Println(truthy.Value(p))

	p = new(int)
	// truthy does not check value underlying pointer!
	fmt.Println(truthy.Value(p))

	var s string
	fmt.Println(truthy.Value(s))

	s = " "
	fmt.Println(truthy.Value(s))

	m := map[string]string{}
	fmt.Println(truthy.Value(m))

	m["a"] = "b"
	fmt.Println(truthy.Value(m))

	var t time.Time
	t = t.Local()
	// t.IsZero() is still true although t is not the empty value
	fmt.Println(truthy.Value(t))

	t = t.Add(1 * time.Second)
	fmt.Println(truthy.Value(t))

	// Output:
	// false
	// true
	// false
	// true
	// false
	// true
	// false
	// true
	// false
	// true
	// false
	// true
}
