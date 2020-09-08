func canBeEqual(target []int, arr []int) bool {
    M := make(map[int]int)
    for i := 0; i < len(target); i++ {
        M[target[i]]++
    }
    for i := 0; i < len(arr); i++ {
        M[arr[i]]--
        if M[arr[i]] == 0 {
            delete(M, arr[i])
        }
    }
    
    return len(M) == 0
}
