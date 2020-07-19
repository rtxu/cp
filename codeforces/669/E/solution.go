package main

import "fmt"

const INF = 1e9+17

type BIT map[int]int

func (bit BIT) prefixSum(i int) int {
    sum := 0
    for ; i > 0; i -= i & -i {
        sum += bit[i]
    }
    return sum
}

func (bit BIT) update(i int, delta int) {
    for ; i < INF; i += i & -i {
        bit[i] += delta
    }
}

func main() {
    bits := make(map[int]BIT)

    var n int
    fmt.Scanln(&n)
    for i := 0; i < n; i++ {
        var op, time, oprand int
        fmt.Scanln(&op, &time, &oprand)
        if _, exist := bits[oprand]; !exist {
            bits[oprand] = make(BIT)
        }
        switch op {
        case 1:
            bits[oprand].update(time, 1)
        case 2:
            bits[oprand].update(time, -1)
        case 3:
            fmt.Println(bits[oprand].prefixSum(time))
        }
    }
}
