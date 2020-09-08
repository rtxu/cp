const MOD = 1e9 + 7

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
    sort.Ints(horizontalCuts)
    sort.Ints(verticalCuts)
    
    heights := make([]int, len(horizontalCuts) + 1)
    last := 0
    for i := 0; i < len(horizontalCuts); i++ {
        current := horizontalCuts[i]
        heights[i] = current - last
        last = current
    }
    heights[len(horizontalCuts)] = h - last
    //fmt.Println(heights)
    
    last = 0
    widths := make([]int, len(verticalCuts) + 1)
    for i := 0; i < len(verticalCuts); i++ {
        current := verticalCuts[i]
        widths[i] = current - last
        last = current
    }
    widths[len(verticalCuts)] = w - last
    //fmt.Println(widths)
    
    sort.Ints(heights)
    sort.Ints(widths)
    /*
    fmt.Println(heights)
    fmt.Println(widths)
    
    fmt.Println(heights[len(horizontalCuts)])
    fmt.Println(widths[len(verticalCuts)])
    fmt.Println(int(int64(heights[len(horizontalCuts)]) * int64(widths[len(verticalCuts)]) % MOD))
    */
    return int(int64(heights[len(horizontalCuts)]) * int64(widths[len(verticalCuts)]) % MOD)
}
