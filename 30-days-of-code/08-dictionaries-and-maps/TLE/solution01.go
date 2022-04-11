package main
import "fmt"

func main() {
    var n int
    var name, number string
    phone := map[string]string{}
    fmt.Scanf("%d ", &n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%s %s ", &name, &number)
        phone[name] = number
    }
    n, _ = fmt.Scanf("%s ", &name)
    for n == 1 {
        if number, ok := phone[name]; ok {
            fmt.Printf("%s=%s\n", name, number)
        } else {
            fmt.Println("Not found")
        }
        n, _ = fmt.Scanf("%s ", &name)
    }
}

