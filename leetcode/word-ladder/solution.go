func ladderLength(beginWord string, endWord string, wordList []string) int {
    n := len(wordList)
    newWordList := make([]string, n+1)
    newWordList[0] = beginWord
    copy(newWordList[1:], wordList)
    wordList = newWordList
    n++
    endWordInList := false
    endWordIndex := -1
    for i := 0; i < n; i++ {
        if wordList[i] == endWord {
            endWordInList = true
            endWordIndex = i
            break
        }
    }
    if !endWordInList {
        return 0
    }
    
    wordLen := len(beginWord)
    connected := func(i, j int) bool {
        diffLen := 0
        for k := 0; k < wordLen && diffLen <= 1; k++ {
            if wordList[i][k] != wordList[j][k] {
                diffLen++
            }
        }
        return diffLen <= 1
    }
    
    G := make([][]int, n)
    for i := 0; i < n; i++ {
        G[i] = make([]int, n)   
    }
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if j != i && connected(i, j) {
                G[i][0]++
                G[j][0]++
                G[i][G[i][0]], G[j][G[j][0]] = j, i
            }
        }
    }
    
    const UNVISITED = 0
    minCost := make([]int, n)
    for i := 1; i < n; i++ {
        minCost[i] = UNVISITED
    }
    minCost[0] = 1
    
    Q := make([]int, 1)
    Q[0] = 0
    for len(Q) > 0 {
        current := Q[0]
        Q = Q[1:]
        for i := 1; i <= G[current][0]; i++ {
            j := G[current][i]
            if minCost[j] == UNVISITED {
                minCost[j] = minCost[current]+1
                Q = append(Q, j)
            }
        }
    }
    
    return minCost[endWordIndex]
}
