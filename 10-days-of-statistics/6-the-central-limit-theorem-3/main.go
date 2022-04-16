package main
import (
    "fmt"
    "math"
)

func main() {
    var n, mean, sd, p, z float64
    fmt.Scanf("%f ", &n)
    fmt.Scanf("%f ", &mean)
    fmt.Scanf("%f ", &sd)
    fmt.Scanf("%f ", &p)
    fmt.Scanf("%f ", &z)
    a := -z * sd / math.Sqrt(n) + mean
    b := z * sd / math.Sqrt(n) + mean
    fmt.Printf("%.2f\n", a)
    fmt.Printf("%.2f\n", b)
}
