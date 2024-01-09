# Truthy [![Go Reference](https://pkg.go.dev/badge/github.com/carlmjohnson/truthy.svg)](https://pkg.go.dev/github.com/carlmjohnson/truthy)

![Truthiness](https://user-images.githubusercontent.com/222245/136619462-f2bc5858-067f-4277-a813-b95c64b3cdac.png)

truthy is a set of packages that use generics to create useful boolean tests of truthiness and also contain some nice helper functions

## Package Examples

### `truthy/condition`
```go
// Ever wish Go had ? : ternary operators?
// Now it has a ternary function.
x := condition.Evaluate(is.Truthy(""), 1, 10) // x == 10

// condition.Evaluate cannot lazily evaluate its arguments,
// but you can use a closure to fake it.
s := condition.Evaluate(is.TruthySlice([]string{""}),
	func() string {
		// do some calculation
		return "foo"
	},
	func() string {
		// do some calculation
		return "bar"
	})()
// s == "foo"
```

### `truthy/defaults`
```go
// How about an equivalent of the nullish coalescing operator ??
// as seen in C#, JavaScript, PHP, etc.:
var s string
defaults.GetFirst(s, "default") // "default"
s = "something"
defaults.GetFirst(s, "default") // "something"
defaults.GetFirst(0, 0*1, 1-1, 0x10-10) // 6

// Easily set defaults
n := getUserInput()
defaults.SetFirst(&n, 42)
```

### `truthy/is`
```go
// is.Truthy returns whether a comparable value
// is equal to the zero value for its type.
//
// is.TruthyAny returns the truthiness of any argument.
// If the value's type has a Bool() bool method, the method is called and returned.
// If the type has an IsZero() bool method, the opposite value is returned.
// Slices and maps are truthy if they have a length greater than zero.
// All other types are truthy if they are not their zero value.

is.Truthy(0) // false
is.Truthy(1) // true

is.Truthy("") // false
is.Truthy(" ") // true

is.TruthySlice([]byte(``)) // false
is.TruthySlice([]byte(` `)) // true

is.TruthySlice([]int{}) // false
is.TruthySlice([]int{1, 2, 3}) // true

var err error
is.Truthy(err) // false
is.Truthy(errors.New("hi")) // true
if is.Truthy(err) {
	panic(err)
}


var p *int
is.Truthy(p) // false

p = new(int)
// Truthy does not check the underlying truthy value for a pointer!
is.Truthy(p) // true
// for that, use TruthyPointer,
is.TruthyPointer(p) // false
// but beware of double pointers!
is.TruthyPointer(&p) // true
```

### `truthy/pointers`
```go
// Start with a null pointer
var strptr *string

// pointers.Deref safely dereferences a nil pointer into its empty value
fmt.Println(pointers.Deref(strptr) == "") // prints true

// pointers.Coalesce lets us specify a default value for a nil pointer
fmt.Println(pointers.Coalesce(strptr, "hello, world")) // prints "hello, world"

// We can create a pointer to a string or other primitive type with pointers.New
newptr := pointers.New("meaning of life") // makes a pointer to a string, wow!

// pointers.First returns the first pointer that isn't nil.
strptr = pointers.First(strptr, newptr) // returns newptr
```
