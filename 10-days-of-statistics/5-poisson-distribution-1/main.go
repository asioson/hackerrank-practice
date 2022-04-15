package main
import (
    "fmt"
    "math"
)

func fact(x int) int {
    if x > 1 {
        return x * fact(x-1)
    }
    return 1
}

func p(x int, mean float64) float64 {
    return math.Pow(mean, float64(x)) * math.Exp(-mean) / float64(fact(x)) 
}

func main() {
    var mean float64
    var x int
    fmt.Scanf("%f ", &mean)
    fmt.Scanf("%d", &x)
    fmt.Printf("%.3f\n", p(x, mean))
}

