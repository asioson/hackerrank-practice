package main
import (
    "fmt"
    "math"
)

func F(mean, sd, x float64) float64 {
    return 0.5 * (1 + math.Erf((x - mean) / (sd * math.Sqrt(2))))
}

func main() {
    var maxCapacity, n, mean, sd float64
    fmt.Scanf("%f ", &maxCapacity) 
    fmt.Scanf("%f ", &n) 
    fmt.Scanf("%f ", &mean) 
    fmt.Scanf("%f ", &sd)
    x := maxCapacity / n
    sd = sd / math.Sqrt(n)
    fmt.Printf("%.4f\n", F(mean, sd, x))
}

