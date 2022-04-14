package main
import (
    "fmt"
    "math"
)

func comb(n, k int) int {
    a, b := 1, 1
    for i := 1; i <= k; i++ {
        a *= n
        b *= i
        n--
    }
    return a / b
}

func bin(n, x int, p float64) float64 {
    return float64(comb(n, x)) * math.Pow(p, float64(x)) * math.Pow(1-p, float64(n-x))
}

func main() {
    var p float64
    var n int
    fmt.Scanf("%f %d", &p, &n)
    p = p / 100
    prob := float64(0)
    for x := 0; x <= 2; x++ {
        prob += bin(n, x, p)
    }
    fmt.Printf("%.3f\n", prob)
    prob = float64(0)
    for x := 2; x <= n; x++ {
        prob += bin(n, x, p)
    }
    fmt.Printf("%.3f\n", prob)
}
