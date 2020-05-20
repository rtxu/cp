func countTriplets(A []int) int {    
    n := len(A)
    M := make(map[int]int)
    count := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            M[A[i] & A[j]]++
        }
    }
    for k := 0; k < n; k++ {
        for a, cnt := range M {
            if A[k] & a == 0 {
                count += cnt
            }
        }
    }
    return count
}
