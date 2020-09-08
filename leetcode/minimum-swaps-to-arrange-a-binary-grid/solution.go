var sum [201][201]int

func rangeSum(row, start, end int) int {
    return sum[row][end] - sum[row][start]
}

func minSwaps(grid [][]int) int {
    n := len(grid)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            sum[i][j+1] = sum[i][j] + grid[i][j]
        }
    }
    
    steps := 0
    for i := 0; i < n; i++ {
        var j int
        for j = i; j < n && rangeSum(j, i+1, n) != 0; j++ {}
        if j >= n {
            return -1
        } else {
            steps += j-i
            temp := sum[j]
            for k := j-1; k >= i; k-- {
                sum[k+1] = sum[k]
            }
            sum[i] = temp
        }
    }
    return steps
}
