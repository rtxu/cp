func subarrayBitwiseORs(A []int) int {
    n := len(A)
    current := map[int]bool{}
    all := map[int]bool{}
    for i := 0; i < n; i++ {
        next := map[int]bool{}
        for num := range current {
            next[num | A[i]] = true
            all[num | A[i]] = true
        }
        next[A[i]] = true
        all[A[i]] = true
        
        current = next
    }
    return len(all)
}
