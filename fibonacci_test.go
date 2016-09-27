package fibonacci

import (
    "math/big"
    "testing"
)

// The tests check that fib(i) == expected[i]
var expected []int64 = []int64{0,1,1,2,3,5,8,13,21,34,55,89,144,233,377,610}

// The benchmarks compute fib(n) for 0 <= n < MAX_N
const MAX_N uint = 30

func TestFib_r(t *testing.T) {
    for i, v := range(expected) {
        got := Fib_r(uint(i))
        if got != uint(v) {
            t.Errorf("Fib_r(%d) != %d", i, v)
        }
    }
}

func TestFib_i(t *testing.T) {
    for i, v := range(expected) {
        got := Fib_i(uint(i))
        if got != uint(v) {
            t.Errorf("Fib_i(%d) != %d", i, v)
        }
    }
}

func TestFibonacci_r(t *testing.T) {
    for i, v := range(expected) {
        got := Fibonacci_r(uint(i))
        if got.Cmp(big.NewInt(v)) != 0 {
            t.Errorf("Fibonacci_r(%d) != %d", i, v)
        }
    }
}

func TestFibonacci_i(t *testing.T) {
    for i, v := range(expected) {
        got := Fibonacci_i(uint(i))
        if got.Cmp(big.NewInt(v)) != 0 {
            t.Errorf("Fibonacci_i(%d) != %d", i, v)
        }
    }
}

func TestFibonacci_c(t *testing.T) {
    for i, v := range(expected) {
        got := Fibonacci_c(uint(i))
        if got.Cmp(big.NewInt(v)) != 0 {
            t.Errorf("Fibonacci_c(%d) != %d", i, v)
        }
    }
}

func BenchmarkFib_r(b *testing.B) {
    for i := 0; i < b.N; i++ {
        for n := uint(0); n < MAX_N; n++ {
            Fib_r(n)
        }
    }
}

func BenchmarkFib_i(b *testing.B) {
    for i := 0; i < b.N; i++ {
        for n := uint(0); n < MAX_N; n++ {
            Fib_i(n)
        }
    }
}

func BenchmarkFibonacci_r(b *testing.B) {
    for i := 0; i < b.N; i++ {
        for n := uint(0); n < MAX_N; n++ {
            Fibonacci_r(n)
        }
    }
}

func BenchmarkFibonacci_i(b *testing.B) {
    for i := 0; i < b.N; i++ {
        for n := uint(0); n < MAX_N; n++ {
            Fibonacci_i(n)
        }
    }
}

func BenchmarkFibonacci_c(b *testing.B) {
    for i := 0; i < b.N; i++ {
        for n := uint(0); n < MAX_N; n++ {
            Fibonacci_c(n)
        }
    }
}
