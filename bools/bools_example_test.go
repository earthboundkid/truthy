package bools_test

import (
	"errors"
	"fmt"
	"time"

	"github.com/carlmjohnson/truthy/bools"
)

func ExampleBools() {
	var err error
	fmt.Println(bools.Comparable(err))

	err = errors.New("hi")
	fmt.Println(bools.Comparable(err))

	var n int
	fmt.Println(bools.Comparable(n))

	n = 1
	fmt.Println(bools.Comparable(n))

	var p *int
	fmt.Println(bools.Comparable(p))

	p = new(int)
	// Comparable does not check underlying pointer value!
	fmt.Println(bools.Comparable(p))
	// for that, use Pointer,
	fmt.Println(bools.Pointer(p))
	// but beware of double pointers!
	fmt.Println(bools.Pointer(&p))

	var s string
	fmt.Println(bools.Comparable(s))

	s = " "
	fmt.Println(bools.Comparable(s))

	var b []byte
	fmt.Println(bools.Slice(b))

	b = []byte(" ")
	fmt.Println(bools.Slice(b))

	m := map[string]string{}
	fmt.Println(bools.Map(m))

	m["a"] = "b"
	fmt.Println(bools.Map(m))

	var t time.Time
	t = t.Local()
	// t.IsZero() is still true although t is not the empty value
	fmt.Println(bools.Any(t))

	t = t.Add(1 * time.Second)
	fmt.Println(bools.Any(t))

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
