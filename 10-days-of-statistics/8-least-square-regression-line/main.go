package main
import "fmt"

func fit(x []float64, y []float64) (float64, float64) {
    var sx, sy, sxy, sxx float64
    n := len(x)
    for i := 0; i < n; i++ {
        sx += x[i]
        sy += y[i]
        sxx += x[i] * x[i]
        sxy += x[i] * y[i]
    }
    b := (float64(n) * sxy - sx * sy) / (float64(n) * sxx - sx * sx)
    a := (sy - b * sx) / float64(n)
    return a, b
}

func main() {
    const n = 5;
    x := make([]float64, n)
    y := make([]float64, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%f %f ", &x[i], &y[i])
    }
    a, b := fit(x, y)
    fmt.Printf("%.3f\n", b * 80 + a)
}
