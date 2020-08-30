// ref: 
// 1. https://blog.csdn.net/Yaokai_AssultMaster/article/details/79492190
// 2. https://cp-algorithms.com/data_structures/fenwick.html

package main

import "fmt"

// Point Update, Range Query
type BITv1 []int

func (bit BITv1) update(i int, delta int) {
    // i++ for 1-based index
    for i++; i < len(bit); i += i & -i {
        bit[i] += delta
    }
}

func (bit BITv1) prefixSum(i int) int {
    var result int
    // i++ for 1-based index
    for i++; i >= 0; i -= i & -i {
        result += bit[i]
    }
    return result
}

func (bit BITv1) rangeSum(l, r int) int {
    return prefixSum(r) - prefixSum(l-1)
}

// Range Update, Point Query

// Range Update, Range Query

func main() {
    for i := 0; i < 10; i++ {
        fmt.Printf("%d & -%d = %d\n", i, i, i & -i)
    }
}
