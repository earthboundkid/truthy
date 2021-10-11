# Truthy

![Truthiness](https://user-images.githubusercontent.com/222245/136619462-f2bc5858-067f-4277-a813-b95c64b3cdac.png)

Truthy is a package which uses generics (Go 1.18+) to create useful boolean tests and helper functions.

## Examples

```
// truthy.Value returns the truthiness of any argument.
// If the value's type has a Bool() bool method, the method is called and returned.
// If the type has an IsZero() bool method, the opposite value is returned.
// Slices and maps are truthy if they have a length greater than zero.
// All other types are truthy if they are not their zero value.

truthy.Value(0) // false
truthy.Value(1) // true

truthy.Value("") // false
truthy.Value(" ") // true

truthy.Value([]byte(``)) // false
truthy.Value([]byte(` `)) // true

truthy.Value([]int{}) // false
truthy.Value([]int{1, 2, 3}) // true

var err error
truthy.Value(err) // false
truthy.Value(errors.New("hi")) // true
if truthy.Value(err) {
	panic(err)
}


var p *int
truthy.Value(p) // false

p = new(int)
// truthy does not check value underlying pointer!
truthy.Value(p) // true

// Ever wish Go had ? : ternary operators?
// Now it has a ternary function.
x := truthy.Cond("", 1, 10) // x == 10

// truthy.Cond cannot lazily evaluate its arguments,
// but you can use a closure to fake it.
s := truthy.Cond([]string{""},
	func() string {
		// do some calculation
		return "foo"
	},
	func() string {
		// do some calculation
		return "bar"
	})()
// s == "foo"


// How about an equivalent of the nullish coalescing operator ?? 
// as seen in C#, JavaScript, PHP, etc.:
var s string
truthy.First(s, "default") // "default"
s = "something"
truthy.First(s, "default") // "something"
truthy.First(0, 0*1, 1-1, 0x10-10) // 6


// Easily set defaults
n := getUserInput()
truthy.SetDefault(&n, 42)

if truthy.Or("1", 0) {
	fmt.Println("yay") // prints yay
}

if truthy.And(300, "") {
	fmt.Println("boo") // not executed
}
```

## Installation

As of October 2018, Go 1.18 is not released, but you can install Go tip with

```
$ go install golang.org/dl/gotip@latest
$ gotip download
$ gotip init me/myproject
$ gotip get github.com/carlmjohnson/truthy
```

## Discussion

Should you use this package? It's a little bit of a joke package, but the set default functionality is useful, especially for strings. Time will tell what best practices around the use of generics in Go turn out to be.
