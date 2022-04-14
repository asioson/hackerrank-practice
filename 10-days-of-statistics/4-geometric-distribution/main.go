package main
import "fmt"

func g(n int, p float64) float64 {
    prob := p
    for i := 0; i < n-1; i++ {
        prob *= (1 - p)
    } 
    return prob
}

func main() {
    var a, b, n int
    fmt.Scanf("%d %d\n%d", &a, &b, &n)
    p := float64(a) / float64(b)
    fmt.Printf("%.3f\n", g(n, p))
}
