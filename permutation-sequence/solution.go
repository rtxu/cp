func backtrack(current []byte, visited []bool, n int, k *int) string {
    if len(current) == n {
        (*k)--
        if (*k) == 0 {
            return string(current)
        }
        return ""
    }
    
    for i := 1; i <= n; i++ {
        if !visited[i] {
            current = append(current, byte('0'+i))
            visited[i] = true
            result := backtrack(current, visited, n, k)
            if result != "" {
                return result
            }
            visited[i] = false
            current = current[0:len(current)-1]
        }
    }
    return ""
}

func getPermutation(n int, k int) string {
    visited := make([]bool, n+1)
    return backtrack([]byte{}, visited, n, &k)
}
