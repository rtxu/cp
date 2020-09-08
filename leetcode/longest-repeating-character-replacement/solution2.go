// Time: O(n)
// Space: O(26)
func characterReplacement(s string, k int) int {
    chCount := make([]int, 26)
    
    var max, maxWindowSize int
    for i, j := 0, 0; j < len(s); j++ {
        ch := int(s[j]-'A')
        chCount[ch]++
        if chCount[ch] > max {
            max = chCount[ch]
        }
        for j-i+1-max > k {
            chCount[int(s[i]-'A')]--
            i++
        }
        if j-i+1 > maxWindowSize {
            maxWindowSize = j-i+1
        }
    }
    return maxWindowSize
}
