package is_test

import (
	"errors"
	"testing"

	"github.com/carlmjohnson/truthy/is"
)

func BenchmarkTruthyAny_error(b *testing.B) {
	fillVal := errors.New("something")
	fill := false
	for i := 0; i < b.N; i++ {
		var value error
		if fill {
			value = fillVal
		}
		if is.TruthyAny(value) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func BenchmarkTruthy_error(b *testing.B) {
	fillVal := errors.New("something")
	fill := false
	for i := 0; i < b.N; i++ {
		var value error
		if fill {
			value = fillVal
		}
		if is.Truthy(value) != fill {
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

func BenchmarkTruthyAny_string(b *testing.B) {
	fillVal := "something"
	fill := false
	for i := 0; i < b.N; i++ {
		var value string
		if fill {
			value = fillVal
		}
		if is.TruthyAny(value) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func BenchmarkTruthy_string(b *testing.B) {
	fillVal := "something"
	fill := false
	for i := 0; i < b.N; i++ {
		var value string
		if fill {
			value = fillVal
		}
		if is.Truthy(value) != fill {
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

func BenchmarkTruthyAny_int(b *testing.B) {
	fillVal := 1
	fill := false
	for i := 0; i < b.N; i++ {
		var value int
		if fill {
			value = fillVal
		}
		if is.TruthyAny(value) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func BenchmarkTruthy_int(b *testing.B) {
	fillVal := 1
	fill := false
	for i := 0; i < b.N; i++ {
		var value int
		if fill {
			value = fillVal
		}
		if is.Truthy(value) != fill {
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
