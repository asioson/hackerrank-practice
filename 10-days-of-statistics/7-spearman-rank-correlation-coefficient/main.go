package main
import (
    "fmt"
    "sort"
)

func getRank(x []float64) map[float64]int {
    rank := map[float64]int{}
    n := len(x)
    tx := make([]float64, n)
    for i := 0; i < n; i++ {
        tx[i] = x[i]
     }
    sort.Float64s(tx)
    for i := 0; i < n; i++ {
        rank[tx[i]] = i+1
    }
    return rank
}

func spearman(x, y []float64) float64 {
    n := len(x)
    xrank := getRank(x)
    yrank := getRank(y)
    ddSum := 0
    for i := 0; i < n; i++ {
        d := xrank[x[i]] - yrank[y[i]]
        ddSum += d * d
    }
    return 1 - (6 * float64(ddSum) / float64(n * n * n - n))
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
    fmt.Printf("%.3f\n", spearman(x, y))
}
