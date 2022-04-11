package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)
    
    maxCount, count := 0, 0
    for n > 0 {
        if n % 2 == 1 {
            count++
        } else {
            if maxCount < count {
                maxCount = count
            }
            count = 0
        }
        n = n / 2
    }
    if maxCount < count {
        maxCount = count
    }
    fmt.Println(maxCount)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

