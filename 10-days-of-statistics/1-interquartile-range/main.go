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

type Data struct {
    v, f, idx int32
}

func getValue(arr []Data, k int32) int32 {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        if arr[i].idx <= k && k < arr[i+1].idx {
            return arr[i].v
        }
    }
    return arr[n-1].v
}

/*
 * Complete the 'interQuartile' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY values
 *  2. INTEGER_ARRAY freqs
 */

func interQuartile(values []int32, freqs []int32) {
    n := len(values)
    data := make([]Data, n)
    for i := 0; i < n; i++ {
        data[i].v = values[i]
        data[i].f = freqs[i]
    }
    total := data[0].f
    sort.Slice(data, func (i int, j int) bool { return data[i].v < data[j].v })
    for i := 1; i < n; i++ {
        data[i].idx = data[i-1].idx + data[i-1].f
        total += data[i].f
    }
    q1, q3 := float64(0), float64(0)
    m := total / 2
    r := total % 2
    if m % 2 == 0 {
        q1 = float64(getValue(data, m / 2 - 1) + getValue(data, m / 2)) / 2.0
        q3 = float64(getValue(data, m + r + m / 2 - 1) + getValue(data, m + r + m / 2)) / 2.0
    } else {
        q1 = float64(getValue(data, m / 2))
        q3 = float64(getValue(data, m + r + m / 2))
    }
    fmt.Printf("%.1f\n", q3 - q1)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    valTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var val []int32

    for i := 0; i < int(n); i++ {
        valItemTemp, err := strconv.ParseInt(valTemp[i], 10, 64)
        checkError(err)
        valItem := int32(valItemTemp)
        val = append(val, valItem)
    }

    freqTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var freq []int32

    for i := 0; i < int(n); i++ {
        freqItemTemp, err := strconv.ParseInt(freqTemp[i], 10, 64)
        checkError(err)
        freqItem := int32(freqItemTemp)
        freq = append(freq, freqItem)
    }

    interQuartile(val, freq)
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

