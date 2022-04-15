package main
import "fmt"

func main() {
    var a, b float64
    fmt.Scanf("%f %f", &a, &b)
    fmt.Printf("%.3f\n", 160 + 40 * (a + a * a))
    fmt.Printf("%.3f\n", 128 + 40 * (b + b * b))
}
