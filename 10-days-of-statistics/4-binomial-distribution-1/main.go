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
    return float64(comb(n, x)) * math.Pow(p, float64(x)) * math.Pow(1 - p, float64(n-x))
}

func main() {
    var b, g float64
    fmt.Scanf("%f %f", &b, &g)
    
    n, p, answer := 6, b / (b + g), float64(0)
    for x := 3; x <= n; x++ {
        answer += bin(n, x, p)
    }
    fmt.Printf("%.3f\n", answer)
}
