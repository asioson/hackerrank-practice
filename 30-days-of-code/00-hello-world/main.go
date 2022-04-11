package main
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    inputString, err := reader.ReadString('\n')
    if err == nil {
        fmt.Println("Hello, World.")
        fmt.Print(inputString)
    }
}
