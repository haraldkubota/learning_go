package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

/*
 * Complete the 'maxMin' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY arr
 */

// type foo []int32

// func (f foo) Len() int {
//     return len(f)
// }

// func (f foo) Less(i, j int) bool {
//     return f[i] < f[j]
// }

// func (f foo) Swap(i, j int) {
//     f[i], f[j] = f[j], f[i]
// }


func maxMin(k int32, arr []int32) int32 {
    // with implementing the Sort interface above
    // sort.Sort(arr) 
    // Better:
    sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

    var m int32
    m=100000000
    for i:=0; i < len(arr)-int(k)+1; i+=1 {
        a:=arr[i+int(k)-1]-arr[i]
        if a < m {
            m=a
        }
    }
    return m
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

    kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    k := int32(kTemp)

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := maxMin(k, arr)

    fmt.Fprintf(writer, "%d\n", result)

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
