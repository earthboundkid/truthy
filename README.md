# Truthy

![Truthiness](https://user-images.githubusercontent.com/222245/136619462-f2bc5858-067f-4277-a813-b95c64b3cdac.png)

Truthy is a package which uses generics (Go 1.18+) to create useful boolean tests and helper functions.

## Examples

```
	var err error
	truthy.Value(err) // false

	err = errors.New("hi")
	truthy.Value(err) // true

	if truthy.Value(err) {
		panic(err)
	}

	var n int
	truthy.Value(n) // false

	n = 1
	truthy.Value(n) // true

	var p *int
	truthy.Value(p) // false

	p = new(int)
	// truthy does not check value underlying pointer!
	truthy.Value(p) // true

	if truthy.Or("1", 0) {
		fmt.Println("yay") // prints yay
	}

	if truthy.And(300, "") {
		fmt.Println("boo") // not executed
	}

	// Ever wish Go has ? : ternary operators?
	x := truthy.Cond("", 1, 10) // x == 10
	// truthy.Cond cannot lazily evaluate its arguments,
	// but you can use a closure to fake it.
	s := truthy.Cond(x,
		func() string {
			// do some calculation
			return "foo"
		},
		func() string {
			// do some calculation
			return "bar"
		})()
	// s == "bar"

	// How about an equivalent of ?? in C#, JavaScript, PHP, etc.:
	var s string
	truthy.First(s, "default") // "default"
	s = "something"
	truthy.First(s, "default") // "something"

	// Easily set defaults
	n := getUserInput()
	truthy.SetDefault(&n, 42)
```

## Discussion

Should you use this package? It's a little bit of a joke package, but the set default functionality is useful, especially for strings. Time will tell what best practices around the use of generics in Go turn out to be.
