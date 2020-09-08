func countTriplets(A []int) int {    
    n := len(A)
    
    D := make([][]int, 4)
    for i := 0; i < 4; i++ {
        D[i] = make([]int, 1<<16)
    }
    D[0][1<<16 - 1] = 1
    for i := 1; i <= 3; i++ {
        for j := 0; j < 1<<16; j++ {
            for k := 0; k < n; k++ {
                D[i][j & A[k]] += D[i-1][j]
            }
        }
    }
    
    return D[3][0]
}
