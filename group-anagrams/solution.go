func getRoot(s string) string {
    C := make([]byte, 26)
    for i := 0; i < len(s); i++ {
        C[s[i] - 'a']++
    }
    return string(C)
}

func groupAnagrams(strs []string) [][]string {
    M := make(map[string]int)
    result := make([][]string, 0)
    n := len(strs)
    for i := 0; i < n; i++ {
        root := getRoot(strs[i])
        var index int
        var exist bool
        if index, exist = M[root]; exist {
        } else {
            index = len(M)
            M[root] = index
            result = append(result, nil)
        }
        result[index] = append(result[index], strs[i])
    }
    
    return result
}
