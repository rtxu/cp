func reorganizeString(S string) string {
    count := map[byte]int{}
    n := len(S)
    var maxCount int
    var maxCountChar byte
    for i := 0; i < n; i++ {
        count[S[i]]++
        if count[S[i]] > maxCount {
            maxCount = count[S[i]]
            maxCountChar = S[i]
        }
    }
    if maxCount > (n+1)/2 {
        return ""
    } else {
        result := make([]byte, n)
        i := 0
        for j := 0; j < maxCount; j++ {
            result[i] = maxCountChar
            i += 2
            if i >= n {
                i = 1
            }
        }
        delete(count, maxCountChar)
        for ch, cnt := range count {
            for j := 0; j < cnt; j++ {
                result[i] = ch
                i += 2
                if i >= n {
                    i = 1
                }
            }
        }
        return string(result)
    }
}
