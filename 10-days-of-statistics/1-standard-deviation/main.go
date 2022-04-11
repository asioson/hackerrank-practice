package main

import (
    "bufio"
    "fmt"
    "io"
    "math"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'stdDev' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func stdDev(arr []int32) {
    n := len(arr)
    if n > 0 {
        total, mean := float64(0), float64(0)
        for i := 0; i < n; i++ {
            total += float64(arr[i])
        }
        mean = total / float64(n)
        total = 0
        for i := 0; i < n; i++ {
            diff := float64(arr[i]) - mean 
            total += diff * diff       
        }
        sd := math.Sqrt(total / float64(n))
        fmt.Printf("%.1f\n", sd)
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    valsTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var vals []int32

    for i := 0; i < int(n); i++ {
        valsItemTemp, err := strconv.ParseInt(valsTemp[i], 10, 64)
        checkError(err)
        valsItem := int32(valsItemTemp)
        vals = append(vals, valsItem)
    }

    stdDev(vals)
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

