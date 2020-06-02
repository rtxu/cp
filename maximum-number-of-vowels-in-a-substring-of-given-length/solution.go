func maxVowels(s string, k int) int {
    isVowel := func(ch byte) bool {
        return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
    }
    
    n := len(s)
    vowelCount := 0
    var maxVowelCount int
    for i := 0; i < n; i++ {
        if isVowel(s[i]) {
            vowelCount++
        }
        if i >= k && isVowel(s[i-k]) {
            vowelCount--
        }
        if vowelCount > maxVowelCount {
            maxVowelCount = vowelCount
        }
    }
    return maxVowelCount
}
