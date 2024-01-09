package defaults_test

import (
	"fmt"
	"time"

	"github.com/carlmjohnson/truthy/defaults"
)

type MyStruct struct {
	Port    int
	Host    string
	Timeout time.Duration
}

func NewMyStruct(port int, host string, timeout time.Duration) *MyStruct {
	s := MyStruct{port, host, timeout}
	defaults.SetFirst(&s.Port, 80)
	defaults.SetFirst(&s.Host, "localhost")
	defaults.SetFirst(&s.Timeout, 10*time.Second)
	return &s
}

func ExampleSetDefault() {
	fmt.Println(NewMyStruct(0, "", 0))
	fmt.Println(NewMyStruct(443, "example.com", 1*time.Minute))
	// Output:
	// &{80 localhost 10s}
	// &{443 example.com 1m0s}
}

func MakeMyStruct(port int, host string, timeout time.Duration) *MyStruct {
	return &MyStruct{
		Port:    defaults.GetFirst(port, 80),
		Host:    defaults.GetFirst(host, "localhost"),
		Timeout: defaults.GetFirst(timeout, 10*time.Second),
	}
}

func ExampleFirst() {
	fmt.Println(MakeMyStruct(0, "", 0))
	fmt.Println(MakeMyStruct(443, "example.com", 1*time.Minute))
	// Output:
	// &{80 localhost 10s}
	// &{443 example.com 1m0s}
}
