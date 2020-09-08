func hasAllCodes(s string, k int) bool {
    M := make(map[int]bool)
    for i := 0; i < (1<<k); i++ {
        M[i] = true
    }
    
    num := 0
    for i := 0; i < len(s); i++ {
        num <<= 1
        if s[i] == '1' {
            num += 1
        }
        if i >= k && s[i-k] == '1' {
            num -= (1 << k)
        }
        
        //fmt.Println(i, num)
        if i >= k-1 {
            if _, exist := M[num]; exist {
                
                delete(M, num)
            }
        }
    }
    
    return len(M) == 0
}
