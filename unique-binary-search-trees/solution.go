var M = map[int]int{
    0:1,
    1:1,
}

func numTrees(n int) int {
    if v, exist := M[n]; exist {
        return v
    }
    
    sum := 0
    for i := 0; i < n; i++ {
        sum += numTrees(i) * numTrees(n-1-i)
    }
    M[n] = sum
    return sum
}
