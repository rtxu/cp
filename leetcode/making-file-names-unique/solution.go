// Time: O(n)
// Space: O(n)
// ref: https://leetcode.com/problems/making-file-names-unique/discuss/697923/Go-O(n)-Solution-with-explanation
func getFolderNames(names []string) []string {
    n := len(names)
    M := map[string]bool{}
    nextSuffix := map[string]int{}
    result := make([]string, n)
    for i := 0; i < n; i++ {
        origName := names[i]
        nextName := origName
        k := nextSuffix[origName]
        if k == 0 {
            k++
        }
        for M[nextName] {
            nextName = origName + fmt.Sprintf("(%d)", k)
            k++
        }
        result[i] = nextName
        M[nextName] = true
        nextSuffix[origName] = k
    }
    return result
}
