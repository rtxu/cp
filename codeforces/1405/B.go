package main

import (
    "fmt"
    "bufio"
    "os"
)

const N = 1e5+5

var a [N]int64

func min(a, b int64) int64 {
    if a < b {
        return a 
    } else {
        return b
    }
}

var scanner *bufio.Scanner
func init() {
    scanner = bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1 << 20), 3 << 20)
	scanner.Split(bufio.ScanWords)
}

func read(p interface{}) {
    scanner.Scan()
    fmt.Sscan(scanner.Text(), p)
}

func main() {
    var t int
    for read(&t); t > 0; t-- {
        var n int
        read(&n)
        for i := 0; i < n; i++ {
            read(&a[i])
        }
        // fmt.Println("case #", t, a[:n])
        var result int64
        var more int64
        for i := 0; i < n; i++ {
            if a[i] < 0 {
                delta := min(-a[i], more)
                a[i] += delta
                more -= delta
                result += -a[i]
            } else {
                more += a[i]
                a[i] = 0
            }
        }
        fmt.Println(result)
    }
}
