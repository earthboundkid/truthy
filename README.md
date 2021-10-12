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

// Collection testing and filtering
truthy.Any(0, 1, 2) // true
truthy.All(0, 1, 2) // false

ss := []string{"", "a", "b", ""}
truthy.Filter(&ss) // ss == []string{"a", "b"}

// Logical operators
if truthy.Or("1", 0) {
	fmt.Println("yay") // prints yay
}

if truthy.And(300, "") {
	fmt.Println("boo") // not executed
}
```

## Installation

As of October 2021, Go 1.18 is not released, but you can install Go tip with

```
$ go install golang.org/dl/gotip@latest
$ gotip download
$ gotip init me/myproject
$ gotip get github.com/carlmjohnson/truthy
```

## FAQs

### Oh god

This is the correct reaction.

### Isn't this just using reflection? Does it even really require generics?

I tried to write a non-generic version of this package first, but you can’t reflect on an interface type. When you do `reflect.Value(x)`, you lose the fact that x was, e.g. an error, because `reflect.Value()` only takes `interface{}` and the conversion loses the interface type. You’d end up saying whether the underlying concrete type was empty or not, which is typically not what you want. To work around that, you could require that everything is passed as a pointer, e.g. `reflect.Value(&err)`, but `truthy.Value(&err)` sucks as an API. If you look at how `truthy.Value()` works, it accepts a value of type `T`, and then passes `*T` to `reflect.Value()` and calls `value.Elem()` to finally get the correct reflection type. So, on a technical level, you couldn’t quite make this API work without generics, although it could be close. However, `truthy.Filter()`, `truthy.SetDefault()`, `truthy.Any()`, and `truthy.All()` could be implemented with pure reflection, although the implementation would be a lot uglier.

Then there’s `truthy.First()`. To be honest, `truthy.First()` is the only part of the package that I consider actually useful, and even that, I mostly expect it to be used for picking a string or default. Anyhow, it requires generics to avoid the cast back from interface type to the concrete type.

### Should I use this package? 
Probably not. It's a little bit of a joke package, but the `truthy.First()` and `truthy.SetDefault()` functionality seem useful, especially for strings. Time will tell what best practices around the use of generics in Go turn out to be.
