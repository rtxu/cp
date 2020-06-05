func spiralOrder(matrix [][]int) []int {
    row := len(matrix)
    if row == 0 {
        return []int{}
    }
    col := len(matrix[0])
    
    var result []int
    var start int
    for start = 0; start < row-1 && start < col-1; start++ {
        // (start, start) -> (start, col-1)
        //       |                 |
        // (row-1, start) <- (row-1, col-1)
        i, j := start, start
        for ; j < col-1; j++ {
            result = append(result, matrix[i][j])
        }
        // j = col-1
        for ; i < row-1; i++ {
            result = append(result, matrix[i][j])
        }
        // i = row-1
        for ; j > start; j-- {
            result = append(result, matrix[i][j])
        }
        // j = start
        for ; i > start; i-- {
            result = append(result, matrix[i][j])
        }
        row--
        col--
    }
    if start == row-1 {
        i, j := start, start
        for ; j <= col-1; j++ {
            result = append(result, matrix[i][j])
        }
    } else if start == col-1 {
        i, j := start, start
        for ; i <= row-1; i++ {
            result = append(result, matrix[i][j])
        }
    }

    return result
}
