package main
import (
    "fmt"
    "math"
)

func mean(x []float64) float64 {
    total := float64(0)
    for i := 0; i < len(x); i++ {
        total += x[i]
    }
    return total / float64(len(x))
}

func sd(x []float64) float64 {
    m := mean(x)
    s := float64(0)
    for i := 0; i < len(x); i++ {
        s += (x[i] - m) * (x[i] - m)
    }
    return math.Sqrt(s / float64(len(x)))
}

func pearson(x, y []float64) float64 {
    mx, my := mean(x), mean(y)
    sdx, sdy := sd(x), sd(y)
    n := len(x)
    total := float64(0)
    for i := 0; i < n; i++ {
        total += (x[i] - mx) * (y[i] - my)
    }
    return total / (float64(n) * sdx * sdy)
}

func readValues(n int) []float64 {
    x := make([]float64, n)
    for i := 0; i < n-1; i++ {
        fmt.Scanf("%f", &x[i])
    }
    fmt.Scanf("%f ", &x[n-1])
    return x    
}

func main() {
    var n int
    fmt.Scanf("%d ", &n)
    x := readValues(n)
    y := readValues(n)
    fmt.Printf("%.3f\n", pearson(x,y))
}
