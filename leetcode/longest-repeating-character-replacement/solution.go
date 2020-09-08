// Time: O(26n)
// Space: O(26)
func characterReplacement(s string, k int) int {
    chCount := make(map[byte]int)
    isValid := func() bool {
        sum, max := 0, 0
        for _, cnt := range chCount {
            sum += cnt
            if cnt > max {
                max = cnt
            }
        }
        return (sum-max) <= k
    }
    
    var maxWindowSize int
    for i, j := 0, 0; j < len(s); j++ {
        chCount[s[j]]++
        for !isValid() {
            chCount[s[i]]--
            i++
        }
        if j-i+1 > maxWindowSize {
            maxWindowSize = j-i+1
        }
    }
    return maxWindowSize
}
