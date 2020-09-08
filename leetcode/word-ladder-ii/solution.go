type param struct {
    current int
    target int
    // 当前已走路径
    path []string
    pathLen int
    // 对应当前路径的 visited
    visited []bool
    G [][]int
    minCost []int
    wordList []string
    result *[][]string
}

func dfs(p param) {
    p.visited[p.current] = true
    p.path[p.pathLen] = p.wordList[p.current]
    p.pathLen++
    
    if p.current == p.target {
        newResult := make([]string, p.pathLen)
        copy(newResult, p.path[:p.pathLen])
        (*p.result) = append((*p.result), newResult)
    }
    
    for k := 1; k <= p.G[p.current][0]; k++ {
        next := p.G[p.current][k]
        if !p.visited[next] && p.pathLen+1 <= p.minCost[next] {
            nextP := p
            nextP.current = next
            dfs(nextP)
        }
    }
    
    p.pathLen--
    p.visited[p.current] = false
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
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
        return [][]string{}
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
    
    Q := make([]int, 1)
    Q[0] = 0
    minCost := make([]int, n)
    minCost[0] = 1
    
    for len(Q) > 0 {
        current := Q[0]
        Q = Q[1:]
        for k := 1; k <= G[current][0]; k++ {
            next := G[current][k]
            if minCost[next] == 0 {
                minCost[next] = minCost[current]+1
                Q = append(Q, next)
            }
        }
    }
    
    result := make([][]string, 0)
    p := param{
        current: 0,
        target: endWordIndex,
        path: make([]string, n),
        pathLen: 0,
        visited: make([]bool, n),
        G: G,
        minCost: minCost,
        wordList: wordList,
        result: &result,
    }
    dfs(p)
    
    return result
}    

