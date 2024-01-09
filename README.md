# Truthy [![Go Reference](https://pkg.go.dev/badge/github.com/carlmjohnson/truthy.svg)](https://pkg.go.dev/github.com/carlmjohnson/truthy)

![Truthiness](https://user-images.githubusercontent.com/222245/136619462-f2bc5858-067f-4277-a813-b95c64b3cdac.png)

`truthy` leverages generics to create a useful collection of boolean tests of truthiness and helper functions

## Package Examples

### `truthy/bools`
```go
// bools.Comparable returns whether a comparable value
// is equal to the zero value for its type.
//
// bools.Any returns the truthiness of any argument.
// If the value's type has a Bool() bool method, the method is called and returned.
// If the type has an IsZero() bool method, the opposite value is returned.
// Slices and maps are truthy if they have a length greater than zero.
// All other types are truthy if they are not their zero value.

bools.Comparable(0) // false
bools.Comparable(1) // true

bools.Comparable("") // false
bools.Comparable(" ") // true

bools.Map(map[string]string{}) // false
bools.Map(map[string]string{"truthy": "is awesome"}) // true

bools.Slice([]int{}) // false
bools.Slice([]int{1, 2, 3}) // true

var err error
bools.Comparable(err) // false
bools.Comparable(errors.New("hi")) // true
if bools.Comparable(err) {
	panic(err)
}


var p *int
bools.Comparable(p) // false

p = new(int)
// Comparable does not check the underlying truthy value for a pointer!
bools.Comparable(p) // true
// for that, use Pointer,
bools.Pointer(p) // false
// but beware of double pointers!
bools.Pointer(&p) // true

// How about an equivalent of the nullish coalescing operator ??
// as seen in C#, JavaScript, PHP, etc.:
var s string
bools.GetFirst(s, "default") // "default"
s = "something"
bools.GetFirst(s, "default") // "something"
bools.GetFirst(0, 0*1, 1-1, 0x10-10) // 6

// Easily set defaults
n := getUserInput()
bools.SetFirst(&n, 42)

// Want to be able to convert from Boolean to Int?
bools.Int(true) // 1
bools.Int(10 > 100) // 0
bools.Int(bools.Comparable("")) // 0

// Ever wish Go had ? : ternary operators?
// Now it has a ternary function.
x := bools.Ternary(bools.Comparable(""), 1, 10) // x == 10

// bools.Ternary cannot lazily evaluate its arguments,
// but you can use a closure to fake it.
s := bools.Ternary(bools.Slice([]string{""}),
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
