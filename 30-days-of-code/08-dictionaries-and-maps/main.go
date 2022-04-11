package main
import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func main() {
    var n int
    phone := map[string]string{}
    fmt.Scanf("%d ", &n)
    scanner := bufio.NewScanner(os.Stdin)
    for i := 0; i < n; i++ {
        scanner.Scan()
        data := strings.Split(scanner.Text(), " ")
        phone[data[0]] = data[1]
    }
    for scanner.Scan() {
        name := scanner.Text()
        if number, ok := phone[name]; ok {
            fmt.Printf("%s=%s\n", name, number)
        } else {
            fmt.Println("Not found")
        }
    }
}

