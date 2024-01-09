package bools_test

import (
	"fmt"
	"time"

	"github.com/carlmjohnson/truthy/bools"
)

type MyStruct struct {
	Port    int
	Host    string
	Timeout time.Duration
}

func NewMyStruct(port int, host string, timeout time.Duration) *MyStruct {
	s := MyStruct{port, host, timeout}
	bools.SetFirst(&s.Port, 80)
	bools.SetFirst(&s.Host, "localhost")
	bools.SetFirst(&s.Timeout, 10*time.Second)
	return &s
}

func ExampleSetFirst() {
	fmt.Println(NewMyStruct(0, "", 0))
	fmt.Println(NewMyStruct(443, "example.com", 1*time.Minute))
	// Output:
	// &{80 localhost 10s}
	// &{443 example.com 1m0s}
}

func MakeMyStruct(port int, host string, timeout time.Duration) *MyStruct {
	return &MyStruct{
		Port:    bools.GetFirst(port, 80),
		Host:    bools.GetFirst(host, "localhost"),
		Timeout: bools.GetFirst(timeout, 10*time.Second),
	}
}

func ExampleGetFirst() {
	fmt.Println(MakeMyStruct(0, "", 0))
	fmt.Println(MakeMyStruct(443, "example.com", 1*time.Minute))
	// Output:
	// &{80 localhost 10s}
	// &{443 example.com 1m0s}
}
