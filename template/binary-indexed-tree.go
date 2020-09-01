// ref: 
// 1. https://blog.csdn.net/Yaokai_AssultMaster/article/details/79492190
// 2. https://cp-algorithms.com/data_structures/fenwick.html

package main

import "fmt"

// Point Update, Range Query
type BITv1 []int

func (bit BITv1) Update(i int, delta int) {
    // i++ for 1-based index
    for i++; i < len(bit); i += i & -i {
        bit[i] += delta
    }
}

func (bit BITv1) PrefixSum(i int) int {
    var result int
    // i++ for 1-based index
    for i++; i > 0; i -= i & -i {
        result += bit[i]
    }
    return result
}

// [l,r] both inclusive
func (bit BITv1) RangeSum(l, r int) int {
    return bit.PrefixSum(r) - bit.PrefixSum(l-1)
}

// Range Update, Point Query: 基于差分思想
type BITv2 []int

func (bit BITv2) update(i int, delta int) {
    // i++ for 1-based index
    for i++; i < len(bit); i += i & -i {
        bit[i] += delta
    }
}

// [l,r] both inclusive
func (bit BITv2) RangeUpdate(l, r int, delta int) {
    bit.update(l, delta)
    bit.update(r+1, -delta)
}


func (bit BITv2) PointUpdate(i int, delta int) {
    bit.RangeUpdate(i, i, delta)
}

func (bit BITv2) PointQuery(i int) int {
    var result int
    // i++ for 1-based index
    for i++; i > 0; i -= i & -i {
        result += bit[i]
    }
    return result
}

// Range Update, Range Query
// ref: https://petr-mitrichev.blogspot.com/2013/05/fenwick-tree-range-updates.html
type BITv3 struct {
    mul []int
    add []int
}

func (bit BITv3) update(i int, mul int, add int) { 
    for i++; i < len(bit.mul); i += i&-i {
        bit.mul[i] += mul
        bit.add[i] += add
    }
}

// [l,r] both inclusive
func (bit BITv3) RangeUpdate(l, r int, delta int) { 
    bit.update(l, delta, -delta * (l-1))
    bit.update(r+1, -delta, delta * r)
}

func (bit BITv3) PrefixSum(i int) int { 
    var mul, add int
    start := i
    for i++; i > 0; i -= i&-i {
        mul += bit.mul[i]
        add += bit.add[i]
    }
    return mul * start + add
}

// [l,r] both inclusive
func (bit BITv3) RangeSum(l, r int) int { 
    return bit.PrefixSum(r) - bit.PrefixSum(l-1)
}

func main() {
    const N = 10
    printSeperator := func(end bool) {
        if end {
            fmt.Println()
        } else {
            fmt.Print(" ")
        }
    }
    printArrayIndex := func(n int) {
        for i := 0; i < n; i++ {
            fmt.Printf("%2d", i)
            printSeperator(i == n-1)
        }
    }

    fmt.Println("Test BITv1: Point Update, Range Query")
    bit1 := make(BITv1, N+1)
    fmt.Println("Update sequence")
    printArrayIndex(N)
    for i := 0; i < N; i++ {
        bit1.Update(i, i+1)
        fmt.Printf("%2d", i+1)
        printSeperator(i == N-1)
    }
    fmt.Printf("PrefixSum(%d) = %d\n", N-1, bit1.PrefixSum(N-1))
    fmt.Printf("RangeSum(%d, %d) = %d\n", 4, N-1, bit1.RangeSum(4, N-1))

    fmt.Println()
    fmt.Println("Test BITv2: Range Update, Point Query")
    bit2 := make(BITv2, N+1)
    bit2.RangeUpdate(N/2, N-1, 2)
    fmt.Printf("RangeUpdate(l = %d, r = %d, delta = %d)\n", N/2, N-1, 2)
    printArrayIndex(N)
    for i := 0; i < N; i++ {
        fmt.Printf("%2d", bit2.PointQuery(i))
        printSeperator(i == N-1)
    }

    fmt.Println()
    fmt.Println("Test BITv3: Range Update, Range Query")
    bit3 := BITv3{
        mul: make([]int, N+1),
        add: make([]int, N+1),
    }
    bit3.RangeUpdate(N/2, N-1, 2)
    fmt.Printf("RangeUpdate(l = %d, r = %d, delta = %d)\n", N/2, N-1, 2)
    fmt.Printf("RangeSum(l = %d, r = %d) = %d\n", 1, N/2-1, bit3.RangeSum(1, N/2-1))
    fmt.Printf("RangeSum(l = %d, r = %d) = %d\n", 1, N/2, bit3.RangeSum(1, N/2))
    fmt.Printf("RangeSum(l = %d, r = %d) = %d\n", N/2, N-1, bit3.RangeSum(N/2, N-1))
    bit3.RangeUpdate(4, 7, 3)
    fmt.Printf("RangeUpdate(l = %d, r = %d, delta = %d)\n", 4, 7, 3)
    fmt.Printf("RangeSum(l = %d, r = %d) = %d\n", 1, N/2-1, bit3.RangeSum(1, N/2-1))
    fmt.Printf("RangeSum(l = %d, r = %d) = %d\n", 1, N/2, bit3.RangeSum(1, N/2))
    fmt.Printf("RangeSum(l = %d, r = %d) = %d\n", N/2, N-1, bit3.RangeSum(N/2, N-1))
}
