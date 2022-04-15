package main
import (
    "fmt"
    "math"
)

func F(mean, sd, x float64) float64 {
    return 0.5 * (1 + math.Erf((x - mean)/(sd * math.Sqrt(2))))
}

func main() {
    var mean, sd, x1, x2 float64
    fmt.Scanf("%f %f ", &mean, &sd)
    fmt.Scanf("%f ", &x1)
    fmt.Scanf("%f ", &x2)
    fmt.Printf("%.2f\n", 100 * (1 - F(mean, sd, x1)))
    fmt.Printf("%.2f\n", 100 * (1 - F(mean, sd, x2)))
    fmt.Printf("%.2f\n", 100 * F(mean, sd, x2))
}

