package fibonacci

import (
    "math/big"
    "sync"
)

// Recursive implementation without big int support
func Fib_r(n uint) uint {
    if n < 2 {
        return n
    }
    return Fib_r(n-2) + Fib_r(n-1)
}

// Iterative implementation without big int support
func Fib_i(n uint) uint {
    var a, b uint = 0, 1
    for i := uint(0); i < n; i++ {
        a, b = b, a + b
    }
    return a
}

// Recursive implementation with big int support
func Fibonacci_r(n uint) *big.Int {
    if n == 0 {
        return big.NewInt(0)
    }
    if n == 1 {
        return big.NewInt(1)
    }
    x := Fibonacci_r(n-1)
    y := Fibonacci_r(n-2)
    return x.Add(x, y)
}

// Iterative implementation with big int support
func Fibonacci_i(n uint) *big.Int {
    a := big.NewInt(0)
    b := big.NewInt(1)
    for i := uint(0); i < n; i++ {
        a.Add(a,b)
        a, b = b, a
    }
    return a
}

// Cacheing implementation with big int support
var cache []*big.Int = []*big.Int{big.NewInt(0), big.NewInt(1)}
var mu sync.Mutex

func Fibonacci_c(n uint) *big.Int {
    mu.Lock()
    for i := uint(len(cache)); i <= n; i++ {
        var t big.Int
        t.Add(cache[i-1], cache[i-2])
        cache = append(cache, &t)
    }
    mu.Unlock()
    return cache[n]
}
