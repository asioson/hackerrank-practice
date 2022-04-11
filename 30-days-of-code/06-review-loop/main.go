package main
import "fmt"

func main() {
    var n int
    fmt.Scanf("%d", &n)
    for i := 0; i < n; i++ {
        var s, os, es string
        fmt.Scanf("%s", &s)
        m := len(s)
        for j := 0; j < m - 1; j += 2 {
            os += string(s[j])
            es += string(s[j+1])
        }
        if m % 2 == 1 {
            os += string(s[m-1])
        }
        fmt.Println(os, es)
    }
}
