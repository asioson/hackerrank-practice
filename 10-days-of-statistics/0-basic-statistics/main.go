package main
import (
    "fmt"
    "sort"
)

func main() {
    var n, item int
    fmt.Scanf("%d", &n)
    if n > 0 {
        x := make([]int, n)
        for i := 0; i < n; i++ {
            fmt.Scanf("%d", &item)
            x[i] = item
        }
        sort.Ints(x)
        
        var mean float64
        total := 0
        for i := 0; i < n; i++ {
            total += x[i]
        }
        mean = float64(total) / float64(n)
        
        var median float64 
        if n % 2 == 0 {
            median = float64(x[(n-1)/2] + x[((n-1)/2)+1]) / 2.0
        } else {
            median = float64(x[n / 2])
        }
        
        var mode, modeCount, cMode, cModeCount int
        mode, modeCount = x[0], 1
        cMode, cModeCount = x[0], 1
        for i := 1; i < n; i++ {
            if cMode == x[i] {
                cModeCount++
            } else {
                if modeCount < cModeCount {
                    mode = cMode
                    modeCount = cModeCount
                }
                cMode = x[i]
                cModeCount = 1
            }
        }
        if modeCount < cModeCount {
            mode = cMode
            modeCount = cModeCount
        }

        fmt.Printf("%.1f\n", mean)
        fmt.Printf("%.1f\n", median)
        fmt.Printf("%d\n", mode)        
    }    
}

