func fourSumCount(A []int, B []int, C []int, D []int) int {
    sumAB := map[int]int{}
    n := len(A)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            sumAB[A[i]+B[j]]++
        }
    }
    var result int
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            result += sumAB[-C[i]-D[j]]
        }
    }
    return result
}
