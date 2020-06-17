func findSubstring(s string, words []string) []int { 
    wordCnt := len(words)
    if wordCnt == 0 {
        return []int{}
    }
    wordLen := len(words[0])
    M := make(map[string]int)
    for i := 0; i < wordCnt; i++ {
        M[words[i]]++
    }
    
    var result []int
    for i := 0; i+wordCnt*wordLen <= len(s); i++ {
        var j int
        for j = 0; j < wordCnt; j++ {
            w := s[i+j*wordLen:i+j*wordLen+wordLen]
            if M[w] > 0 {
                M[w]--
            } else {
                break
            }
        }
        if j == wordCnt {
            result = append(result, i)
        }
        for k := 0; k < j; k++ {
            w := s[i+k*wordLen:i+k*wordLen+wordLen]
            M[w]++
        }
    }
    return result
}
