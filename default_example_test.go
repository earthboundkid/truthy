package truthy_test

import (
	"fmt"
	"time"

	"github.com/carlmjohnson/truthy"
)

type MyStruct struct {
	Port    int
	Host    string
	Timeout time.Duration
}

func NewMyStruct(port int, host string, timeout time.Duration) *MyStruct {
	s := MyStruct{
		Port:    port,
		Host:    host,
		Timeout: timeout,
	}
	truthy.SetDefault(&s.Port, 80)
	truthy.SetDefault(&s.Host, "localhost")
	truthy.SetDefault(&s.Timeout, 10*time.Second)
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
		Port:    truthy.First(port, 80),
		Host:    truthy.First(host, "localhost"),
		Timeout: truthy.First(timeout, 10*time.Second),
	}
}

func ExampleFirst() {
	fmt.Println(MakeMyStruct(0, "", 0))
	fmt.Println(MakeMyStruct(443, "example.com", 1*time.Minute))
	// Output:
	// &{80 localhost 10s}
	// &{443 example.com 1m0s}
}
