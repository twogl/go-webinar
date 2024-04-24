package my_math_test

import (
	"fmt"
	"testing"

	"github.com/twogl/go-webinar/godoc_and_tests/my_math"
)

func Test_Add(t *testing.T) {
	testCases := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"1 + 2", 1, 2, 3},
		{"4 + (-1)", 4, -1, 3},
		{"0 + 0", 0, 0, 0},
		{"4 + (-7)", 4, -7, -3},
		{"(-3) + (-7)", -3, -7, -10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := my_math.Add(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("expected: %d, got: %d", tc.want, got)
			}
		})
	}
}

func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		my_math.Add(1, 2)
	}
}

func ExampleAdd() {
	fmt.Println(my_math.Add(5, 7))
	// Output: 12
}
func ExampleMultiply() {
	fmt.Println(my_math.Multiply(5, 7))
	// Output: 35
}
