package is_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/carlmjohnson/truthy/is"
)

func ExampleTruthy() {
	var err error
	fmt.Println(is.Truthy(err))

	err = errors.New("hi")
	fmt.Println(is.Truthy(err))

	var n int
	fmt.Println(is.Truthy(n))

	n = 1
	fmt.Println(is.Truthy(n))

	var p *int
	fmt.Println(is.Truthy(p))

	p = new(int)
	// Truthy does not check underlying pointer value!
	fmt.Println(is.Truthy(p))
	// for that, use TruthyPointer,
	fmt.Println(is.TruthyPointer(p))
	// but beware of double pointers!
	fmt.Println(is.TruthyPointer(&p))

	var s string
	fmt.Println(is.Truthy(s))

	s = " "
	fmt.Println(is.Truthy(s))

	var b []byte
	fmt.Println(is.TruthySlice(b))

	b = []byte(" ")
	fmt.Println(is.TruthySlice(b))

	m := map[string]string{}
	fmt.Println(is.TruthyMap(m))

	m["a"] = "b"
	fmt.Println(is.TruthyMap(m))

	var t time.Time
	t = t.Local()
	// t.IsZero() is still true although t is not the empty value
	fmt.Println(is.TruthyAny(t))

	t = t.Add(1 * time.Second)
	fmt.Println(is.TruthyAny(t))

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
	// false
	// true
	// false
	// true
}
