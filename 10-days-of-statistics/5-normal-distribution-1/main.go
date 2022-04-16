package main
import (
    "fmt"
    "math"
)

func F(mean, sd, x float64) float64 {
    return 0.5 * (1 + math.Erf((x - mean) / (sd * math.Sqrt(2))))
}

func main() {
    var mean, sd, x, x1, x2 float64
    fmt.Scanf("%f %f ", &mean, &sd)
    fmt.Scanf("%f ", &x)
    fmt.Scanf("%f %f ", &x1, &x2)
    fmt.Printf("%.3f\n", F(mean, sd, x))
    fmt.Printf("%.3f\n", F(mean, sd, x2) - F(mean, sd, x1))
}
