package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "sort"
    "strconv"
    "strings"
)

func median(arr []int32) int32 {
    n := len(arr)
    if n % 2 == 1 {
        return arr[n/2]
    }
    return (arr[n/2] + arr[n/2 - 1]) / 2
}

/*
 * Complete the 'quartiles' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func quartiles(arr []int32) []int32 {
    q := make([]int32,3)
    m := len(arr) / 2
    sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
    q[0] = median(arr[:m])
    q[1] = median(arr)
    q[2] = median(arr[m+len(arr)%2:])
    return q
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    dataTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var data []int32

    for i := 0; i < int(n); i++ {
        dataItemTemp, err := strconv.ParseInt(dataTemp[i], 10, 64)
        checkError(err)
        dataItem := int32(dataItemTemp)
        data = append(data, dataItem)
    }

    res := quartiles(data)

    for i, resItem := range res {
        fmt.Fprintf(writer, "%d", resItem)

        if i != len(res) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
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

