func lowbit(i int) int {
    i32 := uint32(i)
    return int(i32 & (i32 ^ (1<<32-1) + 1))
}

func countBits(num int) []int {
    result := make([]int, num+1)
    result[0] = 0
    for i := 1; i <= num; i++ {
        fmt.Println(num, i, lowbit(i))
        result[i] = result[i-lowbit(i)] + 1
    }
    return result
}
