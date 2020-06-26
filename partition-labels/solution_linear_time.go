// Time: O(n)
// Space: O(1)

func partitionLabels(S string) []int {
    L, R := [26]int{}, [26]int{}
    for i := 0; i < 26; i++ {
        L[i], R[i] = -1, -1
    }
    n := len(S)
    for i := 0; i < n; i++ {
        char := S[i] - 'a'
        if L[char] == -1 {
            L[char] = i
        }
        R[char] = i
    }
    var result []int
    for i := 0; i < n; i++ {
        char := S[i] - 'a'
        l, r := L[char], R[char]
        for j := l+1; j <= r; j++ {
            char := S[j] - 'a'
            if R[char] > r {
                r = R[char]
            }
        }
        result = append(result, r-l+1)
        i = r
    }
    
    return result
}
