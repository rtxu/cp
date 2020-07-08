// Time: O(n^3), n is the max(row, col)
// Space: O(n)

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func numSubmat(mat [][]int) int {
    row, col := len(mat), len(mat[0])
    h := make([]int, col)
    count := 0
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if mat[i][j] == 1 {
                h[j] += 1
            } else {
                h[j] = 0
            }
        
            for s, currentH := j, h[j]; s >= 0 && currentH > 0; s-- {
                count += currentH
                if s > 0 {
                    currentH = min(currentH, h[s-1])
                }
            }
        }
    }
    
    return count
}
