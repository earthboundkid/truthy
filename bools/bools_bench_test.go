package bools_test

import (
	"errors"
	"testing"

	"github.com/carlmjohnson/truthy/bools"
)

func BenchmarkAny_error(b *testing.B) {
	fillVal := errors.New("something")
	fill := false
	for i := 0; i < b.N; i++ {
		var value error
		if fill {
			value = fillVal
		}
		if bools.Any(value) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func BenchmarkComparable_error(b *testing.B) {
	fillVal := errors.New("something")
	fill := false
	for i := 0; i < b.N; i++ {
		var value error
		if fill {
			value = fillVal
		}
		if bools.Comparable(value) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func Benchmark_error(b *testing.B) {
	fillVal := errors.New("something")
	fill := false
	for i := 0; i < b.N; i++ {
		var value error
		if fill {
			value = fillVal
		}
		if value != nil != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func BenchmarkAny_string(b *testing.B) {
	fillVal := "something"
	fill := false
	for i := 0; i < b.N; i++ {
		var value string
		if fill {
			value = fillVal
		}
		if bools.Any(value) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func BenchmarkComparable_string(b *testing.B) {
	fillVal := "something"
	fill := false
	for i := 0; i < b.N; i++ {
		var value string
		if fill {
			value = fillVal
		}
		if bools.Comparable(value) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func Benchmark_string(b *testing.B) {
	fillVal := "something"
	fill := false
	for i := 0; i < b.N; i++ {
		var value string
		if fill {
			value = fillVal
		}
		if value != "" != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func BenchmarkAny_int(b *testing.B) {
	fillVal := 1
	fill := false
	for i := 0; i < b.N; i++ {
		var value int
		if fill {
			value = fillVal
		}
		if bools.Any(value) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func BenchmarkComparable_int(b *testing.B) {
	fillVal := 1
	fill := false
	for i := 0; i < b.N; i++ {
		var value int
		if fill {
			value = fillVal
		}
		if bools.Comparable(value) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func Benchmark_int(b *testing.B) {
	fillVal := 1
	fill := false
	for i := 0; i < b.N; i++ {
		var value int
		if fill {
			value = fillVal
		}
		if (value != 0) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}
