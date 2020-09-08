func canArrange(arr []int, k int) bool {
    n := len(arr)
    count := map[int]int{}
    for i := 0; i < n; i++ {
        remainder := arr[i] % k
        if remainder < 0 {
            remainder += k
        }
        
        count[remainder]++
    }
    
    if count[0] % 2 != 0 {
        return false
    }
    for i := 1; i < k; i++ {
        if count[i] != count[k-i] {
            return false
        }
    }
    
    return true
}
