package pointers_test

import (
	"fmt"

	"github.com/carlmjohnson/truthy/pointers"
)

func ExampleNew() {
	strptr1 := pointers.New("meaning of life")
	strptr2 := pointers.New("meaning of life")
	fmt.Println(strptr1 != strptr2)
	fmt.Println(*strptr1 == *strptr2)

	intp1 := pointers.New(42)
	intp2 := pointers.New(42)
	fmt.Println(intp1 != intp2)
	fmt.Println(*intp1 == *intp2)

	type MyFloat float64
	fp := pointers.New[MyFloat](42)
	fmt.Println(fp != nil)
	fmt.Println(*fp == 42)

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
}
