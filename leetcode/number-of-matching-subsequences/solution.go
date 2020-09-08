func numMatchingSubseq(S string, words []string) int {
    n := len(S)
    // G(i, ch) 表示从 S[i-1] 位置往后找到的第一个 ch 的位置，0 表示不可达
    // G(0, ch) 表示空串时的状态
    G := make([][26]int, n+1)
    var chPos [26]int
    for i := n; i >= 0; i-- {
        G[i] = chPos
        if i > 0 {
            chPos[S[i-1] - 'a'] = i
        }
    }
    wordCnt := len(words)
    var result int
    for i := 0; i < wordCnt; i++ {
        word := words[i]
        currentPos, matchedLen := 0, 0
        for matchedLen < len(word) && G[currentPos][word[matchedLen]-'a'] > 0 {
            currentPos = G[currentPos][word[matchedLen]-'a']
            matchedLen++
        }
        if matchedLen == len(word) {
            result++
        }
    }
    return result
}
